package commons

import "go.mongodb.org/mongo-driver/v2/bson"

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type Event struct {
	Type string        `json:"type,omitempty" bson:"type,omitempty" yaml:"type,omitempty"`
	Who  string        `json:"who,omitempty" bson:"who,omitempty" yaml:"who,omitempty"`
	Text string        `json:"text,omitempty" bson:"text,omitempty" yaml:"text,omitempty"`
	When bson.DateTime `json:"when,omitempty" bson:"when,omitempty" yaml:"when,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Event) IsZero() bool {
	return s.Type == "" && s.Who == "" && s.Text == "" && s.When == 0
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
