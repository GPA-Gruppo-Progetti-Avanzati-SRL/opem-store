package sequence

import "go.mongodb.org/mongo-driver/bson/primitive"

// @tpm-schematics:start-region("top-file-section")

const (
	CollectionId = "sequence"
	EntityType   = "sequence"
)

// @tpm-schematics:end-region("top-file-section")

type Sequence struct {
	OId   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Bid   string             `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Et    string             `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Value int32              `json:"value,omitempty" bson:"value,omitempty" yaml:"value,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Sequence) IsZero() bool {
	return s.OId == primitive.NilObjectID && s.Bid == "" && s.Et == "" && s.Value == 0
}

type QueryResult struct {
	Records int        `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []Sequence `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

// @tpm-schematics:start-region("bottom-file-section")

type NextValOptions struct {
	SeqId           string
	Increment       int32
	CreateIfMissing bool
}

type NextValOption func(*NextValOptions)

func WithSeqId(s string) NextValOption {
	return func(opts *NextValOptions) {
		opts.SeqId = s
	}
}

func WithCreateIfMissing(b bool, d string) NextValOption {
	return func(opts *NextValOptions) {
		opts.CreateIfMissing = b
	}
}

// @tpm-schematics:end-region("bottom-file-section")
