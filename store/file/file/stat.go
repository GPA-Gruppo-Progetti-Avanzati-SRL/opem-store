package file

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type Stat struct {
	NumRecords      int32 `json:"num_records,omitempty" bson:"num_records,omitempty" yaml:"num_records,omitempty"`
	NumErrors       int32 `json:"num_errors,omitempty" bson:"num_errors,omitempty" yaml:"num_errors,omitempty"`
	NumRecordsOk    int32 `json:"num_records_ok,omitempty" bson:"num_records_ok,omitempty" yaml:"num_records_ok,omitempty"`
	NumRecordsFixed int32 `json:"num_records_fixed,omitempty" bson:"num_records_fixed,omitempty" yaml:"num_records_fixed,omitempty"`
	Size            int32 `json:"size,omitempty" bson:"size,omitempty" yaml:"size,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Stat) IsZero() bool {
	return s.NumRecords == 0 && s.NumErrors == 0 && s.NumRecordsOk == 0 && s.NumRecordsFixed == 0 && s.Size == 0
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
