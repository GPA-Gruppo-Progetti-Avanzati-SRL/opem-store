package domain

import "go.mongodb.org/mongo-driver/bson/primitive"
import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type Domain struct {
	OId         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Code        string             `json:"code,omitempty" bson:"code,omitempty" yaml:"code,omitempty"`
	ObjType     string             `json:"objType,omitempty" bson:"objType,omitempty" yaml:"objType,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty" yaml:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty" yaml:"description,omitempty"`
	Langs       string             `json:"langs,omitempty" bson:"langs,omitempty" yaml:"langs,omitempty"`
	Members     []Member           `json:"members,omitempty" bson:"members,omitempty" yaml:"members,omitempty"`
	Apps        []commons.App      `json:"apps,omitempty" bson:"apps,omitempty" yaml:"apps,omitempty"`
	SysInfo     commons.SysInfo    `json:"sysInfo,omitempty" bson:"sysInfo,omitempty" yaml:"sysInfo,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Domain) IsZero() bool {
	return s.OId == primitive.NilObjectID && s.Code == "" && s.ObjType == "" && s.Name == "" && s.Description == "" && s.Langs == "" && len(s.Members) == 0 && len(s.Apps) == 0 && s.SysInfo.IsZero()
}

// @tpm-schematics:start-region("bottom-file-section")

// GetAppByObjTypeAndId search for a match of App.Id and App.ObjType. If not found it returns an empty App
// or the last configured App with same Id (and different objType). It works in the assumption to redirect to
// the counter-part of the defined app.
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

type QueryResult struct {
	Records int      `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []Domain `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

// @tpm-schematics:end-region("bottom-file-section")
