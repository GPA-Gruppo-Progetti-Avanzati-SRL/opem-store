package commons

import "go.mongodb.org/mongo-driver/v2/bson"

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type FileReference struct {
	OId    bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	SrcSet []FileVariant `json:"src_set,omitempty" bson:"src_set,omitempty" yaml:"src_set,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s FileReference) IsZero() bool {
	return s.OId == bson.NilObjectID && len(s.SrcSet) == 0
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
