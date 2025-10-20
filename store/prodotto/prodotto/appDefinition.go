package prodotto

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type AppDefinition struct {
	Func      string              `json:"func,omitempty" bson:"func,omitempty" yaml:"func,omitempty"`
	IsCustom  bool                `json:"is_custom,omitempty" bson:"is_custom,omitempty" yaml:"is_custom,omitempty"`
	Id        string              `json:"id,omitempty" bson:"id,omitempty" yaml:"id,omitempty"`
	Name      string              `json:"name,omitempty" bson:"name,omitempty" yaml:"name,omitempty"`
	Tag       string              `json:"tag,omitempty" bson:"tag,omitempty" yaml:"tag,omitempty"`
	AppNumber AppNumberDefinition `json:"app_number,omitempty" bson:"app_number,omitempty" yaml:"app_number,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s AppDefinition) IsZero() bool {
	return s.Func == "" && !s.IsCustom && s.Id == "" && s.Name == "" && s.Tag == "" && s.AppNumber.IsZero()
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
