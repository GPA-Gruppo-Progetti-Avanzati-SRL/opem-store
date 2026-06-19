package acl

import (
	"context"
	"errors"
	"time"

	model "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// activeFilter è il filtro standard per escludere documenti non attivi.
var activeFilter = bson.E{Key: "sys_info.status", Value: "active"}

// optionalFieldFilter costruisce un filtro MongoDB che matcha documenti in cui:
//   - il campo ha valore `value` oppure "*"
//   - oppure il campo non esiste affatto (semantica "wildcard implicita")
func optionalFieldFilter(field, value string) bson.D {
	return bson.D{{Key: "$or", Value: bson.A{
		bson.D{{Key: field, Value: bson.D{{Key: "$in", Value: bson.A{value, "*"}}}}},
		bson.D{{Key: field, Value: bson.D{{Key: "$exists", Value: false}}}},
	}}}
}

// ResolveCapabilities risolve il grafo completo role-caps → cap-group → cap-def
// per tutti i ruoli forniti e ritorna map[app]map[category][]model.CapDef.
func ResolveCapabilities(ctx context.Context, coll *mongo.Collection, domain, site string, roles []string) (map[string]map[string][]model.CapDef, error) {
	const semLogContext = "acl::resolve-capabilities"

	entries, err := FindRoleCaps(ctx, coll, domain, site, "", roles)
	if err != nil {
		return nil, err
	}
	if len(entries) == 0 {
		log.Debug().Str("domain", domain).Str("site", site).Strs("roles", roles).
			Msg(semLogContext + " - no role-caps found")
		return nil, nil
	}

	allCgCodes := make(map[string]struct{})
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

	cgByCode := make(map[string]model.CapGroup)
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
			cgByCode[cg.OId] = cg
		}
		log.Debug().Int("codes", len(codes)).Int("cap-groups", len(allCgs)).
			Msg(semLogContext + " - cap-groups bulk loaded")
	}

	allCapDefIds := make(map[string]struct{})
	for id := range directCapDefIds {
		allCapDefIds[id] = struct{}{}
	}

	for _, entry := range entries {
		log.Debug().Str("role", entry.Role).Msg(semLogContext + " - expanding role-caps entry")
		for _, cgCode := range entry.CapGroups {
			cg, exists := cgByCode[cgCode]
			if !exists {
				log.Warn().Str("code", cgCode).Str("role", entry.Role).
					Msg(semLogContext + " - cap-group not found, skipping")
				continue
			}
			for _, capId := range cg.Capabilities {
				if capId != "" {
					allCapDefIds[capId] = struct{}{}
				}
			}
		}
	}

	result := make(map[string]map[string][]model.CapDef)
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
				result[cd.App] = make(map[string][]model.CapDef)
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
func ResolveCapabilitiesPerSite(ctx context.Context, coll *mongo.Collection, domain string, roles []string) (map[string]map[string]map[string][]model.CapDef, error) {
	const semLogContext = "acl::resolve-capabilities-per-site"

	entries, err := FindRoleCaps(ctx, coll, domain, "", "", roles)
	if err != nil {
		return nil, err
	}
	if len(entries) == 0 {
		log.Debug().Str("domain", domain).Strs("roles", roles).
			Msg(semLogContext + " - no role-caps found")
		return nil, nil
	}

	allCgCodes := make(map[string]struct{})
	for _, entry := range entries {
		for _, cg := range entry.CapGroups {
			if cg != "" {
				allCgCodes[cg] = struct{}{}
			}
		}
	}

	cgByCode := make(map[string]model.CapGroup)
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
			cgByCode[cg.OId] = cg
		}
		log.Debug().Int("codes", len(codes)).Int("cap-groups", len(allCgs)).
			Msg(semLogContext + " - cap-groups bulk loaded")
	}

	type siteCapKey struct{ site, capId string }
	siteCapIds := make(map[siteCapKey]struct{})

	for _, entry := range entries {
		siteKey := entry.Site
		for _, capId := range entry.Capabilities {
			if capId != "" {
				siteCapIds[siteCapKey{siteKey, capId}] = struct{}{}
			}
		}
		for _, cgCode := range entry.CapGroups {
			cg, exists := cgByCode[cgCode]
			if !exists {
				log.Warn().Str("code", cgCode).Str("role", entry.Role).
					Msg(semLogContext + " - cap-group not found, skipping")
				continue
			}
			for _, capId := range cg.Capabilities {
				if capId != "" {
					siteCapIds[siteCapKey{siteKey, capId}] = struct{}{}
				}
			}
		}
	}

	if len(siteCapIds) == 0 {
		return nil, nil
	}

	allCapDefIds := make(map[string]struct{})
	for k := range siteCapIds {
		allCapDefIds[k.capId] = struct{}{}
	}
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

	cdById := make(map[string]model.CapDef, len(allCds))
	for _, cd := range allCds {
		cdById[cd.OId] = cd
	}

	result := make(map[string]map[string]map[string][]model.CapDef)
	for k := range siteCapIds {
		cd, ok := cdById[k.capId]
		if !ok {
			continue
		}
		s := k.site
		if result[s] == nil {
			result[s] = make(map[string]map[string][]model.CapDef)
		}
		if result[s][cd.App] == nil {
			result[s][cd.App] = make(map[string][]model.CapDef)
		}
		result[s][cd.App][cd.Category] = append(result[s][cd.App][cd.Category], cd)
	}

	log.Debug().Str("domain", domain).Strs("roles", roles).
		Int("site-buckets", len(result)).Msg(semLogContext + " - resolved")
	return result, nil
}

// FindRoleCaps cerca tutti i documenti role-caps per un insieme di ruoli.
func FindRoleCaps(ctx context.Context, collection *mongo.Collection, domain, site, app string, roles []string) ([]model.RoleCapsEntry, error) {
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

	var entries []model.RoleCapsEntry
	if err := cursor.All(ctx, &entries); err != nil {
		log.Error().Err(err).Msg(semLogContext + " - decode error")
		return nil, err
	}

	log.Debug().Strs("roles", roles).Str("domain", domain).Str("site", site).
		Str("app", app).Int("entries", len(entries)).Msg(semLogContext)
	return entries, nil
}

// FindCapGroup cerca un documento cap-group tramite il suo _id.
func FindCapGroup(ctx context.Context, collection *mongo.Collection, code string) (*model.CapGroup, error) {
	const semLogContext = "acl::find-cap-group"

	filter := bson.D{
		{Key: "_id", Value: code},
		{Key: "_et", Value: EntityTypeCapGroup},
		activeFilter,
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var cg model.CapGroup
	if err := collection.FindOne(ctx, filter).Decode(&cg); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Debug().Str("code", code).Msg(semLogContext + " - not found")
			return nil, nil
		}
		log.Error().Err(err).Str("code", code).Msg(semLogContext + " - db error")
		return nil, err
	}

	log.Debug().Str("code", code).Msg(semLogContext + " - found")
	return &cg, nil
}

// FindCapDef cerca un documento cap-def tramite il suo _id.
func FindCapDef(ctx context.Context, collection *mongo.Collection, id string) (*model.CapDef, error) {
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

	var cd model.CapDef
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

// FindCapGroupsByCodes restituisce tutti i cap-group i cui _id compaiono nell'elenco.
func FindCapGroupsByCodes(ctx context.Context, collection *mongo.Collection, codes []string) ([]model.CapGroup, error) {
	const semLogContext = "acl::find-cap-groups-by-codes"

	if len(codes) == 0 {
		return nil, nil
	}

	filter := bson.D{
		{Key: "_id", Value: bson.D{{Key: "$in", Value: codes}}},
		{Key: "_et", Value: EntityTypeCapGroup},
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

	var cgs []model.CapGroup
	if err := cursor.All(ctx, &cgs); err != nil {
		log.Error().Err(err).Msg(semLogContext + " - decode error")
		return nil, err
	}

	log.Debug().Strs("codes", codes).Int("found", len(cgs)).Msg(semLogContext)
	return cgs, nil
}

// FindCapDefsByIds restituisce tutte le cap-def i cui _id compaiono nell'elenco.
func FindCapDefsByIds(ctx context.Context, collection *mongo.Collection, ids []string) ([]model.CapDef, error) {
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

	var cds []model.CapDef
	if err := cursor.All(ctx, &cds); err != nil {
		log.Error().Err(err).Msg(semLogContext + " - decode error")
		return nil, err
	}

	log.Debug().Strs("ids", ids).Int("found", len(cds)).Msg(semLogContext)
	return cds, nil
}

// FindAllCapGroupsByCode è mantenuta per compatibilità.
// Deprecated: usare FindCapGroupsByCodes(ctx, coll, []string{code}).
func FindAllCapGroupsByCode(ctx context.Context, collection *mongo.Collection, code string) ([]model.CapGroup, error) {
	return FindCapGroupsByCodes(ctx, collection, []string{code})
}
