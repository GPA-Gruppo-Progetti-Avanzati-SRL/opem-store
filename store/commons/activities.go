package commons

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type Activities struct {
	Todos []Activity `json:"todos,omitempty" bson:"todos,omitempty" yaml:"todos,omitempty"`
	Logs  []Activity `json:"logs,omitempty" bson:"logs,omitempty" yaml:"logs,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Activities) IsZero() bool {
	return len(s.Todos) == 0 && len(s.Logs) == 0
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
