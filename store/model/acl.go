package model

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
)

// CapDef è la definizione canonica di una capability.
type CapDef struct {
	OId         string          `json:"_id,omitempty"         bson:"_id,omitempty"         yaml:"_id,omitempty"`
	Et          string          `json:"_et,omitempty"         bson:"_et,omitempty"         yaml:"_et,omitempty"`
	App         string          `json:"app"                   bson:"app"                   yaml:"app,omitempty"`
	Category    string          `json:"category"              bson:"category"              yaml:"category,omitempty"`
	SysInfo     commons.SysInfo `json:"sys_info,omitempty"    bson:"sys_info,omitempty"    yaml:"sys_info,omitempty"`
	Description string          `json:"description,omitempty" bson:"description,omitempty" yaml:"description,omitempty"`
	Name        string          `json:"name,omitempty"        bson:"name,omitempty"        yaml:"name,omitempty"`
	Endpoint    string          `json:"endpoint,omitempty"    bson:"endpoint,omitempty"    yaml:"endpoint,omitempty"`
	Icon        string          `json:"icon,omitempty"        bson:"icon,omitempty"        yaml:"icon,omitempty"`
	Order       int             `json:"order,omitempty"       bson:"order,omitempty"       yaml:"order,omitempty"`
	Menu        bool            `json:"menu"                  bson:"menu"                  yaml:"menu,omitempty"`
	Method      string          `json:"method,omitempty"      bson:"method,omitempty"      yaml:"method,omitempty"`
	Deny        bool            `json:"deny,omitempty"        bson:"deny,omitempty"        yaml:"deny,omitempty"`
}

func (c CapDef) IsZero() bool {
	return c.OId == "" && c.App == ""
}

// CapGroup raggruppa IDs di cap-def per riuso nelle assegnazioni role-caps.
type CapGroup struct {
	OId          string          `json:"_id,omitempty"         bson:"_id,omitempty"         yaml:"_id,omitempty"`
	Et           string          `json:"_et,omitempty"         bson:"_et,omitempty"         yaml:"_et,omitempty"`
	Description  string          `json:"description,omitempty" bson:"description,omitempty" yaml:"description,omitempty"`
	Capabilities []string        `json:"capabilities"          bson:"capabilities"          yaml:"capabilities,omitempty"`
	SysInfo      commons.SysInfo `json:"sys_info,omitempty"    bson:"sys_info,omitempty"    yaml:"sys_info,omitempty"`
}

func (c CapGroup) IsZero() bool {
	return c.OId == ""
}

// RoleCapsEntry assegna cap-groups e/o IDs individuali di cap-def a un ruolo.
type RoleCapsEntry struct {
	OId          string          `json:"_id,omitempty"      bson:"_id,omitempty"      yaml:"_id,omitempty"`
	Et           string          `json:"_et,omitempty"      bson:"_et,omitempty"      yaml:"_et,omitempty"`
	Role         string          `json:"role"               bson:"role"               yaml:"role,omitempty"`
	Domain       string          `json:"domain,omitempty"   bson:"domain,omitempty"   yaml:"domain,omitempty"`
	Site         string          `json:"site,omitempty"     bson:"site,omitempty"     yaml:"site,omitempty"`
	App          string          `json:"app,omitempty"      bson:"app,omitempty"      yaml:"app,omitempty"`
	CapGroups    []string        `json:"cap_groups"         bson:"cap_groups"         yaml:"cap_groups,omitempty"`
	Capabilities []string        `json:"capabilities"       bson:"capabilities"       yaml:"capabilities,omitempty"`
	SysInfo      commons.SysInfo `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`
}

func (r RoleCapsEntry) IsZero() bool {
	return r.Role == ""
}
