package keyvaluepackage

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type KeyValue struct {
	Key   string `json:"key,omitempty" bson:"key,omitempty" yaml:"key,omitempty"`
	Value string `json:"value,omitempty" bson:"value,omitempty" yaml:"value,omitempty"`
	Order int32  `json:"order,omitempty" bson:"order,omitempty" yaml:"order,omitempty"`
	Kind  string `json:"kind,omitempty" bson:"kind,omitempty" yaml:"kind,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s KeyValue) IsZero() bool {
	return s.Key == "" && s.Value == "" && s.Order == 0 && s.Kind == ""
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
