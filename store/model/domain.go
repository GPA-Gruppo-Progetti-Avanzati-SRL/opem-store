package model

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
	"gopkg.in/yaml.v3"
)

type Member struct {
	Code    string `json:"code,omitempty" bson:"code,omitempty" yaml:"code,omitempty"`
	ObjType string `json:"obj_type,omitempty" bson:"obj_type,omitempty" yaml:"obj_type,omitempty"`
}

func (s Member) IsZero() bool {
	return s.Code == "" && s.ObjType == ""
}

type Domain struct {
	OId         string          `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Bid         string          `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Et          string          `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Name        string          `json:"name,omitempty" bson:"name,omitempty" yaml:"name,omitempty"`
	Description string          `json:"description,omitempty" bson:"description,omitempty" yaml:"description,omitempty"`
	LogoUrl     string          `json:"logo_url,omitempty" bson:"logo_url,omitempty" yaml:"logo_url,omitempty"`
	Langs       string          `json:"langs,omitempty" bson:"langs,omitempty" yaml:"langs,omitempty"`
	Members     []Member        `json:"members,omitempty" bson:"members,omitempty" yaml:"members,omitempty"`
	Apps        []commons.App   `json:"apps,omitempty" bson:"apps,omitempty" yaml:"apps,omitempty"`
	SysInfo     commons.SysInfo `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`
}

func (d *Domain) UnmarshalYAML(value *yaml.Node) error {
	type rawDomain Domain
	var raw rawDomain
	if err := value.Decode(&raw); err != nil {
		return err
	}
	*d = Domain(raw)
	if d.OId == "" {
		d.OId = newOID()
	}
	return nil
}

func (s Domain) IsZero() bool {
	return s.OId == "" && s.Bid == "" && s.Et == "" && s.Name == "" && s.Description == "" && s.LogoUrl == "" && s.Langs == "" && len(s.Members) == 0 && len(s.Apps) == 0 && s.SysInfo.IsZero()
}

func (s Domain) GetAppByObjTypeAndId(objType commons.AppObjType, appId string) (commons.App, bool) {
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

type DomainQueryResult struct {
	Records int      `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []Domain `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}
