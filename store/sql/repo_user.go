package sql

import (
	model "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/model"
	"context"
	"strings"
	"time"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
)

type sqlUserRepo struct{ db *bun.DB }

func NewUserRepo(db *bun.DB) *sqlUserRepo { return &sqlUserRepo{db: db} }

func (r *sqlUserRepo) FindByNickname(ctx context.Context, nickname string) (*model.User, error) {
	const semLogContext = "store-sql::user::find-by-nickname"

	var row sqlUser
	err := r.db.NewSelect().Model(&row).
		Where("nickname = ?", nickname).
		Where("et != ?", "oidc-user").
		Where("status = ?", "active").
		Limit(1).Scan(ctx)
	if err != nil {
		if isNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	roles, err := r.loadRoles(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	u := toOpemUser(&row, roles)
	log.Debug().Str("nickname", nickname).Str("id", row.ID).Msg(semLogContext + " - found")
	return u, nil
}

func (r *sqlUserRepo) FindRoleByNickname(ctx context.Context, nickname string) (*model.User, error) {
	const semLogContext = "store-sql::user::find-role-by-nickname"

	var row sqlUser
	err := r.db.NewSelect().Model(&row).
		Where("nickname = ?", nickname).
		Where("et = ?", "oidc-user").
		Limit(1).Scan(ctx)
	if err != nil {
		if isNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	roles, err := r.loadRoles(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	u := toOpemUser(&row, roles)
	log.Debug().Str("nickname", nickname).Str("id", row.ID).Msg(semLogContext + " - found")
	return u, nil
}

func (r *sqlUserRepo) FindByID(ctx context.Context, id string) (*model.User, error) {
	const semLogContext = "store-sql::user::find-by-id"

	var row sqlUser
	err := r.db.NewSelect().Model(&row).Where("id = ?", id).Limit(1).Scan(ctx)
	if err != nil {
		if isNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	roles, err := r.loadRoles(ctx, id)
	if err != nil {
		return nil, err
	}
	u := toOpemUser(&row, roles)
	log.Debug().Str("id", id).Msg(semLogContext + " - found")
	return u, nil
}

func (r *sqlUserRepo) Insert(ctx context.Context, u *model.User) (string, error) {
	row, roles := fromOpemUser(u)
	if row.ID == "" {
		row.ID = newID()
	}

	if _, err := r.db.NewInsert().Model(row).Exec(ctx); err != nil {
		return "", err
	}
	for i := range roles {
		roles[i].UserID = row.ID
		if roles[i].ID == "" {
			roles[i].ID = newID()
		}
	}
	if len(roles) > 0 {
		if _, err := r.db.NewInsert().Model(&roles).Exec(ctx); err != nil {
			return "", err
		}
	}
	return row.ID, nil
}

func (r *sqlUserRepo) UpdateInfo(ctx context.Context, id, firstname, lastname string) (*model.User, error) {
	_, err := r.db.NewUpdate().TableExpr("opem_user").
		Set("firstname = ?", firstname).
		Set("lastname = ?", lastname).
		Set("modified_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return r.FindByID(ctx, id)
}

func (r *sqlUserRepo) UpdateRoles(ctx context.Context, id string, roles commons.UserRoleList) (*model.User, error) {
	if _, err := r.db.NewDelete().TableExpr("opem_user_role").Where("user_id = ?", id).Exec(ctx); err != nil {
		return nil, err
	}
	sqlRoles := make([]sqlUserRole, 0, len(roles))
	for _, role := range roles {
		sqlRoles = append(sqlRoles, sqlUserRole{
			ID:     newID(),
			UserID: id,
			Domain: role.Domain,
			Site:   role.Site,
			Apps:   role.Apps,
		})
	}
	if len(sqlRoles) > 0 {
		if _, err := r.db.NewInsert().Model(&sqlRoles).Exec(ctx); err != nil {
			return nil, err
		}
	}
	_, err := r.db.NewUpdate().TableExpr("opem_user").
		Set("modified_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return r.FindByID(ctx, id)
}

func (r *sqlUserRepo) UpdateProfilePicture(_ context.Context, _ string, _ commons.FileReference) (*model.User, error) {
	return nil, ErrNotSupported
}

func (r *sqlUserRepo) Query(ctx context.Context, opts model.UserQueryOptions) (model.UserQueryResult, error) {
	q := r.db.NewSelect().Model((*sqlUser)(nil)).
		Where("et != ?", "oidc-user").
		Where("status = ?", "active")

	if opts.SearchTerm != "" {
		term := "%" + strings.ToLower(opts.SearchTerm) + "%"
		q = q.Where("(LOWER(nickname) LIKE ? OR LOWER(firstname) LIKE ? OR LOWER(lastname) LIKE ? OR LOWER(email) LIKE ?)",
			term, term, term, term)
	}

	var total int
	if opts.WithCount {
		cnt, err := q.Count(ctx)
		if err != nil {
			return model.UserQueryResult{}, err
		}
		total = cnt
	}

	if opts.SortBy != "" {
		dir := "ASC"
		if opts.SortDirection == "desc" {
			dir = "DESC"
		}
		q = q.OrderExpr("? ?", bun.Ident(opts.SortBy), bun.Safe(dir))
	}
	if opts.Limit > 0 {
		q = q.Limit(int(opts.Limit))
	}
	if opts.Offset > 0 {
		q = q.Offset(int(opts.Offset))
	}

	var rows []sqlUser
	if err := q.Scan(ctx, &rows); err != nil {
		return model.UserQueryResult{}, err
	}

	data := make([]model.User, 0, len(rows))
	for i := range rows {
		roles, err := r.loadRoles(ctx, rows[i].ID)
		if err != nil {
			return model.UserQueryResult{}, err
		}
		data = append(data, *toOpemUser(&rows[i], roles))
	}
	return model.UserQueryResult{Records: total, Data: data}, nil
}

func (r *sqlUserRepo) loadRoles(ctx context.Context, userID string) ([]sqlUserRole, error) {
	var roles []sqlUserRole
	err := r.db.NewSelect().Model(&roles).Where("user_id = ?", userID).Scan(ctx)
	if err != nil && !isNotFound(err) {
		return nil, err
	}
	return roles, nil
}
