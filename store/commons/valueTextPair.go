package commons

import "strings"

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type ValueTextPair struct {
	Value string `json:"value,omitempty" bson:"value,omitempty" yaml:"value,omitempty"`
	Text  string `json:"text,omitempty" bson:"text,omitempty" yaml:"text,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s ValueTextPair) IsZero() bool {
	return s.Value == "" && s.Text == ""
}

// @tpm-schematics:start-region("bottom-file-section")

func NewValueTextPair(value string, maskValue string) ValueTextPair {
	if ndx := strings.LastIndex(value, maskValue); ndx >= 0 {
		return ValueTextPair{Value: value[ndx+1:], Text: value}
	}

	return ValueTextPair{Value: value}
}

// @tpm-schematics:end-region("bottom-file-section")
