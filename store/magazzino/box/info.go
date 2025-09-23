package box

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type Info struct {
	BoxCode     string `json:"box_code,omitempty" bson:"box_code,omitempty" yaml:"box_code,omitempty"`
	PackageCode string `json:"package_code,omitempty" bson:"package_code,omitempty" yaml:"package_code,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Info) IsZero() bool {
	return s.BoxCode == "" && s.PackageCode == ""
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
