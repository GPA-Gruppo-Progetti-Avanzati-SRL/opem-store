package card

import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type CardApp struct {
	AppType   string                `json:"app_type,omitempty" bson:"app_type,omitempty" yaml:"app_type,omitempty"`
	AppNumber commons.ValueTextPair `json:"app_number,omitempty" bson:"app_number,omitempty" yaml:"app_number,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s CardApp) IsZero() bool {
	return s.AppType == "" && s.AppNumber.IsZero()
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
