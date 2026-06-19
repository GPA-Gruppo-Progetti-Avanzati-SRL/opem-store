package sql

import (
	model "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/model"
	"context"

	"github.com/uptrace/bun"
)

type sqlKVRepo struct{ db *bun.DB }

func NewKVRepo(db *bun.DB) *sqlKVRepo { return &sqlKVRepo{db: db} }

func (r *sqlKVRepo) FindByDomainSiteNameCategory(ctx context.Context, domainCode, siteCode, pkgName, category string) (model.KeyValuePackage, bool, error) {
	siteCode = translateSiteCode(siteCode)
	scopes := scopesForQuery(domainCode, siteCode)

	q := r.db.NewSelect().Model((*sqlKV)(nil)).
		Where("bid = ?", pkgName).
		Where("scope IN (?)", bun.In(scopes)).
		Where("status = ?", "active")
	if category != "" {
		q = q.Where("category = ?", category)
	}

	var rows []sqlKV
	if err := q.Scan(ctx, &rows); err != nil && !isNotFound(err) {
		return model.KeyValuePackage{}, false, err
	}

	var best *sqlKV
	for i := range rows {
		if best == nil || len(rows[i].Scope) > len(best.Scope) {
			best = &rows[i]
		}
	}
	if best == nil {
		return model.KeyValuePackage{}, false, nil
	}

	props, err := r.loadProps(ctx, best.ID)
	if err != nil {
		return model.KeyValuePackage{}, false, err
	}

	kvp := toOpemKV(best, props)

	reqScope := scopeFromDomainSite(domainCode, siteCode)
	if best.Scope != reqScope {
		kvp.Inherited = true
	}

	return kvp, true, nil
}

func (r *sqlKVRepo) FindByDomainSiteCategories(ctx context.Context, domainCode, siteCode string, categories []string) ([]model.KeyValuePackage, error) {
	siteCode = translateSiteCode(siteCode)
	scopes := scopesForQuery(domainCode, siteCode)

	q := r.db.NewSelect().Model((*sqlKV)(nil)).
		Where("scope IN (?)", bun.In(scopes)).
		Where("status = ?", "active")
	if len(categories) > 0 {
		q = q.Where("category IN (?)", bun.In(categories))
	}

	var rows []sqlKV
	if err := q.Scan(ctx, &rows); err != nil && !isNotFound(err) {
		return nil, err
	}

	type bidCatKey struct{ bid, cat string }
	byBidCat := make(map[bidCatKey]*sqlKV)
	for i := range rows {
		k := bidCatKey{rows[i].Bid, rows[i].Category}
		existing, found := byBidCat[k]
		if !found || len(rows[i].Scope) > len(existing.Scope) {
			byBidCat[k] = &rows[i]
		}
	}

	reqScope := scopeFromDomainSite(domainCode, siteCode)
	result := make([]model.KeyValuePackage, 0, len(byBidCat))
	for _, row := range byBidCat {
		props, err := r.loadProps(ctx, row.ID)
		if err != nil {
			return nil, err
		}
		kvp := toOpemKV(row, props)
		if row.Scope != reqScope {
			kvp.Inherited = true
		}
		result = append(result, kvp)
	}
	return result, nil
}

func (r *sqlKVRepo) loadProps(ctx context.Context, kvID string) ([]sqlKVProperty, error) {
	var props []sqlKVProperty
	err := r.db.NewSelect().Model(&props).
		Where("kv_id = ?", kvID).
		OrderExpr("ord ASC").
		Scan(ctx)
	if err != nil && !isNotFound(err) {
		return nil, err
	}
	return props, nil
}

// translateSiteCode converte "star" in "*" (SiteWildCard).
func translateSiteCode(siteCode string) string {
	if siteCode == "star" {
		return "*"
	}
	return siteCode
}
