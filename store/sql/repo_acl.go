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

	result := make([]model.RoleCapsEntry, 0, len(rows))
	for _, row := range rows {
		capGroups, err := r.loadCapGroupIDs(ctx, row.ID)
		if err != nil {
			return nil, err
		}
		caps, err := r.loadCapDefIDs(ctx, row.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, toOpemRoleCaps(&row, capGroups, caps))
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
	capDefIDs, err := r.loadCapGroupDefIDs(ctx, code)
	if err != nil {
		return nil, err
	}
	return toOpemCapGroup(&row, capDefIDs), nil
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

	type siteCapKey struct{ site, capID string }
	siteCapIDs := make(map[siteCapKey]struct{})

	for _, entry := range entries {
		siteKey := entry.Site
		for _, capID := range entry.Capabilities {
			if capID != "" {
				siteCapIDs[siteCapKey{siteKey, capID}] = struct{}{}
			}
		}
		for _, cgCode := range entry.CapGroups {
			cg, err := r.FindCapGroup(ctx, cgCode)
			if err != nil || cg == nil {
				continue
			}
			for _, capID := range cg.Capabilities {
				if capID != "" {
					siteCapIDs[siteCapKey{siteKey, capID}] = struct{}{}
				}
			}
		}
	}

	if len(siteCapIDs) == 0 {
		return nil, nil
	}

	result := make(map[string]map[string]map[string][]model.CapDef)
	for k := range siteCapIDs {
		cd, err := r.FindCapDef(ctx, k.capID)
		if err != nil || cd == nil {
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

func (r *sqlACLRepo) loadCapGroupIDs(ctx context.Context, roleCapsID string) ([]string, error) {
	var rows []sqlRoleCapGroup
	err := r.db.NewSelect().Model(&rows).Where("role_caps_id = ?", roleCapsID).Scan(ctx)
	if err != nil && !isNotFound(err) {
		return nil, err
	}
	ids := make([]string, 0, len(rows))
	for _, row := range rows {
		ids = append(ids, row.CapGroupID)
	}
	return ids, nil
}

func (r *sqlACLRepo) loadCapDefIDs(ctx context.Context, roleCapsID string) ([]string, error) {
	var rows []sqlRoleCapDef
	err := r.db.NewSelect().Model(&rows).Where("role_caps_id = ?", roleCapsID).Scan(ctx)
	if err != nil && !isNotFound(err) {
		return nil, err
	}
	ids := make([]string, 0, len(rows))
	for _, row := range rows {
		ids = append(ids, row.CapDefID)
	}
	return ids, nil
}

func (r *sqlACLRepo) loadCapGroupDefIDs(ctx context.Context, capGroupID string) ([]string, error) {
	var rows []sqlCapGroupDef
	err := r.db.NewSelect().Model(&rows).Where("cap_group_id = ?", capGroupID).Scan(ctx)
	if err != nil && !isNotFound(err) {
		return nil, err
	}
	ids := make([]string, 0, len(rows))
	for _, row := range rows {
		ids = append(ids, row.CapDefID)
	}
	return ids, nil
}

func (r *sqlACLRepo) resolveCapDefIDs(ctx context.Context, entries []model.RoleCapsEntry) (map[string]struct{}, error) {
	allIDs := make(map[string]struct{})
	for _, entry := range entries {
		for _, capID := range entry.Capabilities {
			if capID != "" {
				allIDs[capID] = struct{}{}
			}
		}
		for _, cgCode := range entry.CapGroups {
			defIDs, err := r.loadCapGroupDefIDs(ctx, cgCode)
			if err != nil {
				return nil, err
			}
			for _, id := range defIDs {
				allIDs[id] = struct{}{}
			}
		}
	}
	return allIDs, nil
}

func (r *sqlACLRepo) buildCapMap(ctx context.Context, capDefIDs map[string]struct{}) (map[string]map[string][]model.CapDef, error) {
	result := make(map[string]map[string][]model.CapDef)
	for capID := range capDefIDs {
		cd, err := r.FindCapDef(ctx, capID)
		if err != nil || cd == nil {
			continue
		}
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
