package site

import "go.mongodb.org/mongo-driver/v2/bson"
import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

// @tpm-schematics:start-region("top-file-section")
const (
	EntityType   = "site"
	CollectionId = "site"
)

// @tpm-schematics:end-region("top-file-section")

type Site struct {
	OId         bson.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Bid         string          `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Domain      string          `json:"domain,omitempty" bson:"domain,omitempty" yaml:"domain,omitempty"`
	Et          string          `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Name        string          `json:"name,omitempty" bson:"name,omitempty" yaml:"name,omitempty"`
	Description string          `json:"description,omitempty" bson:"description,omitempty" yaml:"description,omitempty"`
	Bookmark    bool            `json:"bookmark,omitempty" bson:"bookmark,omitempty" yaml:"bookmark,omitempty"`
	Langs       string          `json:"langs,omitempty" bson:"langs,omitempty" yaml:"langs,omitempty"`
	Apps        []commons.App   `json:"apps,omitempty" bson:"apps,omitempty" yaml:"apps,omitempty"`
	SysInfo     commons.SysInfo `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Site) IsZero() bool {
	return s.OId == bson.NilObjectID && s.Bid == "" && s.Domain == "" && s.Et == "" && s.Name == "" && s.Description == "" && !s.Bookmark && s.Langs == "" && len(s.Apps) == 0 && s.SysInfo.IsZero()
}

type QueryResult struct {
	Records int    `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []Site `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

// @tpm-schematics:start-region("bottom-file-section")

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

// @tpm-schematics:end-region("bottom-file-section")
