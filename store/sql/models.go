package sql

import (
	"time"

	"github.com/uptrace/bun"
)

// ── opem_user ─────────────────────────────────────────────────────────────────

type sqlUser struct {
	bun.BaseModel `bun:"table:opem_user,alias:u"`

	ID         string    `bun:"id,pk"`
	Et         string    `bun:"et,notnull"`
	Nickname   string    `bun:"nickname,notnull"`
	Firstname  string    `bun:"firstname,notnull"`
	Lastname   string    `bun:"lastname,notnull"`
	Email      string    `bun:"email,notnull"`
	Password   string    `bun:"password,notnull"`
	Status     string    `bun:"status,notnull"`
	CreatedAt  time.Time `bun:"created_at,notnull"`
	ModifiedAt time.Time `bun:"modified_at,notnull"`
}

// ── opem_user_role ────────────────────────────────────────────────────────────

type sqlUserRole struct {
	bun.BaseModel `bun:"table:opem_user_role,alias:ur"`

	ID     string `bun:"id,pk"`
	UserID string `bun:"user_id,notnull"`
	Domain string `bun:"domain,notnull"`
	Site   string `bun:"site,notnull"`
	Apps   string `bun:"apps,notnull"`
}

// ── opem_session ──────────────────────────────────────────────────────────────

type sqlSession struct {
	bun.BaseModel `bun:"table:opem_session,alias:s"`

	ID         string    `bun:"id,pk"`
	Userid     string    `bun:"userid,notnull"`
	Nickname   string    `bun:"nickname,notnull"`
	RemoteAddr string    `bun:"remote_addr,notnull"`
	Flags      string    `bun:"flags,notnull"`
	Status     string    `bun:"status,notnull"`
	CreatedAt  time.Time `bun:"created_at,notnull"`
	ModifiedAt time.Time `bun:"modified_at,notnull"`
}

// ── opem_domain ───────────────────────────────────────────────────────────────

type sqlDomain struct {
	bun.BaseModel `bun:"table:opem_domain,alias:d"`

	ID          string    `bun:"id,pk"`
	Code        string    `bun:"code,notnull"`
	Name        string    `bun:"name,notnull"`
	Description string    `bun:"description,notnull"`
	LogoUrl     string    `bun:"logo_url,notnull"`
	Langs       string    `bun:"langs,notnull"`
	Status      string    `bun:"status,notnull"`
	CreatedAt   time.Time `bun:"created_at,notnull"`
	ModifiedAt  time.Time `bun:"modified_at,notnull"`
}

// ── opem_site ─────────────────────────────────────────────────────────────────

type sqlSite struct {
	bun.BaseModel `bun:"table:opem_site,alias:s"`

	ID          string    `bun:"id,pk"`
	Code        string    `bun:"code,notnull"`
	DomainCode  string    `bun:"domain_code,notnull"`
	Name        string    `bun:"name,notnull"`
	Description string    `bun:"description,notnull"`
	Icon        string    `bun:"icon,notnull"`
	Order       int       `bun:"ord,notnull"`
	Bookmark    bool      `bun:"bookmark,notnull"`
	Langs       string    `bun:"langs,notnull"`
	Status      string    `bun:"status,notnull"`
	CreatedAt   time.Time `bun:"created_at,notnull"`
	ModifiedAt  time.Time `bun:"modified_at,notnull"`
}

// ── opem_app ──────────────────────────────────────────────────────────────────

type sqlApp struct {
	bun.BaseModel `bun:"table:opem_app,alias:app"`

	ID           string `bun:"id,pk"`
	OwnerType    string `bun:"owner_type,notnull"`
	OwnerID      string `bun:"owner_id,notnull"`
	AppID        string `bun:"app_id,notnull"`
	ObjType      string `bun:"obj_type,notnull"`
	Name         string `bun:"name,notnull"`
	Title        string `bun:"title,notnull"`
	SubTitle     string `bun:"sub_title,notnull"`
	Description  string `bun:"description,notnull"`
	Path         string `bun:"path,notnull"`
	Version      string `bun:"version,notnull"`
	RoleRequired bool   `bun:"role_required,notnull"`
	Icon         string `bun:"icon,notnull"`
	Order        int    `bun:"ord,notnull"`
	ApiContext   string `bun:"api_context,notnull"`
}

// ── opem_kv ───────────────────────────────────────────────────────────────────

type sqlKV struct {
	bun.BaseModel `bun:"table:opem_kv,alias:kv"`

	ID          string `bun:"id,pk"`
	Bid         string `bun:"bid,notnull"`
	Scope       string `bun:"scope,notnull"`
	Category    string `bun:"category,notnull"`
	IsSystem    bool   `bun:"is_system,notnull"`
	Description string `bun:"description,notnull"`
	Status      string `bun:"status,notnull"`
}

// ── opem_kv_property ──────────────────────────────────────────────────────────

type sqlKVProperty struct {
	bun.BaseModel `bun:"table:opem_kv_property,alias:kvp"`

	ID          string `bun:"id,pk"`
	KVID        string `bun:"kv_id,notnull"`
	Key         string `bun:"key,notnull"`
	Value       string `bun:"value,notnull"`
	Kind        string `bun:"kind,notnull"`
	Order       int32  `bun:"ord,notnull"`
	Name        string `bun:"name,notnull"`
	Description string `bun:"description,notnull"`
	Hint        string `bun:"hint,notnull"`
	Icon        string `bun:"icon,notnull"`
	Status      string `bun:"status,notnull"`
	SysName     string `bun:"sys_name,notnull"`
}

// ── opem_sequence ─────────────────────────────────────────────────────────────

type sqlSequence struct {
	bun.BaseModel `bun:"table:opem_sequence,alias:seq"`

	ID         string `bun:"id,pk"`
	Bid        string `bun:"bid,notnull"`
	DomainCode string `bun:"domain_code,notnull"`
	SiteCode   string `bun:"site_code,notnull"`
	Value      int32  `bun:"val,notnull"`
	Format     string `bun:"format,notnull"`
	Prefix     string `bun:"prefix,notnull"`
}

// ── opem_acl_role_caps ────────────────────────────────────────────────────────

type sqlRoleCaps struct {
	bun.BaseModel `bun:"table:opem_acl_role_caps,alias:rc"`

	ID         string `bun:"id,pk"`
	Role       string `bun:"role,notnull"`
	DomainCode string `bun:"domain_code,notnull"`
	SiteCode   string `bun:"site_code,notnull"`
	App        string `bun:"app,notnull"`
	Status     string `bun:"status,notnull"`
}

// ── opem_acl_role_cap_group (junction) ────────────────────────────────────────

type sqlRoleCapGroup struct {
	bun.BaseModel `bun:"table:opem_acl_role_cap_group,alias:rcg"`

	RoleCapsID string `bun:"role_caps_id,pk"`
	CapGroupID string `bun:"cap_group_id,pk"`
}

// ── opem_acl_role_cap_def (junction diretta) ──────────────────────────────────

type sqlRoleCapDef struct {
	bun.BaseModel `bun:"table:opem_acl_role_cap_def,alias:rcd"`

	RoleCapsID string `bun:"role_caps_id,pk"`
	CapDefID   string `bun:"cap_def_id,pk"`
}

// ── opem_acl_cap_group ────────────────────────────────────────────────────────

type sqlCapGroup struct {
	bun.BaseModel `bun:"table:opem_acl_cap_group,alias:cg"`

	ID          string `bun:"id,pk"`
	Description string `bun:"description,notnull"`
	Status      string `bun:"status,notnull"`
}

// ── opem_acl_cap_group_def (junction) ─────────────────────────────────────────

type sqlCapGroupDef struct {
	bun.BaseModel `bun:"table:opem_acl_cap_group_def,alias:cgd"`

	CapGroupID string `bun:"cap_group_id,pk"`
	CapDefID   string `bun:"cap_def_id,pk"`
}

// ── opem_acl_cap_def ──────────────────────────────────────────────────────────

type sqlCapDef struct {
	bun.BaseModel `bun:"table:opem_acl_cap_def,alias:cd"`

	ID          string `bun:"id,pk"`
	App         string `bun:"app,notnull"`
	Category    string `bun:"category,notnull"`
	Description string `bun:"description,notnull"`
	Name        string `bun:"name,notnull"`
	Endpoint    string `bun:"endpoint,notnull"`
	Icon        string `bun:"icon,notnull"`
	Order       int    `bun:"ord,notnull"`
	Menu        bool   `bun:"menu,notnull"`
	Method      string `bun:"method,notnull"`
	Status      string `bun:"status,notnull"`
}
