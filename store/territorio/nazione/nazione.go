package nazione

import "go.mongodb.org/mongo-driver/v2/bson"
import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

// @tpm-schematics:start-region("top-file-section")
const (
	EntityType   = "country"
	CollectionId = "territorio"

	CodeUicITA = "086"
	EE         = "EE"
)

// @tpm-schematics:end-region("top-file-section")

type Nazione struct {
	OId           bson.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Et            string          `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Bid           string          `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Name          string          `json:"name,omitempty" bson:"name,omitempty" yaml:"name,omitempty"`
	CodeUic       string          `json:"code_uic,omitempty" bson:"code_uic,omitempty" yaml:"code_uic,omitempty"`
	CodeIso3      string          `json:"code_iso3,omitempty" bson:"code_iso3,omitempty" yaml:"code_iso3,omitempty"`
	CodeCatastale string          `json:"code_catastale,omitempty" bson:"code_catastale,omitempty" yaml:"code_catastale,omitempty"`
	SysInfo       commons.SysInfo `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Nazione) IsZero() bool {
	return s.OId == bson.NilObjectID && s.Et == "" && s.Bid == "" && s.Name == "" && s.CodeUic == "" && s.CodeIso3 == "" && s.CodeCatastale == "" && s.SysInfo.IsZero()
}

type QueryResult struct {
	Records int       `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []Nazione `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
