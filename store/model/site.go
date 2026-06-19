package model

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
	"gopkg.in/yaml.v3"
)

type Site struct {
	OId         string          `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Bid         string          `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Domain      string          `json:"domain,omitempty" bson:"domain,omitempty" yaml:"domain,omitempty"`
	Et          string          `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Name        string          `json:"name,omitempty" bson:"name,omitempty" yaml:"name,omitempty"`
	Description string          `json:"description,omitempty" bson:"description,omitempty" yaml:"description,omitempty"`
	Icon        string          `json:"icon,omitempty" bson:"icon,omitempty" yaml:"icon,omitempty"`
	Order       int             `json:"order,omitempty" bson:"order,omitempty" yaml:"order,omitempty"`
	Bookmark    bool            `json:"bookmark,omitempty" bson:"bookmark,omitempty" yaml:"bookmark,omitempty"`
	Langs       string          `json:"langs,omitempty" bson:"langs,omitempty" yaml:"langs,omitempty"`
	Apps        []commons.App   `json:"apps,omitempty" bson:"apps,omitempty" yaml:"apps,omitempty"`
	SysInfo     commons.SysInfo `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`
}

func (s *Site) UnmarshalYAML(value *yaml.Node) error {
	type rawSite Site
	var raw rawSite
	if err := value.Decode(&raw); err != nil {
		return err
	}
	*s = Site(raw)
	if s.OId == "" {
		s.OId = newOID()
	}
	return nil
}

func (s Site) IsZero() bool {
	return s.OId == "" && s.Bid == "" && s.Domain == "" && s.Et == "" && s.Name == "" && s.Description == "" && !s.Bookmark && s.Langs == "" && len(s.Apps) == 0 && s.SysInfo.IsZero()
}

func (s Site) GetAppByObjTypeAndId(objType commons.AppObjType, appId string) (commons.App, bool) {
	app := commons.App{Id: appId}
	for _, a := range s.Apps {
		if a.Id == appId {
			app = a
			if a.ObjType == string(objType) {
				return a, true
			}
		}
	}
	return app, false
}

type SiteQueryResult struct {
	Records int    `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []Site `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

type SiteQueryOptions struct {
	Limit         int64  `form:"limit"         query:"limit"         json:"limit,omitempty"         bson:"limit,omitempty"         yaml:"limit,omitempty"`
	Offset        int64  `form:"offset"        query:"offset"        json:"offset,omitempty"        bson:"offset,omitempty"        yaml:"offset,omitempty"`
	SortBy        string `form:"sortBy"        query:"sortBy"        json:"sortBy,omitempty"        bson:"sortBy,omitempty"        yaml:"sortBy,omitempty"`
	SortDirection string `form:"sortDirection" query:"sortDirection" json:"sortDirection,omitempty" bson:"sortDirection,omitempty" yaml:"sortDirection,omitempty"`
	SearchTerm    string `form:"ssearch"       query:"ssearch"       json:"ssearch,omitempty"       bson:"ssearch,omitempty"       yaml:"ssearch,omitempty"`
	WithCount     bool   `form:"withCount"     query:"withCount"     json:"withCount,omitempty"     bson:"withCount,omitempty"     yaml:"withCount,omitempty"`
}
