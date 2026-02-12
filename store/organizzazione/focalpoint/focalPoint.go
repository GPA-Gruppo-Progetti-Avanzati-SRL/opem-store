package focalpoint

import (
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)
import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

// @tpm-schematics:start-region("top-file-section")
const (
	EntityType   = "focal-point"
	CollectionId = "focal-point"
)

// @tpm-schematics:end-region("top-file-section")

type FocalPoint struct {
	OId          bson.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Domain       string          `json:"domain,omitempty" bson:"domain,omitempty" yaml:"domain,omitempty"`
	Site         string          `json:"site,omitempty" bson:"site,omitempty" yaml:"site,omitempty"`
	Bid          string          `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Et           string          `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Name         string          `json:"name,omitempty" bson:"name,omitempty" yaml:"name,omitempty"`
	OfficerName  string          `json:"officer_name,omitempty" bson:"officer_name,omitempty" yaml:"officer_name,omitempty"`
	ExtendedName string          `json:"extended_name,omitempty" bson:"extended_name,omitempty" yaml:"extended_name,omitempty"`
	ReducedName  string          `json:"reduced_name,omitempty" bson:"reduced_name,omitempty" yaml:"reduced_name,omitempty"`
	Address      commons.Address `json:"address,omitempty" bson:"address,omitempty" yaml:"address,omitempty"`
	SysInfo      commons.SysInfo `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s FocalPoint) IsZero() bool {
	return s.OId == bson.NilObjectID && s.Domain == "" && s.Site == "" && s.Bid == "" && s.Et == "" && s.Name == "" && s.OfficerName == "" && s.ExtendedName == "" && s.ReducedName == "" && s.Address.IsZero() && s.SysInfo.IsZero()
}

type QueryResult struct {
	Records int          `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []FocalPoint `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

type FormResponseError struct {
	Field string `json:"field,omitempty" bson:"field,omitempty" yaml:"field,omitempty"`
	Error string `json:"message,omitempty" bson:"message,omitempty" yaml:"message,omitempty"`
}

type FormResponse struct {
	Status      int                 `json:"status,omitempty" bson:"status,omitempty" yaml:"status,omitempty"`
	Message     string              `json:"message,omitempty" bson:"message,omitempty" yaml:"message,omitempty"`
	FieldErrors []FormResponseError `json:"fieldErrors,omitempty" bson:"fieldErrors,omitempty" yaml:"fieldErrors,omitempty"`
	Document    *FocalPoint         `json:"document,omitempty" bson:"document,omitempty" yaml:"document,omitempty"`
}

// @tpm-schematics:start-region("bottom-file-section")

func (fp FocalPoint) BidAsAmmFocalPoint() string {
	return BidAsAmmFocalPoint(fp.Bid)
}

func BidAsAmmFocalPoint(bid string) string {
	s := fmt.Sprintf("%-9s", bid)
	return s[6:] + s[:6]
}

// @tpm-schematics:end-region("bottom-file-section")
