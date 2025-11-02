package sequence

import "go.mongodb.org/mongo-driver/v2/bson"

// @tpm-schematics:start-region("top-file-section")

const (
	CollectionId = "sequence"
	EntityType   = "sequence"

	BoxSequenceBid = "seq-mag"
	CardSequenceId = "seq-card"
)

// @tpm-schematics:end-region("top-file-section")

type Sequence struct {
	OId    bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Bid    string        `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Et     string        `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Domain string        `json:"domain,omitempty" bson:"domain,omitempty" yaml:"domain,omitempty"`
	Site   string        `json:"site,omitempty" bson:"site,omitempty" yaml:"site,omitempty"`
	Value  int32         `json:"value,omitempty" bson:"value,omitempty" yaml:"value,omitempty"`
	Format string        `json:"format,omitempty" bson:"format,omitempty" yaml:"format,omitempty"`
	Prefix string        `json:"prefix,omitempty" bson:"prefix,omitempty" yaml:"prefix,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Sequence) IsZero() bool {
	return s.OId == bson.NilObjectID && s.Bid == "" && s.Et == "" && s.Domain == "" && s.Site == "" && s.Value == 0 && s.Format == "" && s.Prefix == ""
}

type QueryResult struct {
	Records int        `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []Sequence `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

// @tpm-schematics:start-region("bottom-file-section")

type Range struct {
	From   int32
	To     int32
	Format string
}
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

func WithCreateIfMissing(b bool) NextValOption {
	return func(opts *NextValOptions) {
		opts.CreateIfMissing = b
	}
}

func WithIncrement(inc int32) NextValOption {
	return func(opts *NextValOptions) {
		opts.Increment = inc
	}
}

func (seq Sequence) FormatSpecifier() string {
	format := "%d"
	pfix := ""

	if seq.Format != "" {
		format = seq.Format
	}

	if seq.Prefix != "" {
		pfix = seq.Prefix
	}

	return pfix + format
}

// @tpm-schematics:end-region("bottom-file-section")
