package commons

import "time"

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type SysInfo struct {
	Status     string    `json:"status,omitempty" bson:"status,omitempty" yaml:"status,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty" yaml:"created_at,omitempty"`
	ModifiedAt time.Time `json:"modified_at,omitempty" bson:"modified_at,omitempty" yaml:"modified_at,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s SysInfo) IsZero() bool {
	return s.Status == "" && s.CreatedAt.IsZero() && s.ModifiedAt.IsZero()
}

// IsActive restituisce true se lo status è attivo.
// "" (assente) è considerato attivo per compatibilità con i file YAML statici.
func (s SysInfo) IsActive() bool {
	return s.Status == "" || s.Status == "active"
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
