package model

import (
	"encoding/json"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

type User struct {
	OId            string                `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Nickname       string                `json:"nickname,omitempty" bson:"nickname,omitempty" yaml:"nickname,omitempty"`
	Et             string                `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Firstname      string                `json:"firstname,omitempty" bson:"firstname,omitempty" yaml:"firstname,omitempty"`
	Lastname       string                `json:"lastname,omitempty" bson:"lastname,omitempty" yaml:"lastname,omitempty"`
	Email          string                `json:"email,omitempty" bson:"email,omitempty" yaml:"email,omitempty"`
	Password       string                `json:"password,omitempty" bson:"password,omitempty" yaml:"password,omitempty"`
	Roles          []commons.UserRole    `json:"roles,omitempty" bson:"roles,omitempty" yaml:"roles,omitempty"`
	SysInfo        commons.SysInfo       `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`
	ProfilePicture commons.FileReference `json:"profile_picture,omitempty" bson:"profile_picture,omitempty" yaml:"profile_picture,omitempty"`
}

func (u *User) UnmarshalYAML(value *yaml.Node) error {
	type rawUser User
	var raw rawUser
	if err := value.Decode(&raw); err != nil {
		return err
	}
	*u = User(raw)
	if u.OId == "" {
		u.OId = newOID()
	}
	return nil
}

func (s User) IsZero() bool {
	return s.OId == "" && s.Nickname == "" && s.Et == "" && s.Firstname == "" && s.Lastname == "" && s.Email == "" && s.Password == "" && len(s.Roles) == 0 && s.SysInfo.IsZero() && s.ProfilePicture.IsZero()
}

func (s User) ToJson() ([]byte, error) {
	const semLogContext = "user::to-json"
	b, err := json.Marshal(s)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return nil, err
	}
	return b, nil
}

func (s User) HasRole4DomainSiteAppId(domain, site, appId, appType, role string) bool {
	return AnyRole4DomainSiteAppId(s.Roles, domain, site, appId, appType, role)
}

func AnyRole4DomainSiteAppId(roles []commons.UserRole, domain, site, appId, appType, role string) bool {
	const semLogContext = "user::any-role"
	for _, r := range roles {
		if !r.MatchDomainAndSite(domain, site) {
			continue
		}
		roleSet, err := r.ParseApps()
		if err != nil {
			log.Error().Err(err).Msg(semLogContext)
			continue
		}
		if roleSet.MatchRole(appId, appType, role) {
			return true
		}
	}
	return false
}

type UserQueryResult struct {
	Records int    `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []User `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

type FormResponseError struct {
	Field string `json:"field,omitempty" bson:"field,omitempty" yaml:"field,omitempty"`
	Error string `json:"message,omitempty" bson:"message,omitempty" yaml:"message,omitempty"`
}

type FormResponse struct {
	Status      int                 `json:"status,omitempty" bson:"status,omitempty" yaml:"status,omitempty"`
	Message     string              `json:"message,omitempty" bson:"message,omitempty" yaml:"message,omitempty"`
	FieldErrors []FormResponseError `json:"fieldErrors,omitempty" bson:"fieldErrors,omitempty" yaml:"fieldErrors,omitempty"`
	Document    *User               `json:"document,omitempty" bson:"document,omitempty" yaml:"document,omitempty"`
}

type UserQueryOptions struct {
	Limit         int64  `form:"limit"         query:"limit"         json:"limit,omitempty"         bson:"limit,omitempty"         yaml:"limit,omitempty"`
	Offset        int64  `form:"offset"        query:"offset"        json:"offset,omitempty"        bson:"offset,omitempty"        yaml:"offset,omitempty"`
	SortBy        string `form:"sortBy"        query:"sortBy"        json:"sortBy,omitempty"        bson:"sortBy,omitempty"        yaml:"sortBy,omitempty"`
	SortDirection string `form:"sortDirection" query:"sortDirection" json:"sortDirection,omitempty" bson:"sortDirection,omitempty" yaml:"sortDirection,omitempty"`
	SearchTerm    string `form:"ssearch"       query:"ssearch"       json:"ssearch,omitempty"       bson:"ssearch,omitempty"       yaml:"ssearch,omitempty"`
	WithCount     bool   `form:"withCount"     query:"withCount"     json:"withCount,omitempty"     bson:"withCount,omitempty"     yaml:"withCount,omitempty"`
}
