package file

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type EntRefStruct struct {
	Dom     string `json:"dom,omitempty" bson:"dom,omitempty" yaml:"dom,omitempty"`
	Ns      string `json:"ns,omitempty" bson:"ns,omitempty" yaml:"ns,omitempty"`
	EntType string `json:"ent_type,omitempty" bson:"ent_type,omitempty" yaml:"ent_type,omitempty"`
	EntId   string `json:"ent_id,omitempty" bson:"ent_id,omitempty" yaml:"ent_id,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s EntRefStruct) IsZero() bool {
	return s.Dom == "" && s.Ns == "" && s.EntType == "" && s.EntId == ""
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
