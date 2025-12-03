package commons

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type StatusCodeTextPair struct {
	Code string `json:"code,omitempty" bson:"code,omitempty" yaml:"code,omitempty"`
	Text string `json:"text,omitempty" bson:"text,omitempty" yaml:"text,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s StatusCodeTextPair) IsZero() bool {
	return s.Code == "" && s.Text == ""
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
