package card

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type CardHolder struct {
	Bid            string `json:"bid,omitempty" bson:"bid,omitempty" yaml:"bid,omitempty"`
	EmbossingName  string `json:"embossing_name,omitempty" bson:"embossing_name,omitempty" yaml:"embossing_name,omitempty"`
	RegistrationId string `json:"registration_id,omitempty" bson:"registration_id,omitempty" yaml:"registration_id,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s CardHolder) IsZero() bool {
	return s.Bid == "" && s.EmbossingName == "" && s.RegistrationId == ""
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
