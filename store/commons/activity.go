package commons

// @tpm-schematics:start-region("top-file-section")

const (
	ActivityStatusToDo = "todo"
	ActivityStatusDone = "done"
)

// @tpm-schematics:end-region("top-file-section")

type Activity struct {
	Id     string `json:"id,omitempty" bson:"id,omitempty" yaml:"id,omitempty"`
	Status string `json:"status,omitempty" bson:"status,omitempty" yaml:"status,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Activity) IsZero() bool {
	return s.Id == "" && s.Status == ""
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
