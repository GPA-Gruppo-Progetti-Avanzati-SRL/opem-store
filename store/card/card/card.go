package card

import "go.mongodb.org/mongo-driver/v2/bson"
import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

// @tpm-schematics:start-region("top-file-section")

import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/card/person"

const (
	EntityType   = "card"
	CollectionId = "card"

	StatusErr                        = "error"
	StatusAMagazzinoInAttesaConferma = "ME" // value: "ME A magazzino in attesa conferma da ente/processor"
	StatusAMagazzinoDisponibili      = "MA"
	StatusAttiva                     = "AT"
	EventReintegroMagazzino          = "reintegro-magazzino"
)

// @tpm-schematics:end-region("top-file-section")

type Card struct {
	OId                   bson.ObjectID         `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Domain                string                `json:"domain,omitempty" bson:"domain,omitempty" yaml:"domain,omitempty"`
	Site                  string                `json:"site,omitempty" bson:"site,omitempty" yaml:"site,omitempty"`
	Bid                   string                `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Et                    string                `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	CardNumber            commons.ValueTextPair `json:"card_number,omitempty" bson:"card_number,omitempty" yaml:"card_number,omitempty"`
	CardType              string                `json:"card_type,omitempty" bson:"card_type,omitempty" yaml:"card_type,omitempty"`
	Status                string                `json:"status,omitempty" bson:"status,omitempty" yaml:"status,omitempty"`
	IdCardExt             string                `json:"id_card_ext,omitempty" bson:"id_card_ext,omitempty" yaml:"id_card_ext,omitempty"`
	EnvelopeNumber        commons.ValueTextPair `json:"envelope_number,omitempty" bson:"envelope_number,omitempty" yaml:"envelope_number,omitempty"`
	Funct                 string                `json:"funct,omitempty" bson:"funct,omitempty" yaml:"funct,omitempty"`
	FocalPoint            commons.BidTextPair   `json:"focal_point,omitempty" bson:"focal_point,omitempty" yaml:"focal_point,omitempty"`
	Product               commons.BidTextPair   `json:"product,omitempty" bson:"product,omitempty" yaml:"product,omitempty"`
	LayoutCode            string                `json:"layout_code,omitempty" bson:"layout_code,omitempty" yaml:"layout_code,omitempty"`
	CorporateCode         string                `json:"corporate_code,omitempty" bson:"corporate_code,omitempty" yaml:"corporate_code,omitempty"`
	Box                   commons.BidTextPair   `json:"box,omitempty" bson:"box,omitempty" yaml:"box,omitempty"`
	Holder                CardHolder            `json:"holder,omitempty" bson:"holder,omitempty" yaml:"holder,omitempty"`
	Apps                  []CardApp             `json:"apps,omitempty" bson:"apps,omitempty" yaml:"apps,omitempty"`
	Addresses             []commons.Address     `json:"addresses,omitempty" bson:"addresses,omitempty" yaml:"addresses,omitempty"`
	Events                []commons.Event       `json:"events,omitempty" bson:"events,omitempty" yaml:"events,omitempty"`
	ExpiresAt             bson.DateTime         `json:"expires_at,omitempty" bson:"expires_at,omitempty" yaml:"expires_at,omitempty"`
	IssueDate             bson.DateTime         `json:"issue_date,omitempty" bson:"issue_date,omitempty" yaml:"issue_date,omitempty"`
	IssueConfirmationDate bson.DateTime         `json:"issue_confirmation_date,omitempty" bson:"issue_confirmation_date,omitempty" yaml:"issue_confirmation_date,omitempty"`
	ActDate               bson.DateTime         `json:"act_date,omitempty" bson:"act_date,omitempty" yaml:"act_date,omitempty"`
	SysInfo               commons.SysInfo       `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	DocPerson person.Person `json:"doc_person,omitempty" bson:"doc_person,omitempty" yaml:"doc_person,omitempty"`
	// @tpm-schematics:end-region("struct-section")
}

func (s Card) IsZero() bool {
	return s.OId == bson.NilObjectID && s.Domain == "" && s.Site == "" && s.Bid == "" && s.Et == "" && s.CardNumber.IsZero() && s.CardType == "" && s.Status == "" && s.IdCardExt == "" && s.EnvelopeNumber.IsZero() && s.Funct == "" && s.FocalPoint.IsZero() && s.Product.IsZero() && s.LayoutCode == "" && s.CorporateCode == "" && s.Box.IsZero() && s.Holder.IsZero() && len(s.Apps) == 0 && len(s.Addresses) == 0 && len(s.Events) == 0 && s.ExpiresAt == 0 && s.IssueDate == 0 && s.IssueConfirmationDate == 0 && s.ActDate == 0 && s.SysInfo.IsZero()
}

type QueryResult struct {
	Records int    `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []Card `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

// @tpm-schematics:start-region("bottom-file-section")

func (s Card) StatusIn(stat ...string) bool {
	if len(stat) == 0 {
		return false
	}

	for _, v := range stat {
		if v == s.Status {
			return true
		}
	}

	return false
}

// @tpm-schematics:end-region("bottom-file-section")
