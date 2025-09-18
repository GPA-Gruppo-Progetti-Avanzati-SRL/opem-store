package focalpoint

import "go.mongodb.org/mongo-driver/bson/primitive"

// @tpm-schematics:start-region("top-file-section")
const (
	EntityType   = "FOCAL-POINT"
	CollectionId = "org_organizzazione"
)

// @tpm-schematics:end-region("top-file-section")

type FocalPoint struct {
	OId         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Domain      string             `json:"domain,omitempty" bson:"domain,omitempty" yaml:"domain,omitempty"`
	Site        string             `json:"site,omitempty" bson:"site,omitempty" yaml:"site,omitempty"`
	Bid         string             `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Et          string             `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	OfficerName string             `json:"officer_name,omitempty" bson:"officer_name,omitempty" yaml:"officer_name,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty" yaml:"status,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s FocalPoint) IsZero() bool {
	return s.OId == primitive.NilObjectID && s.Domain == "" && s.Site == "" && s.Bid == "" && s.Et == "" && s.OfficerName == "" && s.Status == ""
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
