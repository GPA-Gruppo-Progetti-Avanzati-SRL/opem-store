package person

import "go.mongodb.org/mongo-driver/v2/bson"
import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type BirthInfo struct {
	Date     bson.DateTime       `json:"date,omitempty" bson:"date,omitempty" yaml:"date,omitempty"`
	Country  commons.BidTextPair `json:"country,omitempty" bson:"country,omitempty" yaml:"country,omitempty"`
	County   commons.BidTextPair `json:"county,omitempty" bson:"county,omitempty" yaml:"county,omitempty"`
	Townhall commons.BidTextPair `json:"townhall,omitempty" bson:"townhall,omitempty" yaml:"townhall,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s BirthInfo) IsZero() bool {
	return s.Date == 0 && s.Country.IsZero() && s.County.IsZero() && s.Townhall.IsZero()
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
