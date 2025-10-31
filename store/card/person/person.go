package person

import "go.mongodb.org/mongo-driver/v2/bson"
import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

// @tpm-schematics:start-region("top-file-section")
const (
	EntityType   = "person"
	CollectionId = "person"
)

// @tpm-schematics:end-region("top-file-section")

type Person struct {
	OId        bson.ObjectID       `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Domain     string              `json:"domain,omitempty" bson:"domain,omitempty" yaml:"domain,omitempty"`
	Site       string              `json:"site,omitempty" bson:"site,omitempty" yaml:"site,omitempty"`
	Bid        string              `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Et         string              `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	FirstName  string              `json:"first_name,omitempty" bson:"first_name,omitempty" yaml:"first_name,omitempty"`
	LastName   string              `json:"last_name,omitempty" bson:"last_name,omitempty" yaml:"last_name,omitempty"`
	Cf         string              `json:"cf,omitempty" bson:"cf,omitempty" yaml:"cf,omitempty"`
	MaleFemale string              `json:"male_female,omitempty" bson:"male_female,omitempty" yaml:"male_female,omitempty"`
	Birth      BirthInfo           `json:"birth,omitempty" bson:"birth,omitempty" yaml:"birth,omitempty"`
	IdDocument IdentityCard        `json:"id_document,omitempty" bson:"id_document,omitempty" yaml:"id_document,omitempty"`
	Country    commons.BidTextPair `json:"country,omitempty" bson:"country,omitempty" yaml:"country,omitempty"`
	Addresses  []commons.Address   `json:"addresses,omitempty" bson:"addresses,omitempty" yaml:"addresses,omitempty"`
	SysInfo    commons.SysInfo     `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Person) IsZero() bool {
	return s.OId == bson.NilObjectID && s.Domain == "" && s.Site == "" && s.Bid == "" && s.Et == "" && s.FirstName == "" && s.LastName == "" && s.Cf == "" && s.MaleFemale == "" && s.Birth.IsZero() && s.IdDocument.IsZero() && s.Country.IsZero() && len(s.Addresses) == 0 && s.SysInfo.IsZero()
}

type QueryResult struct {
	Records int      `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []Person `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
