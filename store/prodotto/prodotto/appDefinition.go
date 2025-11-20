package prodotto

// @tpm-schematics:start-region("top-file-section")

const (
	XXT03_TAG = "XXT03"
	XXT06_TAG = "XXT06"
	XXT08_TAG = "XXT08"
)

// @tpm-schematics:end-region("top-file-section")

type AppDefinition struct {
	Func             string              `json:"func,omitempty" bson:"func,omitempty" yaml:"func,omitempty"`
	IsCustom         bool                `json:"is_custom,omitempty" bson:"is_custom,omitempty" yaml:"is_custom,omitempty"`
	Id               string              `json:"id,omitempty" bson:"id,omitempty" yaml:"id,omitempty"`
	Name             string              `json:"name,omitempty" bson:"name,omitempty" yaml:"name,omitempty"`
	Tag              string              `json:"tag,omitempty" bson:"tag,omitempty" yaml:"tag,omitempty"`
	AppNumber        AppNumberDefinition `json:"app_number,omitempty" bson:"app_number,omitempty" yaml:"app_number,omitempty"`
	ReportRange      bool                `json:"report_range,omitempty" bson:"report_range,omitempty" yaml:"report_range,omitempty"`
	ReportHolder     bool                `json:"report_holder,omitempty" bson:"report_holder,omitempty" yaml:"report_holder,omitempty"`
	IsEnvelopeNumber bool                `json:"is_envelope_number,omitempty" bson:"is_envelope_number,omitempty" yaml:"is_envelope_number,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s AppDefinition) IsZero() bool {
	return s.Func == "" && !s.IsCustom && s.Id == "" && s.Name == "" && s.Tag == "" && s.AppNumber.IsZero() && !s.ReportRange && !s.ReportHolder && !s.IsEnvelopeNumber
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
