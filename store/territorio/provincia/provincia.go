package provincia

import "go.mongodb.org/mongo-driver/bson/primitive"
import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

// @tpm-schematics:start-region("top-file-section")
const (
	EntityType   = "PROVINCIA"
	CollectionId = "trt_territorio"
)

// @tpm-schematics:end-region("top-file-section")

type Provincia struct {
	OId            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Et             string             `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Code           string             `json:"code,omitempty" bson:"code,omitempty" yaml:"code,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty" yaml:"name,omitempty"`
	CodeUicNazione string             `json:"code_uic_nazione,omitempty" bson:"code_uic_nazione,omitempty" yaml:"code_uic_nazione,omitempty"`
	SysInfo        commons.SysInfo    `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Provincia) IsZero() bool {
	return s.OId == primitive.NilObjectID && s.Et == "" && s.Code == "" && s.Name == "" && s.CodeUicNazione == "" && s.SysInfo.IsZero()
}

type QueryResult struct {
	Records int         `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []Provincia `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
