package commons

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type BidTextPair struct {
	Bid  string `json:"bid,omitempty" bson:"bid,omitempty" yaml:"bid,omitempty"`
	Text string `json:"text,omitempty" bson:"text,omitempty" yaml:"text,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s BidTextPair) IsZero() bool {
	return s.Bid == "" && s.Text == ""
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
