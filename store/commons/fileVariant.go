package commons

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type FileVariant struct {
	Ct   string `json:"ct,omitempty" bson:"ct,omitempty" yaml:"ct,omitempty"`
	Wd   int32  `json:"wd,omitempty" bson:"wd,omitempty" yaml:"wd,omitempty"`
	Ht   int32  `json:"ht,omitempty" bson:"ht,omitempty" yaml:"ht,omitempty"`
	Lks  string `json:"lks,omitempty" bson:"lks,omitempty" yaml:"lks,omitempty"`
	Bln  string `json:"bln,omitempty" bson:"bln,omitempty" yaml:"bln,omitempty"`
	Cnt  string `json:"cnt,omitempty" bson:"cnt,omitempty" yaml:"cnt,omitempty"`
	Url  string `json:"url,omitempty" bson:"url,omitempty" yaml:"url,omitempty"`
	Role string `json:"role,omitempty" bson:"role,omitempty" yaml:"role,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s FileVariant) IsZero() bool {
	return s.Ct == "" && s.Wd == 0 && s.Ht == 0 && s.Lks == "" && s.Bln == "" && s.Cnt == "" && s.Url == "" && s.Role == ""
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
