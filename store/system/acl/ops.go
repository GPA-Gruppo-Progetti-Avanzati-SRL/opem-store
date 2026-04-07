package acl

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// ResolveCapabilities risolve il grafo completo role-caps → cap-group → cap-def
// per tutti i ruoli forniti e ritorna map[app]map[category][]CapDef.
//
// Esegue esattamente 3 query MongoDB (indipendentemente dalla cardinalità dei dati):
//   1. FindRoleCaps          — 1 query con $in sui ruoli
//   2. FindCapGroupsByCodes  — 1 query con $in su tutti i codici di cap-group raccolti
//   3. FindCapDefsByIds      — 1 query con $in su tutti gli _id di cap-def raccolti
//
// Tutto il join è fatto in memoria. Questa funzione è chiamata solo in caso di
// cache miss; la gestione della cache è in cache.go.
func ResolveCapabilities(ctx context.Context, coll *mongo.Collection, domain, site string, roles []string) (map[string]map[string][]CapDef, error) {
	const semLogContext = "acl::resolve-capabilities"

	// ── Query 1: role-caps ────────────────────────────────────────────────────
	entries, err := FindRoleCaps(ctx, coll, domain, site, "", roles)
	if err != nil {
		return nil, err
	}
	if len(entries) == 0 {
		log.Debug().Str("domain", domain).Str("site", site).Strs("roles", roles).
			Msg(semLogContext + " - no role-caps found")
		return nil, nil
	}

	// ── Raccolta codici cap-group e cap-def diretti ───────────────────────────
	// Set di codici cap-group da risolvere con un'unica query bulk.
	allCgCodes := make(map[string]struct{})
	// cap-def ids diretti (da entry.Capabilities, senza passare per cap-group).
	directCapDefIds := make(map[string]struct{})

	for _, entry := range entries {
		for _, cgCode := range entry.CapGroups {
			if cgCode != "" {
				allCgCodes[cgCode] = struct{}{}
			}
		}
		for _, capId := range entry.Capabilities {
			if capId != "" {
				directCapDefIds[capId] = struct{}{}
			}
		}
	}

	// ── Query 2: cap-groups bulk ──────────────────────────────────────────────
	// Una sola query con $in su tutti i codici; risultato in memoria.
	//
	// cgByCodeApp: code → app → CapGroup
	// Permette in memoria di applicare la stessa semantica di FindCapGroup:
	//   - entry globale (App=""): usa TUTTI i cap-group per quel code (qualsiasi app)
	//   - entry app-specifica: preferisce app esatta, fallback su "*"
	cgByCodeApp := make(map[string]map[string]CapGroup)

	if len(allCgCodes) > 0 {
		codes := make([]string, 0, len(allCgCodes))
		for c := range allCgCodes {
			codes = append(codes, c)
		}

		allCgs, cgErr := FindCapGroupsByCodes(ctx, coll, codes)
		if cgErr != nil {
			return nil, cgErr
		}
		for _, cg := range allCgs {
			if cgByCodeApp[cg.Code] == nil {
				cgByCodeApp[cg.Code] = make(map[string]CapGroup)
			}
			// In caso di duplicati (code, app) — situazione anomala — l'ultimo vince.
			cgByCodeApp[cg.Code][cg.App] = cg
		}
		log.Debug().Int("codes", len(codes)).Int("cap-groups", len(allCgs)).
			Msg(semLogContext + " - cap-groups bulk loaded")
	}

	// ── Join in memoria: role-caps + cap-groups → set di cap-def ids ──────────
	allCapDefIds := make(map[string]struct{})
	// Propaga cap-def ids diretti già raccolti.
	for id := range directCapDefIds {
		allCapDefIds[id] = struct{}{}
	}

	for _, entry := range entries {
		isGlobal := entry.App == ""
		log.Debug().Str("role", entry.Role).Str("app", entry.App).Bool("global", isGlobal).
			Msg(semLogContext + " - expanding role-caps entry")

		for _, cgCode := range entry.CapGroups {
			appMap, exists := cgByCodeApp[cgCode]
			if !exists || len(appMap) == 0 {
				log.Warn().Str("code", cgCode).Str("app", entry.App).
					Msg(semLogContext + " - cap-group not found, skipping")
				continue
			}

			if isGlobal {
				// Usa TUTTI i cap-group per questo codice (tutte le app).
				for _, cg := range appMap {
					log.Debug().Str("code", cgCode).Str("cg-app", cg.App).
						Int("ids", len(cg.Capabilities)).
						Msg(semLogContext + " - cap-group (global) expanded")
					for _, capId := range cg.Capabilities {
						if capId != "" {
							allCapDefIds[capId] = struct{}{}
						}
					}
				}
			} else {
				// Preferisce il cap-group app-specifico; fallback su "*".
				cg, ok := appMap[entry.App]
				if !ok {
					cg, ok = appMap["*"]
				}
				if !ok {
					log.Warn().Str("code", cgCode).Str("app", entry.App).
						Msg(semLogContext + " - cap-group not found for app or *, skipping")
					continue
				}
				for _, capId := range cg.Capabilities {
					if capId != "" {
						allCapDefIds[capId] = struct{}{}
					}
				}
			}
		}
	}

	// ── Query 3: cap-def bulk ─────────────────────────────────────────────────
	result := make(map[string]map[string][]CapDef)

	if len(allCapDefIds) > 0 {
		ids := make([]string, 0, len(allCapDefIds))
		for id := range allCapDefIds {
			ids = append(ids, id)
		}

		allCds, cdErr := FindCapDefsByIds(ctx, coll, ids)
		if cdErr != nil {
			return nil, cdErr
		}
		log.Debug().Int("ids", len(ids)).Int("cap-defs", len(allCds)).
			Msg(semLogContext + " - cap-defs bulk loaded")

		for _, cd := range allCds {
			if result[cd.App] == nil {
				result[cd.App] = make(map[string][]CapDef)
			}
			result[cd.App][cd.Category] = append(result[cd.App][cd.Category], cd)
		}
	}

	log.Debug().Str("domain", domain).Str("site", site).Strs("roles", roles).
		Int("apps", len(result)).Msg(semLogContext + " - resolved")
	return result, nil
}

// ResolveCapabilitiesPerSite risolve le capabilities per tutti i ruoli forniti
// mantenendo la contestualizzazione per site derivata dai documenti role-caps.
//
// Restituisce map[site]map[app]map[category][]CapDef dove:
//   - site="" = capabilities globali (role-caps senza campo site → valgono per tutti i site)
//   - site="FA", site="CC", ecc. = capabilities specifiche per quel site
//
// Il chiamante deve fare il merge di site="" + site specifico per ogni site a cui
// l'utente ha accesso (vedi resolveApps in host/resource.go).
//
// Esegue le stesse 3 query di ResolveCapabilities (bulk, senza N+1).
func ResolveCapabilitiesPerSite(ctx context.Context, coll *mongo.Collection, domain string, roles []string) (map[string]map[string]map[string][]CapDef, error) {
	const semLogContext = "acl::resolve-capabilities-per-site"

	// Query 1: tutti i role-caps per questi ruoli (site="" = wildcard, prende tutto)
	entries, err := FindRoleCaps(ctx, coll, domain, "", "", roles)
	if err != nil {
		return nil, err
	}
	if len(entries) == 0 {
		log.Debug().Str("domain", domain).Strs("roles", roles).
			Msg(semLogContext + " - no role-caps found")
		return nil, nil
	}

	// Raccolta codici cap-group
	allCgCodes := make(map[string]struct{})
	for _, entry := range entries {
		for _, cg := range entry.CapGroups {
			if cg != "" {
				allCgCodes[cg] = struct{}{}
			}
		}
	}

	// Query 2: cap-groups bulk
	cgByCodeApp := make(map[string]map[string]CapGroup)
	if len(allCgCodes) > 0 {
		codes := make([]string, 0, len(allCgCodes))
		for c := range allCgCodes {
			codes = append(codes, c)
		}
		allCgs, cgErr := FindCapGroupsByCodes(ctx, coll, codes)
		if cgErr != nil {
			return nil, cgErr
		}
		for _, cg := range allCgs {
			if cgByCodeApp[cg.Code] == nil {
				cgByCodeApp[cg.Code] = make(map[string]CapGroup)
			}
			cgByCodeApp[cg.Code][cg.App] = cg
		}
		log.Debug().Int("codes", len(codes)).Int("cap-groups", len(allCgs)).
			Msg(semLogContext + " - cap-groups bulk loaded")
	}

	// Tracking (siteKey, capDefId) → associa ogni cap-def al site dell'entry
	// siteKey="" = globale (nessun vincolo site sull'entry role-caps)
	type siteCapKey struct{ site, capId string }
	siteCapIds := make(map[siteCapKey]struct{})

	for _, entry := range entries {
		siteKey := entry.Site // "" per entry globali
		isGlobal := entry.App == ""

		// Cap-def dirette
		for _, capId := range entry.Capabilities {
			if capId != "" {
				siteCapIds[siteCapKey{siteKey, capId}] = struct{}{}
			}
		}

		// Cap-def via cap-group
		for _, cgCode := range entry.CapGroups {
			appMap, exists := cgByCodeApp[cgCode]
			if !exists || len(appMap) == 0 {
				log.Warn().Str("code", cgCode).Str("role", entry.Role).
					Msg(semLogContext + " - cap-group not found, skipping")
				continue
			}
			if isGlobal {
				for _, cg := range appMap {
					for _, capId := range cg.Capabilities {
						if capId != "" {
							siteCapIds[siteCapKey{siteKey, capId}] = struct{}{}
						}
					}
				}
			} else {
				cg, ok := appMap[entry.App]
				if !ok {
					cg, ok = appMap["*"]
				}
				if !ok {
					log.Warn().Str("code", cgCode).Str("app", entry.App).
						Msg(semLogContext + " - cap-group not found for app or *, skipping")
					continue
				}
				for _, capId := range cg.Capabilities {
					if capId != "" {
						siteCapIds[siteCapKey{siteKey, capId}] = struct{}{}
					}
				}
			}
		}
	}

	if len(siteCapIds) == 0 {
		return nil, nil
	}

	// Raccogli tutti gli _id di cap-def unici (indipendentemente dal site)
	allCapDefIds := make(map[string]struct{})
	for k := range siteCapIds {
		allCapDefIds[k.capId] = struct{}{}
	}
	ids := make([]string, 0, len(allCapDefIds))
	for id := range allCapDefIds {
		ids = append(ids, id)
	}

	// Query 3: cap-def bulk
	allCds, cdErr := FindCapDefsByIds(ctx, coll, ids)
	if cdErr != nil {
		return nil, cdErr
	}
	log.Debug().Int("ids", len(ids)).Int("cap-defs", len(allCds)).
		Msg(semLogContext + " - cap-defs bulk loaded")

	cdById := make(map[string]CapDef, len(allCds))
	for _, cd := range allCds {
		cdById[cd.OId] = cd
	}

	// Costruisci risultato: map[site]map[app]map[category][]CapDef
	result := make(map[string]map[string]map[string][]CapDef)
	for k := range siteCapIds {
		cd, ok := cdById[k.capId]
		if !ok {
			continue
		}
		s := k.site
		if result[s] == nil {
			result[s] = make(map[string]map[string][]CapDef)
		}
		if result[s][cd.App] == nil {
			result[s][cd.App] = make(map[string][]CapDef)
		}
		result[s][cd.App][cd.Category] = append(result[s][cd.App][cd.Category], cd)
	}

	log.Debug().Str("domain", domain).Strs("roles", roles).
		Int("site-buckets", len(result)).Msg(semLogContext + " - resolved")
	return result, nil
}

// ── helpers ───────────────────────────────────────────────────────────────────

// activeFilter è il filtro standard per escludere documenti non attivi.
var activeFilter = bson.E{Key: "sys_info.status", Value: "active"}

// optionalFieldFilter costruisce un filtro MongoDB che matcha documenti in cui:
//   - il campo ha valore `value` oppure "*"
//   - oppure il campo non esiste affatto (semantica "wildcard implicita")
//
// Se value è "" non aggiunge alcun vincolo (campo completamente ignorato).
func optionalFieldFilter(field, value string) bson.D {
	return bson.D{{Key: "$or", Value: bson.A{
		bson.D{{Key: field, Value: bson.D{{Key: "$in", Value: bson.A{value, "*"}}}}},
		bson.D{{Key: field, Value: bson.D{{Key: "$exists", Value: false}}}},
	}}}
}

// ── FindRoleCaps ──────────────────────────────────────────────────────────────

// FindRoleCaps cerca tutti i documenti role-caps per un insieme di ruoli.
//
// I filtri domain, site, app sono opzionali: se "" il campo non viene filtrato
// (come se fosse "*" — vale per qualsiasi valore).
// Vengono restituiti tutti i documenti che matchano almeno un ruolo in `roles`;
// le capabilities risultanti sono l'UNIONE di tutti i documenti trovati.
func FindRoleCaps(ctx context.Context, collection *mongo.Collection, domain, site, app string, roles []string) ([]RoleCapsEntry, error) {
	const semLogContext = "acl::find-role-caps"

	if len(roles) == 0 {
		log.Debug().Msg(semLogContext + " - empty roles list, returning empty")
		return nil, nil
	}

	filter := bson.D{
		{Key: "_et", Value: EntityTypeRoleCaps},
		{Key: "role", Value: bson.D{{Key: "$in", Value: roles}}},
		activeFilter,
	}

	// Aggiungi filtri contestuali opzionali solo se il valore è fornito.
	var andClauses bson.A
	if domain != "" {
		andClauses = append(andClauses, optionalFieldFilter("domain", domain))
	}
	if site != "" {
		andClauses = append(andClauses, optionalFieldFilter("site", site))
	}
	if app != "" {
		andClauses = append(andClauses, optionalFieldFilter("app", app))
	}
	if len(andClauses) > 0 {
		filter = append(filter, bson.E{Key: "$and", Value: andClauses})
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Error().Err(err).Strs("roles", roles).Str("domain", domain).
			Str("site", site).Str("app", app).Msg(semLogContext + " - find error")
		return nil, err
	}
	defer cursor.Close(ctx)

	var entries []RoleCapsEntry
	if err := cursor.All(ctx, &entries); err != nil {
		log.Error().Err(err).Msg(semLogContext + " - decode error")
		return nil, err
	}

	log.Debug().Strs("roles", roles).Str("domain", domain).Str("site", site).
		Str("app", app).Int("entries", len(entries)).Msg(semLogContext)
	return entries, nil
}

// ── FindCapGroup ──────────────────────────────────────────────────────────────

// FindCapGroup cerca un documento cap-group per (app, code).
// Effettua fallback su app="*" per i gruppi cross-app.
// Ritorna (nil, nil) se non trovato.
func FindCapGroup(ctx context.Context, collection *mongo.Collection, app, code string) (*CapGroup, error) {
	const semLogContext = "acl::find-cap-group"

	filter := bson.D{
		{Key: "_et", Value: EntityTypeCapGroup},
		{Key: "code", Value: code},
		{Key: "app", Value: bson.D{{Key: "$in", Value: bson.A{app, "*"}}}},
		activeFilter,
	}
	// Priorità all'app-specifico rispetto a "*":
	// ordinamento discendente su "app" → il valore specifico precede "*".
	opts := options.FindOne().SetSort(bson.D{{Key: "app", Value: -1}})

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var cg CapGroup
	if err := collection.FindOne(ctx, filter, opts).Decode(&cg); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Debug().Str("app", app).Str("code", code).Msg(semLogContext + " - not found")
			return nil, nil
		}
		log.Error().Err(err).Str("app", app).Str("code", code).Msg(semLogContext + " - db error")
		return nil, err
	}

	log.Debug().Str("app", app).Str("code", code).Msg(semLogContext + " - found")
	return &cg, nil
}

// ── FindCapDef ────────────────────────────────────────────────────────────────

// FindCapDef cerca un documento cap-def tramite il suo _id (stringa).
// Il formato è "<resource>:<azione>.<categoria>", es. "sim:list.ui".
// Lookup per _id è sempre O(1) grazie all'indice primario MongoDB.
// Ritorna (nil, nil) se non trovato o se id è vuoto.
func FindCapDef(ctx context.Context, collection *mongo.Collection, id string) (*CapDef, error) {
	const semLogContext = "acl::find-cap-def"

	if id == "" {
		log.Warn().Msg(semLogContext + " - empty id, skipping")
		return nil, nil
	}

	filter := bson.D{
		{Key: "_id", Value: id},
		activeFilter,
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var cd CapDef
	if err := collection.FindOne(ctx, filter).Decode(&cd); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Debug().Str("id", id).Msg(semLogContext + " - not found")
			return nil, nil
		}
		log.Error().Err(err).Str("id", id).Msg(semLogContext + " - db error")
		return nil, err
	}

	log.Debug().Str("id", id).Str("app", cd.App).Str("category", cd.Category).Msg(semLogContext + " - found")
	return &cd, nil
}

// ── FindCapGroupsByCodes (bulk) ───────────────────────────────────────────────

// FindCapGroupsByCodes restituisce TUTTI i cap-group il cui campo code compare
// nell'elenco fornito, indipendentemente dall'app (query unica con $in).
//
// Sostituisce le chiamate N×FindCapGroup e N×FindAllCapGroupsByCode usate in
// precedenza da ResolveCapabilities: una sola query di rete invece di K query.
//
// Il join (app-specifico vs "*" vs globale) viene poi fatto in memoria da
// ResolveCapabilities.
func FindCapGroupsByCodes(ctx context.Context, collection *mongo.Collection, codes []string) ([]CapGroup, error) {
	const semLogContext = "acl::find-cap-groups-by-codes"

	if len(codes) == 0 {
		return nil, nil
	}

	filter := bson.D{
		{Key: "_et", Value: EntityTypeCapGroup},
		{Key: "code", Value: bson.D{{Key: "$in", Value: codes}}},
		activeFilter,
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Error().Err(err).Strs("codes", codes).Msg(semLogContext + " - find error")
		return nil, err
	}
	defer cursor.Close(ctx)

	var cgs []CapGroup
	if err := cursor.All(ctx, &cgs); err != nil {
		log.Error().Err(err).Msg(semLogContext + " - decode error")
		return nil, err
	}

	log.Debug().Strs("codes", codes).Int("found", len(cgs)).Msg(semLogContext)
	return cgs, nil
}

// ── FindCapDefsByIds (bulk) ───────────────────────────────────────────────────

// FindCapDefsByIds restituisce TUTTE le cap-def i cui _id compaiono nell'elenco
// fornito (query unica con $in sull'indice primario MongoDB).
//
// Sostituisce le chiamate M×FindCapDef usate in precedenza da ResolveCapabilities:
// una sola query di rete invece di M query.
func FindCapDefsByIds(ctx context.Context, collection *mongo.Collection, ids []string) ([]CapDef, error) {
	const semLogContext = "acl::find-cap-defs-by-ids"

	if len(ids) == 0 {
		return nil, nil
	}

	filter := bson.D{
		{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}},
		activeFilter,
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Error().Err(err).Strs("ids", ids).Msg(semLogContext + " - find error")
		return nil, err
	}
	defer cursor.Close(ctx)

	var cds []CapDef
	if err := cursor.All(ctx, &cds); err != nil {
		log.Error().Err(err).Msg(semLogContext + " - decode error")
		return nil, err
	}

	log.Debug().Strs("ids", ids).Int("found", len(cds)).Msg(semLogContext)
	return cds, nil
}

// ── FindAllCapGroupsByCode (deprecated) ───────────────────────────────────────

// FindAllCapGroupsByCode è mantenuta per compatibilità ma non è più usata da
// ResolveCapabilities. Usare FindCapGroupsByCodes per query bulk.
//
// Deprecated: usare FindCapGroupsByCodes(ctx, coll, []string{code}).
func FindAllCapGroupsByCode(ctx context.Context, collection *mongo.Collection, code string) ([]CapGroup, error) {
	return FindCapGroupsByCodes(ctx, collection, []string{code})
}

// fine ops.go
