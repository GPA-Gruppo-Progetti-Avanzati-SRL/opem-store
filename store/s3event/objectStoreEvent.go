package s3event

import "go.mongodb.org/mongo-driver/v2/bson"
import "github.com/aws/aws-lambda-go/events"

// @tpm-schematics:start-region("top-file-section")

const (
	CollectionId = "s3-events"
	EntityType   = "s3-event"

	StatusReady = "ready"
	StatusDone  = "done"
	StatusIDC   = "idc"
)

// @tpm-schematics:end-region("top-file-section")

type ObjectStoreEvent struct {
	OId          bson.ObjectID        `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Bid          string               `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Et           string               `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Rip          bson.DateTime        `json:"_rip,omitempty" bson:"_rip,omitempty" yaml:"_rip,omitempty"`
	Status       string               `json:"_status,omitempty" bson:"_status,omitempty" yaml:"_status,omitempty"`
	StatusReason string               `json:"_status_reason,omitempty" bson:"_status_reason,omitempty" yaml:"_status_reason,omitempty"`
	Event        events.S3EventRecord `json:"event,omitempty" bson:"event,omitempty" yaml:"event,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s ObjectStoreEvent) IsZero() bool {
	return s.OId == bson.NilObjectID && s.Bid == "" && s.Et == "" && s.Rip == 0 && s.Status == "" && s.StatusReason == ""
}

type QueryResult struct {
	Records int                `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []ObjectStoreEvent `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
