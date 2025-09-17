package nazione

import "go.mongodb.org/mongo-driver/bson/primitive"

// @tpm-schematics:start-region("top-file-section")
const (
	EntityType   = "NAZIONE"
	CollectionId = "trt_territorio"
)

// @tpm-schematics:end-region("top-file-section")

type Nazione struct {
	OId           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Et            string             `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Code          string             `json:"code,omitempty" bson:"code,omitempty" yaml:"code,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty" yaml:"name,omitempty"`
	CodeUic       string             `json:"code_uic,omitempty" bson:"code_uic,omitempty" yaml:"code_uic,omitempty"`
	CodeIso3      string             `json:"code_iso3,omitempty" bson:"code_iso3,omitempty" yaml:"code_iso3,omitempty"`
	CodeCatastale string             `json:"code_catastale,omitempty" bson:"code_catastale,omitempty" yaml:"code_catastale,omitempty"`
	Status        string             `json:"status,omitempty" bson:"status,omitempty" yaml:"status,omitempty"`
	Order         int32              `json:"order,omitempty" bson:"order,omitempty" yaml:"order,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Nazione) IsZero() bool {
	return s.OId == primitive.NilObjectID && s.Et == "" && s.Code == "" && s.Name == "" && s.CodeUic == "" && s.CodeIso3 == "" && s.CodeCatastale == "" && s.Status == "" && s.Order == 0
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
