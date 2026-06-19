package sql

import (
	model "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/model"
	"context"

	"github.com/uptrace/bun"
)

type sqlSequenceRepo struct{ db *bun.DB }

func NewSequenceRepo(db *bun.DB) *sqlSequenceRepo { return &sqlSequenceRepo{db: db} }

func (r *sqlSequenceRepo) QueryByDomainSite(ctx context.Context, domainCode, siteCode string) (model.SequenceQueryResult, error) {
	siteCode = translateSiteCode(siteCode)

	var rows []sqlSequence
	err := r.db.NewSelect().Model(&rows).
		Where("domain_code = ?", domainCode).
		Where("site_code = ?", siteCode).
		Scan(ctx)
	if err != nil && !isNotFound(err) {
		return model.SequenceQueryResult{}, err
	}

	data := make([]model.Sequence, 0, len(rows))
	for i := range rows {
		data = append(data, toOpemSequence(&rows[i]))
	}
	return model.SequenceQueryResult{Records: len(data), Data: data}, nil
}
