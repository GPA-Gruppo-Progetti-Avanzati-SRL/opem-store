package commons

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type Address struct {
	Type     string      `json:"type,omitempty" bson:"type,omitempty" yaml:"type,omitempty"`
	County   BidTextPair `json:"county,omitempty" bson:"county,omitempty" yaml:"county,omitempty"`
	Townhall BidTextPair `json:"townhall,omitempty" bson:"townhall,omitempty" yaml:"townhall,omitempty"`
	Zipcode  string      `json:"zipcode,omitempty" bson:"zipcode,omitempty" yaml:"zipcode,omitempty"`
	AttnTo   string      `json:"attn_to,omitempty" bson:"attn_to,omitempty" yaml:"attn_to,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Address) IsZero() bool {
	return s.Type == "" && s.County.IsZero() && s.Townhall.IsZero() && s.Zipcode == "" && s.AttnTo == ""
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
