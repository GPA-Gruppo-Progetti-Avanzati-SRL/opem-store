package sql

import (
	model "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/model"
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
)

type sqlSessionRepo struct{ db *bun.DB }

func NewSessionRepo(db *bun.DB) *sqlSessionRepo { return &sqlSessionRepo{db: db} }

func (r *sqlSessionRepo) FindBySid(ctx context.Context, sid string) (*model.Session, error) {
	const semLogContext = "store-sql::session::find-by-sid"

	var row sqlSession
	err := r.db.NewSelect().Model(&row).
		Where("id = ?", sid).
		Where("status = ?", "active").
		Limit(1).Scan(ctx)
	if err != nil {
		if isNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	s := toOpemSession(&row)
	log.Debug().Str("sid", sid).Str("nickname", s.Nickname).Msg(semLogContext + " - found")
	return s, nil
}

func (r *sqlSessionRepo) Insert(ctx context.Context, s *model.Session) (string, error) {
	row := fromOpemSession(s)
	if row.ID == "" {
		row.ID = newID()
	}
	if _, err := r.db.NewInsert().Model(row).Exec(ctx); err != nil {
		return "", err
	}
	return row.ID, nil
}

func (r *sqlSessionRepo) UpdateUser(ctx context.Context, sid, nickname, userID string) error {
	_, err := r.db.NewUpdate().TableExpr("opem_session").
		Set("nickname = ?", nickname).
		Set("userid = ?", userID).
		Set("modified_at = ?", time.Now()).
		Where("id = ?", sid).
		Exec(ctx)
	return err
}

func (r *sqlSessionRepo) Invalidate(_ string) {}
