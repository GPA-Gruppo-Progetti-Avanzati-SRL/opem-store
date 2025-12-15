package s3event

import (
	"time"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"go.mongodb.org/mongo-driver/v2/bson"
)
import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
import "github.com/aws/aws-lambda-go/events"

// @tpm-schematics:start-region("top-file-section")

const (
	CollectionId = "s3-events"
	EntityType   = "s3-event"

	StatusReady          = "ready"
	StatusDone           = "done"
	StatusIDC            = "idc"
	StatusError          = "error"
	StatusAlreadyPresent = "already-present"
)

// @tpm-schematics:end-region("top-file-section")

type ObjectStoreEvent struct {
	OId    bson.ObjectID              `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Bid    string                     `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Et     string                     `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Rip    bson.DateTime              `json:"_rip,omitempty" bson:"_rip,omitempty" yaml:"_rip,omitempty"`
	Status commons.StatusCodeTextPair `json:"_status,omitempty" bson:"_status,omitempty" yaml:"_status,omitempty"`
	Event  events.S3EventRecord       `json:"event,omitempty" bson:"event,omitempty" yaml:"event,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s ObjectStoreEvent) IsZero() bool {
	return s.OId == bson.NilObjectID && s.Bid == "" && s.Et == "" && s.Rip == 0 && s.Status.IsZero()
}

type QueryResult struct {
	Records int                `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []ObjectStoreEvent `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

// @tpm-schematics:start-region("bottom-file-section")

func NewObjectStoreEvent(bucket, objectKey string, size int64) ObjectStoreEvent {

	const semLogContext = "object-store-event::new-object-store-event"

	evt := ObjectStoreEvent{
		Bid: util.NewUUID(),
		Et:  EntityType,
		Status: commons.StatusCodeTextPair{
			Code: StatusReady,
		},
		Event: events.S3EventRecord{
			EventVersion: "2.1",
			EventSource:  "aws:s3",
			AWSRegion:    "eu-central-1",
			EventTime:    time.Now(),
			EventName:    "ObjectCreated:Put",
			PrincipalID: events.S3UserIdentity{
				PrincipalID: "AWS:AIDA5MH3DL5HCND5NTUJU",
			},
			RequestParameters: events.S3RequestParameters{
				SourceIPAddress: "127.0.0.0",
			},
			ResponseElements: nil,
			S3: events.S3Entity{
				SchemaVersion:   "1.0",
				ConfigurationID: "1",
				Bucket: events.S3Bucket{
					Name: bucket,
					OwnerIdentity: events.S3UserIdentity{
						PrincipalID: "A25IM5X7I3AW2K",
					},
					Arn: "ws:s3:::gpagroup-dev-opem-flussi",
				},
				Object: events.S3Object{
					Key:           objectKey,
					Size:          size,
					URLDecodedKey: objectKey,
					VersionID:     "",
					ETag:          "d28493bdea99e01427c1c02764e9e8d5",
					Sequencer:     "00691DCEB726E5F05B",
				},
			},
		},
	}

	return evt
}

// @tpm-schematics:end-region("bottom-file-section")
