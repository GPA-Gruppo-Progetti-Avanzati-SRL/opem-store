package prodotto

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type AppNumberDefinition struct {
	Type string `json:"type,omitempty" bson:"type,omitempty" yaml:"type,omitempty"`
	Spec string `json:"spec,omitempty" bson:"spec,omitempty" yaml:"spec,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s AppNumberDefinition) IsZero() bool {
	return s.Type == "" && s.Spec == ""
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
