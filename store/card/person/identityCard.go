package person

import "go.mongodb.org/mongo-driver/v2/bson"
import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type IdentityCard struct {
	Type      string              `json:"type,omitempty" bson:"type,omitempty" yaml:"type,omitempty"`
	Code      string              `json:"code,omitempty" bson:"code,omitempty" yaml:"code,omitempty"`
	IssueDate bson.DateTime       `json:"issue_date,omitempty" bson:"issue_date,omitempty" yaml:"issue_date,omitempty"`
	Townhall  commons.BidTextPair `json:"townhall,omitempty" bson:"townhall,omitempty" yaml:"townhall,omitempty"`
	County    commons.BidTextPair `json:"county,omitempty" bson:"county,omitempty" yaml:"county,omitempty"`
	Issuer    string              `json:"issuer,omitempty" bson:"issuer,omitempty" yaml:"issuer,omitempty"`
	ExpiresAt bson.DateTime       `json:"expires_at,omitempty" bson:"expires_at,omitempty" yaml:"expires_at,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s IdentityCard) IsZero() bool {
	return s.Type == "" && s.Code == "" && s.IssueDate == 0 && s.Townhall.IsZero() && s.County.IsZero() && s.Issuer == "" && s.ExpiresAt == 0
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
