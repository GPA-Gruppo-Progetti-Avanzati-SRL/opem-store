package prodotto

import "go.mongodb.org/mongo-driver/v2/bson"
import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

// @tpm-schematics:start-region("top-file-section")
const (
	EntityType   = "prodotto"
	CollectionId = "prodotto"
)

// @tpm-schematics:end-region("top-file-section")

type Prodotto struct {
	OId          bson.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Domain       string          `json:"domain,omitempty" bson:"domain,omitempty" yaml:"domain,omitempty"`
	Site         string          `json:"site,omitempty" bson:"site,omitempty" yaml:"site,omitempty"`
	Bid          string          `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Et           string          `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Name         string          `json:"name,omitempty" bson:"name,omitempty" yaml:"name,omitempty"`
	PrimaryFunct string          `json:"primary_funct,omitempty" bson:"primary_funct,omitempty" yaml:"primary_funct,omitempty"`
	ExpirationAt string          `json:"expiration_at,omitempty" bson:"expiration_at,omitempty" yaml:"expiration_at,omitempty"`
	PersBureau   string          `json:"pers_bureau,omitempty" bson:"pers_bureau,omitempty" yaml:"pers_bureau,omitempty"`
	Properties   bson.M          `json:"properties,omitempty" bson:"properties,omitempty" yaml:"properties,omitempty"`
	Apps         []AppDefinition `json:"apps,omitempty" bson:"apps,omitempty" yaml:"apps,omitempty"`
	SysInfo      commons.SysInfo `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Prodotto) IsZero() bool {
	return s.OId == bson.NilObjectID && s.Domain == "" && s.Site == "" && s.Bid == "" && s.Et == "" && s.Name == "" && s.PrimaryFunct == "" && s.ExpirationAt == "" && s.PersBureau == "" && len(s.Properties) == 0 && len(s.Apps) == 0 && s.SysInfo.IsZero()
}

type QueryResult struct {
	Records int        `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []Prodotto `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
