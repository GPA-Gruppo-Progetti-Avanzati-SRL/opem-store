package keyvaluepackage

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type KeyValue struct {
	Key         string `json:"key,omitempty" bson:"key,omitempty" yaml:"key,omitempty"`
	Value       string `json:"value,omitempty" bson:"value,omitempty" yaml:"value,omitempty"`
	Order       int32  `json:"order,omitempty" bson:"order,omitempty" yaml:"order,omitempty"`
	Kind        string `json:"kind,omitempty" bson:"kind,omitempty" yaml:"kind,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty" yaml:"name,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty" yaml:"description,omitempty"`
	Hint        string `json:"hint,omitempty" bson:"hint,omitempty" yaml:"hint,omitempty"`
	Icon        string `json:"icon,omitempty" bson:"icon,omitempty" yaml:"icon,omitempty"`
	Status      string `json:"status,omitempty" bson:"status,omitempty" yaml:"status,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s KeyValue) IsZero() bool {
	return s.Key == "" && s.Value == "" && s.Order == 0 && s.Kind == "" && s.Name == "" && s.Description == "" && s.Hint == "" && s.Icon == "" && s.Status == ""
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
