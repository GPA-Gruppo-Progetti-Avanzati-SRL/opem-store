package commons

import "encoding/json"

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

func (s ValueRange) String() string {
	b, err := json.Marshal(s)
	if err != nil {
		return err.Error()
	}

	return string(b)
}

// @tpm-schematics:end-region("bottom-file-section")
