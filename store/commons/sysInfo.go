package commons

import "go.mongodb.org/mongo-driver/bson/primitive"

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type SysInfo struct {
	Status     string             `json:"status,omitempty" bson:"status,omitempty" yaml:"status,omitempty"`
	Createdat  primitive.DateTime `json:"createdat,omitempty" bson:"createdat,omitempty" yaml:"createdat,omitempty"`
	Modifiedat primitive.DateTime `json:"modifiedat,omitempty" bson:"modifiedat,omitempty" yaml:"modifiedat,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s SysInfo) IsZero() bool {
	return s.Status == "" && s.Createdat == 0 && s.Modifiedat == 0
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
