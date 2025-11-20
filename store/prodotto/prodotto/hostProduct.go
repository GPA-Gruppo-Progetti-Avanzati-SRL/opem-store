package prodotto

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type HostProduct struct {
	Code       string `json:"code,omitempty" bson:"code,omitempty" yaml:"code,omitempty"`
	LayoutCode string `json:"layout_code,omitempty" bson:"layout_code,omitempty" yaml:"layout_code,omitempty"`
	Release    string `json:"release,omitempty" bson:"release,omitempty" yaml:"release,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s HostProduct) IsZero() bool {
	return s.Code == "" && s.LayoutCode == "" && s.Release == ""
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
