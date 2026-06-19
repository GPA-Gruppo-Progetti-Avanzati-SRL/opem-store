package sql

import (
	model "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/model"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
)

func newID() string {
	b := make([]byte, 12)
	if _, err := rand.Read(b); err != nil {
		panic("sqllks: newID: " + err.Error())
	}
	return hex.EncodeToString(b)
}

func sysInfo(status string, createdAt, modifiedAt time.Time) commons.SysInfo {
	return commons.SysInfo{
		Status:     status,
		CreatedAt:  createdAt,
		ModifiedAt: modifiedAt,
	}
}

// ── User ──────────────────────────────────────────────────────────────────────

func toOpemUser(row *sqlUser, roles []sqlUserRole) *model.User {
	opemRoles := make([]commons.UserRole, 0, len(roles))
	for _, r := range roles {
		opemRoles = append(opemRoles, commons.UserRole{
			Domain: r.Domain,
			Site:   r.Site,
			Apps:   r.Apps,
		})
	}
	et := row.Et
	if et == "" {
		et = "user"
	}
	return &model.User{
		OId:       row.ID,
		Et:        et,
		Nickname:  row.Nickname,
		Firstname: row.Firstname,
		Lastname:  row.Lastname,
		Email:     row.Email,
		Password:  row.Password,
		Roles:     opemRoles,
		SysInfo:   sysInfo(row.Status, row.CreatedAt, row.ModifiedAt),
	}
}

func fromOpemUser(u *model.User) (*sqlUser, []sqlUserRole) {
	id := u.OId
	if id == "" {
		id = newID()
	}
	et := u.Et
	if et == "" {
		et = "user"
	}
	now := time.Now()
	row := &sqlUser{
		ID:         id,
		Et:         et,
		Nickname:   u.Nickname,
		Firstname:  u.Firstname,
		Lastname:   u.Lastname,
		Email:      u.Email,
		Password:   u.Password,
		Status:     u.SysInfo.Status,
		CreatedAt:  now,
		ModifiedAt: now,
	}
	if row.Status == "" {
		row.Status = "active"
	}
	roles := make([]sqlUserRole, 0, len(u.Roles))
	for _, r := range u.Roles {
		roles = append(roles, sqlUserRole{
			ID:     newID(),
			UserID: id,
			Domain: r.Domain,
			Site:   r.Site,
			Apps:   r.Apps,
		})
	}
	return row, roles
}

// ── Session ───────────────────────────────────────────────────────────────────

func toOpemSession(row *sqlSession) *model.Session {
	return &model.Session{
		OId:        row.ID,
		Userid:     row.Userid,
		Nickname:   row.Nickname,
		RemoteAddr: row.RemoteAddr,
		Flags:      row.Flags,
		SysInfo:    sysInfo(row.Status, row.CreatedAt, row.ModifiedAt),
	}
}

func fromOpemSession(s *model.Session) *sqlSession {
	id := s.OId
	if id == "" {
		id = newID()
	}
	now := time.Now()
	return &sqlSession{
		ID:         id,
		Userid:     s.Userid,
		Nickname:   s.Nickname,
		RemoteAddr: s.RemoteAddr,
		Flags:      s.Flags,
		Status:     s.SysInfo.Status,
		CreatedAt:  now,
		ModifiedAt: now,
	}
}

// ── Domain ────────────────────────────────────────────────────────────────────

func toOpemDomain(row *sqlDomain, apps []sqlApp) *model.Domain {
	opemApps := sqlAppsToCommons(apps)
	return &model.Domain{
		OId:         row.ID,
		Bid:         row.Code,
		Name:        row.Name,
		Description: row.Description,
		LogoUrl:     row.LogoUrl,
		Langs:       row.Langs,
		Apps:        opemApps,
		SysInfo:     sysInfo(row.Status, row.CreatedAt, row.ModifiedAt),
	}
}

func fromOpemDomain(d *model.Domain) (*sqlDomain, []sqlApp) {
	id := d.OId
	if id == "" {
		id = newID()
	}
	now := time.Now()
	row := &sqlDomain{
		ID:          id,
		Code:        d.Bid,
		Name:        d.Name,
		Description: d.Description,
		LogoUrl:     d.LogoUrl,
		Langs:       d.Langs,
		Status:      d.SysInfo.Status,
		CreatedAt:   now,
		ModifiedAt:  now,
	}
	if row.Status == "" {
		row.Status = "active"
	}
	apps := commonsAppsToSQL(id, "domain", d.Apps)
	return row, apps
}

// ── Site ──────────────────────────────────────────────────────────────────────

func toOpemSite(row *sqlSite, apps []sqlApp) *model.Site {
	opemApps := sqlAppsToCommons(apps)
	return &model.Site{
		OId:         row.ID,
		Bid:         row.Code,
		Domain:      row.DomainCode,
		Name:        row.Name,
		Description: row.Description,
		Icon:        row.Icon,
		Order:       row.Order,
		Bookmark:    row.Bookmark,
		Langs:       row.Langs,
		Apps:        opemApps,
		SysInfo:     sysInfo(row.Status, row.CreatedAt, row.ModifiedAt),
	}
}

func fromOpemSite(s *model.Site) (*sqlSite, []sqlApp) {
	id := s.OId
	if id == "" {
		id = newID()
	}
	now := time.Now()
	row := &sqlSite{
		ID:          id,
		Code:        s.Bid,
		DomainCode:  s.Domain,
		Name:        s.Name,
		Description: s.Description,
		Icon:        s.Icon,
		Order:       s.Order,
		Bookmark:    s.Bookmark,
		Langs:       s.Langs,
		Status:      s.SysInfo.Status,
		CreatedAt:   now,
		ModifiedAt:  now,
	}
	if row.Status == "" {
		row.Status = "active"
	}
	apps := commonsAppsToSQL(id, "site", s.Apps)
	return row, apps
}

// ── App helpers ───────────────────────────────────────────────────────────────

func sqlAppsToCommons(rows []sqlApp) []commons.App {
	out := make([]commons.App, 0, len(rows))
	for _, r := range rows {
		out = append(out, commons.App{
			Id:           r.AppID,
			ObjType:      r.ObjType,
			Name:         r.Name,
			Title:        r.Title,
			SubTitle:     r.SubTitle,
			Description:  r.Description,
			Path:         r.Path,
			Version:      r.Version,
			RoleRequired: r.RoleRequired,
			Icon:         r.Icon,
			Order:        r.Order,
			ApiContext:   r.ApiContext,
		})
	}
	return out
}

func commonsAppsToSQL(ownerID, ownerType string, apps []commons.App) []sqlApp {
	out := make([]sqlApp, 0, len(apps))
	for _, a := range apps {
		out = append(out, sqlApp{
			ID:           newID(),
			OwnerType:    ownerType,
			OwnerID:      ownerID,
			AppID:        a.Id,
			ObjType:      a.ObjType,
			Name:         a.Name,
			Title:        a.Title,
			SubTitle:     a.SubTitle,
			Description:  a.Description,
			Path:         a.Path,
			Version:      a.Version,
			RoleRequired: a.RoleRequired,
			Icon:         a.Icon,
			Order:        a.Order,
			ApiContext:   a.ApiContext,
		})
	}
	return out
}

// ── KV ────────────────────────────────────────────────────────────────────────

func toOpemKV(row *sqlKV, props []sqlKVProperty) model.KeyValuePackage {
	opemProps := make([]model.KeyValue, 0, len(props))
	for _, p := range props {
		opemProps = append(opemProps, model.KeyValue{
			Key:         p.Key,
			Value:       p.Value,
			Kind:        p.Kind,
			Order:       p.Order,
			Name:        p.Name,
			Description: p.Description,
			Hint:        p.Hint,
			Icon:        p.Icon,
			Status:      p.Status,
			SysName:     p.SysName,
		})
	}
	return model.KeyValuePackage{
		OId:         row.ID,
		Bid:         row.Bid,
		Scope:       row.Scope,
		Category:    row.Category,
		IsSystem:    row.IsSystem,
		Description: row.Description,
		Properties:  opemProps,
		SysInfo:     commons.SysInfo{Status: row.Status},
	}
}

func sqlKVPropsToKeyValues(props []sqlKVProperty) []model.KeyValue {
	out := make([]model.KeyValue, 0, len(props))
	for _, p := range props {
		out = append(out, model.KeyValue{
			Key:         p.Key,
			Value:       p.Value,
			Kind:        p.Kind,
			Order:       p.Order,
			Name:        p.Name,
			Description: p.Description,
			Hint:        p.Hint,
			Icon:        p.Icon,
			Status:      p.Status,
			SysName:     p.SysName,
		})
	}
	return out
}

// ── Sequence ──────────────────────────────────────────────────────────────────

func toOpemSequence(row *sqlSequence) model.Sequence {
	return model.Sequence{
		OId:    row.ID,
		Bid:    row.Bid,
		Domain: row.DomainCode,
		Site:   row.SiteCode,
		Value:  row.Value,
		Format: row.Format,
		Prefix: row.Prefix,
	}
}

// ── ACL ───────────────────────────────────────────────────────────────────────

func toOpemRoleCaps(row *sqlRoleCaps, capGroups []string, capabilities []string) model.RoleCapsEntry {
	return model.RoleCapsEntry{
		OId:          row.ID,
		Role:         row.Role,
		Domain:       row.DomainCode,
		Site:         row.SiteCode,
		App:          row.App,
		CapGroups:    capGroups,
		Capabilities: capabilities,
		SysInfo:      commons.SysInfo{Status: row.Status},
	}
}

func toOpemCapGroup(row *sqlCapGroup, capDefIDs []string) *model.CapGroup {
	return &model.CapGroup{
		OId:          row.ID,
		Description:  row.Description,
		Capabilities: capDefIDs,
		SysInfo:      commons.SysInfo{Status: row.Status},
	}
}

func toOpemCapDef(row *sqlCapDef) *model.CapDef {
	return &model.CapDef{
		OId:         row.ID,
		App:         row.App,
		Category:    row.Category,
		Description: row.Description,
		Name:        row.Name,
		Endpoint:    row.Endpoint,
		Icon:        row.Icon,
		Order:       row.Order,
		Menu:        row.Menu,
		Method:      row.Method,
		SysInfo:     commons.SysInfo{Status: row.Status},
	}
}

// ── Scope KV ──────────────────────────────────────────────────────────────────

func scopeFromDomainSite(domainCode, siteCode string) string {
	if domainCode == "" || domainCode == "root" {
		return "root"
	}
	if siteCode == "" || siteCode == "*" {
		return fmt.Sprintf("root/%s", domainCode)
	}
	return fmt.Sprintf("root/%s/%s", domainCode, siteCode)
}

func scopesForQuery(domainCode, siteCode string) []string {
	if domainCode == "" || domainCode == "root" {
		return []string{"root"}
	}
	if siteCode == "" || siteCode == "*" {
		return []string{"root", fmt.Sprintf("root/%s", domainCode)}
	}
	return []string{
		"root",
		fmt.Sprintf("root/%s", domainCode),
		fmt.Sprintf("root/%s/%s", domainCode, siteCode),
	}
}
