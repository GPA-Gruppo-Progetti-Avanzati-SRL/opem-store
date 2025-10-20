package box

import "go.mongodb.org/mongo-driver/v2/bson"
import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

// @tpm-schematics:start-region("top-file-section")
const (
	EntityType   = "box"
	CollectionId = "box"
)

// @tpm-schematics:end-region("top-file-section")

type Box struct {
	OId        bson.ObjectID       `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Domain     string              `json:"domain,omitempty" bson:"domain,omitempty" yaml:"domain,omitempty"`
	Site       string              `json:"site,omitempty" bson:"site,omitempty" yaml:"site,omitempty"`
	Bid        string              `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Et         string              `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Magazzino  commons.BidTextPair `json:"magazzino,omitempty" bson:"magazzino,omitempty" yaml:"magazzino,omitempty"`
	Prodotto   commons.BidTextPair `json:"prodotto,omitempty" bson:"prodotto,omitempty" yaml:"prodotto,omitempty"`
	FocalPoint commons.BidTextPair `json:"focal_point,omitempty" bson:"focal_point,omitempty" yaml:"focal_point,omitempty"`
	Info       Info                `json:"info,omitempty" bson:"info,omitempty" yaml:"info,omitempty"`
	Status     Status              `json:"status,omitempty" bson:"status,omitempty" yaml:"status,omitempty"`
	Recipient  commons.Address     `json:"recipient,omitempty" bson:"recipient,omitempty" yaml:"recipient,omitempty"`
	Events     []commons.Event     `json:"events,omitempty" bson:"events,omitempty" yaml:"events,omitempty"`
	Notes      []commons.Note      `json:"notes,omitempty" bson:"notes,omitempty" yaml:"notes,omitempty"`
	SysInfo    commons.SysInfo     `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Box) IsZero() bool {
	return s.OId == bson.NilObjectID && s.Domain == "" && s.Site == "" && s.Bid == "" && s.Et == "" && s.Magazzino.IsZero() && s.Prodotto.IsZero() && s.FocalPoint.IsZero() && s.Info.IsZero() && s.Status.IsZero() && s.Recipient.IsZero() && len(s.Events) == 0 && len(s.Notes) == 0 && s.SysInfo.IsZero()
}

type QueryResult struct {
	Records int   `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []Box `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
