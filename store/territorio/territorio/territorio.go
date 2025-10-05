package territorio

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/territorio/comune"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/territorio/nazione"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/territorio/provincia"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

// @tpm-schematics:start-region("top-file-section")

var EntityTypes = []string{nazione.EntityType, provincia.EntityType, comune.EntityType}

const (
	CollectionId = "trt_territorio"
)

// @tpm-schematics:end-region("top-file-section")

type Territorio struct {
	OId           primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Et            string              `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Bid           string              `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Name          string              `json:"name,omitempty" bson:"name,omitempty" yaml:"name,omitempty"`
	CodeUic       string              `json:"code_uic,omitempty" bson:"code_uic,omitempty" yaml:"code_uic,omitempty"`
	CodeIso3      string              `json:"code_iso3,omitempty" bson:"code_iso3,omitempty" yaml:"code_iso3,omitempty"`
	CodeCatastale string              `json:"code_catastale,omitempty" bson:"code_catastale,omitempty" yaml:"code_catastale,omitempty"`
	Nazione       commons.BidTextPair `json:"nazione,omitempty" bson:"nazione,omitempty" yaml:"nazione,omitempty"`
	Cap1          string              `json:"cap1,omitempty" bson:"cap1,omitempty" yaml:"cap1,omitempty"`
	Cap2          string              `json:"cap2,omitempty" bson:"cap2,omitempty" yaml:"cap2,omitempty"`
	Provincia     commons.BidTextPair `json:"provincia,omitempty" bson:"provincia,omitempty" yaml:"provincia,omitempty"`
	CodeIstat     string              `json:"code_istat,omitempty" bson:"code_istat,omitempty" yaml:"code_istat,omitempty"`
	Cab           string              `json:"cab,omitempty" bson:"cab,omitempty" yaml:"cab,omitempty"`
	SysInfo       commons.SysInfo     `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Territorio) IsZero() bool {
	return s.OId == primitive.NilObjectID && s.Et == "" && s.Bid == "" && s.Name == "" && s.CodeUic == "" && s.CodeIso3 == "" && s.CodeCatastale == "" && s.Nazione.IsZero() && s.Cap1 == "" && s.Cap2 == "" && s.Provincia.IsZero() && s.CodeIstat == "" && s.Cab == "" && s.SysInfo.IsZero()
}

type QueryResult struct {
	Records int          `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []Territorio `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
