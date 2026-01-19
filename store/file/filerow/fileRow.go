package filerow

import (
	"net/http"

	"go.mongodb.org/mongo-driver/v2/bson"
)
import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

// @tpm-schematics:start-region("top-file-section")

const (
	EntityType   = "file-row"
	CollectionId = "file-row"

	StatusError = "Err"
	StatusOK    = "OK"

	RowDataFormatText = "text"
	RowDataFormatJson = "json"

	EditorFixAna02 = "fix-ana-02"
)

// @tpm-schematics:end-region("top-file-section")

type FileRow struct {
	OId           bson.ObjectID       `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Domain        string              `json:"domain,omitempty" bson:"domain,omitempty" yaml:"domain,omitempty"`
	Site          string              `json:"site,omitempty" bson:"site,omitempty" yaml:"site,omitempty"`
	Bid           string              `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Et            string              `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Status        string              `json:"status,omitempty" bson:"status,omitempty" yaml:"status,omitempty"`
	Editor        string              `json:"editor,omitempty" bson:"editor,omitempty" yaml:"editor,omitempty"`
	RowData       string              `json:"row_data,omitempty" bson:"row_data,omitempty" yaml:"row_data,omitempty"`
	RowNumber     int32               `json:"row_number,omitempty" bson:"row_number,omitempty" yaml:"row_number,omitempty"`
	RowDataFormat string              `json:"row_data_format,omitempty" bson:"row_data_format,omitempty" yaml:"row_data_format,omitempty"`
	Errs          []string            `json:"errs,omitempty" bson:"errs,omitempty" yaml:"errs,omitempty"`
	File          commons.BidTextPair `json:"file,omitempty" bson:"file,omitempty" yaml:"file,omitempty"`
	SysInfo       commons.SysInfo     `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s FileRow) IsZero() bool {
	return s.OId == bson.NilObjectID && s.Domain == "" && s.Site == "" && s.Bid == "" && s.Et == "" && s.Status == "" && s.Editor == "" && s.RowData == "" && s.RowNumber == 0 && s.RowDataFormat == "" && len(s.Errs) == 0 && s.File.IsZero() && s.SysInfo.IsZero()
}

type QueryResult struct {
	Records int       `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []FileRow `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

const (
	FormResponseOK = http.StatusOK
)

type FormResponseErrors struct {
	Field string `json:"field,omitempty" bson:"field,omitempty" yaml:"field,omitempty"`
	Error string `json:"message,omitempty" bson:"message,omitempty" yaml:"message,omitempty"`
}

type FormResponse struct {
	Status      int    `json:"status,omitempty" bson:"status,omitempty" yaml:"status,omitempty"`
	Message     string `json:"message,omitempty" bson:"message,omitempty" yaml:"message,omitempty"`
	FieldErrors []FormResponseErrors
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
