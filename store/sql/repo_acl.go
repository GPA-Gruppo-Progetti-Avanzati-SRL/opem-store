package sql

import (
	model "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/model"
	"context"

	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
)

type sqlACLRepo struct{ db *bun.DB }

func NewACLRepo(db *bun.DB) *sqlACLRepo { return &sqlACLRepo{db: db} }

func (r *sqlACLRepo) FindRoleCaps(ctx context.Context, domainReq, siteReq, appReq string, roles []string) ([]model.RoleCapsEntry, error) {
	if len(roles) == 0 {
		return nil, nil
	}

	var rows []sqlRoleCaps
	q := r.db.NewSelect().Model(&rows).
		Where("role IN (?)", bun.List(roles)).
		Where("status = ?", "active")
	if domainReq != "" {
		q = q.Where("(domain_code IS NULL OR domain_code = '' OR domain_code = '*' OR domain_code = ?)", domainReq)
	}
	if siteReq != "" {
		q = q.Where("(site_code IS NULL OR site_code = '' OR site_code = '*' OR site_code = ?)", siteReq)
	}
	if appReq != "" {
		q = q.Where("(app IS NULL OR app = '' OR app = '*' OR app = ?)", appReq)
	}
	if err := q.Scan(ctx); err != nil && !isNotFound(err) {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, nil
	}

	// Batch: carica cap-group e cap-def di tutte le righe in 2 query invece di 2N.
	rowIDs := make([]string, len(rows))
	for i, row := range rows {
		rowIDs[i] = row.ID
	}
	capGroupsByRC, err := r.loadCapGroupIDsBatch(ctx, rowIDs)
	if err != nil {
		return nil, err
	}
	capDefsByRC, err := r.loadCapDefIDsBatch(ctx, rowIDs)
	if err != nil {
		return nil, err
	}

	result := make([]model.RoleCapsEntry, 0, len(rows))
	for _, row := range rows {
		result = append(result, toOpemRoleCaps(&row, capGroupsByRC[row.ID], capDefsByRC[row.ID]))
	}
	return result, nil
}

func (r *sqlACLRepo) FindCapGroup(ctx context.Context, code string) (*model.CapGroup, error) {
	var row sqlCapGroup
	err := r.db.NewSelect().Model(&row).
		Where("id = ?", code).
		Where("status = ?", "active").
		Limit(1).Scan(ctx)
	if err != nil {
		if isNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	capDefsByGroup, err := r.loadCapGroupDefIDsBatch(ctx, []string{code})
	if err != nil {
		return nil, err
	}
	return toOpemCapGroup(&row, capDefsByGroup[code]), nil
}

func (r *sqlACLRepo) FindCapDef(ctx context.Context, id string) (*model.CapDef, error) {
	if id == "" {
		return nil, nil
	}
	var row sqlCapDef
	err := r.db.NewSelect().Model(&row).Where("id = ?", id).Limit(1).Scan(ctx)
	if err != nil {
		if isNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return toOpemCapDef(&row), nil
}

func (r *sqlACLRepo) FindCapabilitiesByRoles(ctx context.Context, domainReq, siteReq string, roles []string) (map[string]map[string][]model.CapDef, error) {
	const semLogContext = "store-sql::acl::find-capabilities-by-roles"

	entries, err := r.FindRoleCaps(ctx, domainReq, siteReq, "", roles)
	if err != nil {
		return nil, err
	}
	if len(entries) == 0 {
		log.Debug().Str("domain", domainReq).Str("site", siteReq).Strs("roles", roles).
			Msg(semLogContext + " - no entries found")
		return nil, nil
	}

	allCapDefIDs, err := r.resolveCapDefIDs(ctx, entries)
	if err != nil {
		return nil, err
	}
	if len(allCapDefIDs) == 0 {
		return nil, nil
	}

	return r.buildCapMap(ctx, allCapDefIDs)
}

func (r *sqlACLRepo) FindCapabilitiesPerSite(ctx context.Context, domainReq string, roles []string) (map[string]map[string]map[string][]model.CapDef, error) {
	const semLogContext = "store-sql::acl::find-capabilities-per-site"

	entries, err := r.FindRoleCaps(ctx, domainReq, "", "", roles)
	if err != nil {
		return nil, err
	}
	if len(entries) == 0 {
		log.Debug().Str("domain", domainReq).Strs("roles", roles).
			Msg(semLogContext + " - no entries found")
		return nil, nil
	}

	// Batch: risolvi tutti i cap-group -> cap-def in 1 query.
	cgIDSet := make(map[string]struct{})
	for _, entry := range entries {
		for _, cgCode := range entry.CapGroups {
			if cgCode != "" {
				cgIDSet[cgCode] = struct{}{}
			}
		}
	}
	cgToCapDefs := make(map[string][]string)
	if len(cgIDSet) > 0 {
		cgIDs := make([]string, 0, len(cgIDSet))
		for id := range cgIDSet {
			cgIDs = append(cgIDs, id)
		}
		var err error
		cgToCapDefs, err = r.loadCapGroupDefIDsBatch(ctx, cgIDs)
		if err != nil {
			return nil, err
		}
	}

	// Costruisce la mappa (site, capID) de-duplicata.
	type siteCapKey struct{ site, capID string }
	siteCapIDs := make(map[siteCapKey]struct{})
	for _, entry := range entries {
		site := entry.Site
		for _, capID := range entry.Capabilities {
			if capID != "" {
				siteCapIDs[siteCapKey{site, capID}] = struct{}{}
			}
		}
		for _, cgCode := range entry.CapGroups {
			for _, capID := range cgToCapDefs[cgCode] {
				if capID != "" {
					siteCapIDs[siteCapKey{site, capID}] = struct{}{}
				}
			}
		}
	}
	if len(siteCapIDs) == 0 {
		return nil, nil
	}

	// Batch: carica tutte le cap-def distinte in 1 query.
	seen := make(map[string]struct{}, len(siteCapIDs))
	distinctIDs := make([]string, 0, len(siteCapIDs))
	for k := range siteCapIDs {
		if _, ok := seen[k.capID]; !ok {
			distinctIDs = append(distinctIDs, k.capID)
			seen[k.capID] = struct{}{}
		}
	}
	capDefMap, err := r.findCapDefsByIDs(ctx, distinctIDs)
	if err != nil {
		return nil, err
	}

	log.Debug().
		Str("domain", domainReq).
		Int("cap-defs", len(capDefMap)).
		Int("site-cap-pairs", len(siteCapIDs)).
		Msg(semLogContext + " - batch load complete")

	result := make(map[string]map[string]map[string][]model.CapDef)
	for k := range siteCapIDs {
		cd, ok := capDefMap[k.capID]
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
		result[s][cd.App][cd.Category] = append(result[s][cd.App][cd.Category], *cd)
	}
	return result, nil
}


// -- helper batch -------------------------------------------------------------

// loadCapGroupIDsBatch carica i cap-group-id per tutti i role-caps-id dati in 1 query.
// Restituisce map[roleCapsID][]capGroupID.
func (r *sqlACLRepo) loadCapGroupIDsBatch(ctx context.Context, roleCapsIDs []string) (map[string][]string, error) {
	if len(roleCapsIDs) == 0 {
		return nil, nil
	}
	var rows []sqlRoleCapGroup
	err := r.db.NewSelect().Model(&rows).Where("role_caps_id IN (?)", bun.List(roleCapsIDs)).Scan(ctx)
	if err != nil && !isNotFound(err) {
		return nil, err
	}
	result := make(map[string][]string, len(roleCapsIDs))
	for _, row := range rows {
		result[row.RoleCapsID] = append(result[row.RoleCapsID], row.CapGroupID)
	}
	return result, nil
}

// loadCapDefIDsBatch carica i cap-def-id per tutti i role-caps-id dati in 1 query.
// Restituisce map[roleCapsID][]capDefID.
func (r *sqlACLRepo) loadCapDefIDsBatch(ctx context.Context, roleCapsIDs []string) (map[string][]string, error) {
	if len(roleCapsIDs) == 0 {
		return nil, nil
	}
	var rows []sqlRoleCapDef
	err := r.db.NewSelect().Model(&rows).Where("role_caps_id IN (?)", bun.List(roleCapsIDs)).Scan(ctx)
	if err != nil && !isNotFound(err) {
		return nil, err
	}
	result := make(map[string][]string, len(roleCapsIDs))
	for _, row := range rows {
		result[row.RoleCapsID] = append(result[row.RoleCapsID], row.CapDefID)
	}
	return result, nil
}

// loadCapGroupDefIDsBatch carica i cap-def-id per tutti i cap-group-id dati in 1 query.
// Restituisce map[capGroupID][]capDefID.
func (r *sqlACLRepo) loadCapGroupDefIDsBatch(ctx context.Context, capGroupIDs []string) (map[string][]string, error) {
	if len(capGroupIDs) == 0 {
		return nil, nil
	}
	var rows []sqlCapGroupDef
	err := r.db.NewSelect().Model(&rows).Where("cap_group_id IN (?)", bun.List(capGroupIDs)).Scan(ctx)
	if err != nil && !isNotFound(err) {
		return nil, err
	}
	result := make(map[string][]string, len(capGroupIDs))
	for _, row := range rows {
		result[row.CapGroupID] = append(result[row.CapGroupID], row.CapDefID)
	}
	return result, nil
}

// findCapDefsByIDs carica tutte le cap-def con gli id dati in 1 query.
// Restituisce map[capDefID]*model.CapDef.
func (r *sqlACLRepo) findCapDefsByIDs(ctx context.Context, ids []string) (map[string]*model.CapDef, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	var rows []sqlCapDef
	err := r.db.NewSelect().Model(&rows).Where("id IN (?)", bun.List(ids)).Scan(ctx)
	if err != nil && !isNotFound(err) {
		return nil, err
	}
	result := make(map[string]*model.CapDef, len(rows))
	for i := range rows {
		cd := toOpemCapDef(&rows[i])
		result[cd.OId] = cd
	}
	return result, nil
}

func (r *sqlACLRepo) resolveCapDefIDs(ctx context.Context, entries []model.RoleCapsEntry) (map[string]struct{}, error) {
	// Batch: risolvi tutti i cap-group in 1 query invece di 1 per cap-group.
	cgIDSet := make(map[string]struct{})
	for _, entry := range entries {
		for _, cgCode := range entry.CapGroups {
			if cgCode != "" {
				cgIDSet[cgCode] = struct{}{}
			}
		}
	}
	cgToCapDefs := make(map[string][]string)
	if len(cgIDSet) > 0 {
		cgIDs := make([]string, 0, len(cgIDSet))
		for id := range cgIDSet {
			cgIDs = append(cgIDs, id)
		}
		var err error
		cgToCapDefs, err = r.loadCapGroupDefIDsBatch(ctx, cgIDs)
		if err != nil {
			return nil, err
		}
	}

	allIDs := make(map[string]struct{})
	for _, entry := range entries {
		for _, capID := range entry.Capabilities {
			if capID != "" {
				allIDs[capID] = struct{}{}
			}
		}
		for _, cgCode := range entry.CapGroups {
			for _, id := range cgToCapDefs[cgCode] {
				if id != "" {
					allIDs[id] = struct{}{}
				}
			}
		}
	}
	return allIDs, nil
}

func (r *sqlACLRepo) buildCapMap(ctx context.Context, capDefIDs map[string]struct{}) (map[string]map[string][]model.CapDef, error) {
	if len(capDefIDs) == 0 {
		return nil, nil
	}
	ids := make([]string, 0, len(capDefIDs))
	for id := range capDefIDs {
		ids = append(ids, id)
	}
	// Batch: carica tutte le cap-def in 1 query invece di 1 per ID.
	capDefMap, err := r.findCapDefsByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	result := make(map[string]map[string][]model.CapDef, len(capDefMap))
	for _, cd := range capDefMap {
		if result[cd.App] == nil {
			result[cd.App] = make(map[string][]model.CapDef)
		}
		result[cd.App][cd.Category] = append(result[cd.App][cd.Category], *cd)
	}
	if len(result) == 0 {
		return nil, nil
	}
	return result, nil
}
