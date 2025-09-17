package prodotto

import "go.mongodb.org/mongo-driver/bson/primitive"

// @tpm-schematics:start-region("top-file-section")
const (
	EntityType   = "PRODOTTO"
	CollectionId = "prd_prodotto"
)

// @tpm-schematics:end-region("top-file-section")

type Prodotto struct {
	OId    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Domain string             `json:"domain,omitempty" bson:"domain,omitempty" yaml:"domain,omitempty"`
	Ns     string             `json:"ns,omitempty" bson:"ns,omitempty" yaml:"ns,omitempty"`
	Bid    string             `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Et     string             `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty" yaml:"name,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Prodotto) IsZero() bool {
	return s.OId == primitive.NilObjectID && s.Domain == "" && s.Ns == "" && s.Bid == "" && s.Et == "" && s.Name == ""
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
