package domain

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type Member struct {
	Code    string `json:"code,omitempty" bson:"code,omitempty" yaml:"code,omitempty"`
	ObjType string `json:"obj_type,omitempty" bson:"obj_type,omitempty" yaml:"obj_type,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Member) IsZero() bool {
	return s.Code == "" && s.ObjType == ""
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
