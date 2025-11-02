package commons

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type ValueRange struct {
	From string `json:"from,omitempty" bson:"from,omitempty" yaml:"from,omitempty"`
	To   string `json:"to,omitempty" bson:"to,omitempty" yaml:"to,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s ValueRange) IsZero() bool {
	return s.From == "" && s.To == ""
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
