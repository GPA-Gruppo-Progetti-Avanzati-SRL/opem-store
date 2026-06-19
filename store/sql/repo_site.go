package sql

import (
	model "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/model"
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
)

type sqlSiteRepo struct{ db *bun.DB }

func NewSiteRepo(db *bun.DB) *sqlSiteRepo { return &sqlSiteRepo{db: db} }

func (r *sqlSiteRepo) FindByCode(ctx context.Context, domainCode, siteCode string) (*model.Site, bool, error) {
	const semLogContext = "store-sql::site::find-by-code"

	var row sqlSite
	err := r.db.NewSelect().Model(&row).
		Where("domain_code = ?", domainCode).
		Where("code = ?", siteCode).
		Where("status = ?", "active").
		Limit(1).Scan(ctx)
	if err != nil {
		if isNotFound(err) {
			return nil, false, nil
		}
		return nil, false, err
	}

	apps, err := r.loadApps(ctx, row.ID)
	if err != nil {
		return nil, false, err
	}
	s := toOpemSite(&row, apps)
	log.Debug().Str("domain", domainCode).Str("site", siteCode).Msg(semLogContext + " - found")
	return s, true, nil
}

func (r *sqlSiteRepo) QueryByDomain(ctx context.Context, domainCode string, opts model.SiteQueryOptions) (model.SiteQueryResult, error) {
	q := r.db.NewSelect().Model((*sqlSite)(nil)).
		Where("domain_code = ?", domainCode).
		Where("status = ?", "active")

	if opts.SortBy != "" {
		dir := "ASC"
		if opts.SortDirection == "desc" {
			dir = "DESC"
		}
		q = q.OrderExpr(fmt.Sprintf("%s %s", opts.SortBy, dir))
	} else {
		q = q.OrderExpr("ord ASC")
	}
	if opts.Limit > 0 {
		q = q.Limit(int(opts.Limit))
	}
	if opts.Offset > 0 {
		q = q.Offset(int(opts.Offset))
	}

	var rows []sqlSite
	if err := q.Scan(ctx, &rows); err != nil && !isNotFound(err) {
		return model.SiteQueryResult{}, err
	}

	data := make([]model.Site, 0, len(rows))
	for i := range rows {
		apps, err := r.loadApps(ctx, rows[i].ID)
		if err != nil {
			return model.SiteQueryResult{}, err
		}
		data = append(data, *toOpemSite(&rows[i], apps))
	}
	return model.SiteQueryResult{Records: len(data), Data: data}, nil
}

func (r *sqlSiteRepo) loadApps(ctx context.Context, ownerID string) ([]sqlApp, error) {
	var apps []sqlApp
	err := r.db.NewSelect().Model(&apps).
		Where("owner_type = ?", "site").
		Where("owner_id = ?", ownerID).
		OrderExpr("ord ASC").
		Scan(ctx)
	if err != nil && !isNotFound(err) {
		return nil, err
	}
	return apps, nil
}
