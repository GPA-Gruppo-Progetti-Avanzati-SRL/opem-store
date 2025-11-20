package commons

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type Address struct {
	Type     string      `json:"type,omitempty" bson:"type,omitempty" yaml:"type,omitempty"`
	Country  BidTextPair `json:"country,omitempty" bson:"country,omitempty" yaml:"country,omitempty"`
	County   BidTextPair `json:"county,omitempty" bson:"county,omitempty" yaml:"county,omitempty"`
	Townhall BidTextPair `json:"townhall,omitempty" bson:"townhall,omitempty" yaml:"townhall,omitempty"`
	ZipCode  string      `json:"zip_code,omitempty" bson:"zip_code,omitempty" yaml:"zip_code,omitempty"`
	Address  string      `json:"address,omitempty" bson:"address,omitempty" yaml:"address,omitempty"`
	AttnTo   string      `json:"attn_to,omitempty" bson:"attn_to,omitempty" yaml:"attn_to,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Address) IsZero() bool {
	return s.Type == "" && s.Country.IsZero() && s.County.IsZero() && s.Townhall.IsZero() && s.ZipCode == "" && s.Address == "" && s.AttnTo == ""
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
