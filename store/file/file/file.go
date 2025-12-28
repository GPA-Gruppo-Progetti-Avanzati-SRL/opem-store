package file

import "go.mongodb.org/mongo-driver/v2/bson"
import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

// @tpm-schematics:start-region("top-file-section")

const (
	EntityType   = "file"
	CollectionId = "file"

	StreamDirectionIn   = "I"
	StreamDirectionOut  = "O"
	StatusProcessed     = "PR"
	StatusInProcess     = "IP"
	StatusToBeProcessed = "TP"
	TypeCardRequest     = "RC01"

	TypeAnonymousCardRequest         = "RC02"
	TypeAnonymousCardRequestResponse = "RC02"
	TypeAnagraficheConsegnaCarte02   = "ANA02"
	TypeC3CardFile                   = "RC04"

	TypeNamedCardRequest        = "RC03"
	TypeCardDelivery            = "RC04"
	TypeCardRequestReportRange  = "WHRR"
	TypeCardRequestReportHolder = "WHRH"

	OPERATION_TYPE_CARD_REQUEST          = "RC"
	OPERATION_TYPE_WHAREHOUSE_ANONYMOUS  = "MA"
	OPERATION_TYPE_WHAREHOUSE_NOMINATIVE = "MN"
)

// @tpm-schematics:end-region("top-file-section")

type File struct {
	OId        bson.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Domain     string          `json:"domain,omitempty" bson:"domain,omitempty" yaml:"domain,omitempty"`
	Site       string          `json:"site,omitempty" bson:"site,omitempty" yaml:"site,omitempty"`
	Bid        string          `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Et         string          `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Status     string          `json:"status,omitempty" bson:"status,omitempty" yaml:"status,omitempty"`
	Stats      Stat            `json:"stats,omitempty" bson:"stats,omitempty" yaml:"stats,omitempty"`
	Type       string          `json:"type,omitempty" bson:"type,omitempty" yaml:"type,omitempty"`
	Name       string          `json:"name,omitempty" bson:"name,omitempty" yaml:"name,omitempty"`
	BlobBucket string          `json:"blob_bucket,omitempty" bson:"blob_bucket,omitempty" yaml:"blob_bucket,omitempty"`
	BlobKey    string          `json:"blob_key,omitempty" bson:"blob_key,omitempty" yaml:"blob_key,omitempty"`
	InOut      string          `json:"in_out,omitempty" bson:"in_out,omitempty" yaml:"in_out,omitempty"`
	SysInfo    commons.SysInfo `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	BlobTmpFile string `json:"-" bson:"-" yaml:"-"`
	// @tpm-schematics:end-region("struct-section")
}

func (s File) IsZero() bool {
	return s.OId == bson.NilObjectID && s.Domain == "" && s.Site == "" && s.Bid == "" && s.Et == "" && s.Status == "" && s.Stats.IsZero() && s.Type == "" && s.Name == "" && s.BlobBucket == "" && s.BlobKey == "" && s.InOut == "" && s.SysInfo.IsZero()
}

type QueryResult struct {
	Records int    `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []File `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
