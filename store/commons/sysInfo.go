package commons

import "go.mongodb.org/mongo-driver/v2/bson"

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type SysInfo struct {
	Status     string        `json:"status,omitempty" bson:"status,omitempty" yaml:"status,omitempty"`
	CreatedAt  bson.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty" yaml:"created_at,omitempty"`
	ModifiedAt bson.DateTime `json:"modified_at,omitempty" bson:"modified_at,omitempty" yaml:"modified_at,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s SysInfo) IsZero() bool {
	return s.Status == "" && s.CreatedAt == 0 && s.ModifiedAt == 0
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
