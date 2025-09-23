package box

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type Status struct {
	NumIn  int32  `json:"num_in,omitempty" bson:"num_in,omitempty" yaml:"num_in,omitempty"`
	NumOut int32  `json:"num_out,omitempty" bson:"num_out,omitempty" yaml:"num_out,omitempty"`
	Status string `json:"status,omitempty" bson:"status,omitempty" yaml:"status,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Status) IsZero() bool {
	return s.NumIn == 0 && s.NumOut == 0 && s.Status == ""
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
