package sql

import (
	model "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/model"
	"context"

	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
)

type sqlDomainRepo struct{ db *bun.DB }

func NewDomainRepo(db *bun.DB) *sqlDomainRepo { return &sqlDomainRepo{db: db} }

func (r *sqlDomainRepo) FindByCode(ctx context.Context, code string) (*model.Domain, bool, error) {
	const semLogContext = "store-sql::domain::find-by-code"

	var row sqlDomain
	err := r.db.NewSelect().Model(&row).
		Where("code = ?", code).
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
	d := toOpemDomain(&row, apps)
	log.Debug().Str("code", code).Msg(semLogContext + " - found")
	return d, true, nil
}

func (r *sqlDomainRepo) Query(ctx context.Context) (model.DomainQueryResult, error) {
	var rows []sqlDomain
	err := r.db.NewSelect().Model(&rows).
		Where("status = ?", "active").
		Where("code != ?", "root").
		Scan(ctx)
	if err != nil && !isNotFound(err) {
		return model.DomainQueryResult{}, err
	}

	data := make([]model.Domain, 0, len(rows))
	for i := range rows {
		apps, err := r.loadApps(ctx, rows[i].ID)
		if err != nil {
			return model.DomainQueryResult{}, err
		}
		data = append(data, *toOpemDomain(&rows[i], apps))
	}
	return model.DomainQueryResult{Records: len(data), Data: data}, nil
}

func (r *sqlDomainRepo) loadApps(ctx context.Context, ownerID string) ([]sqlApp, error) {
	var apps []sqlApp
	err := r.db.NewSelect().Model(&apps).
		Where("owner_type = ?", "domain").
		Where("owner_id = ?", ownerID).
		OrderExpr("ord ASC").
		Scan(ctx)
	if err != nil && !isNotFound(err) {
		return nil, err
	}
	return apps, nil
}
