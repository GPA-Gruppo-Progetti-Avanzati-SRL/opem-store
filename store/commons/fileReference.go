package commons

import "go.mongodb.org/mongo-driver/bson/primitive"

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type FileReference struct {
	OId    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	SrcSet []FileVariant      `json:"srcSet,omitempty" bson:"srcSet,omitempty" yaml:"srcSet,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s FileReference) IsZero() bool {
	return s.OId == primitive.NilObjectID && len(s.SrcSet) == 0
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
