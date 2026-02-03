package sequence

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// @tpm-schematics:start-region("top-file-section")

const (
	CollectionId = "sequence"
	EntityType   = "sequence"

	BoxSequenceBid   = "seq-mag"
	CardSequenceId   = "seq-card"
	PersonSequenceId = "seq-pers"
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

type NextValueProvider interface {
	Next() (string, error)
}

type Range struct {
	from    int32
	to      int32
	format  string
	current int32
}

func (r *Range) String() string {
	return fmt.Sprintf("[%d]: from: %d, to: %d, current: %d, format: %s", r.to-r.from, r.from, r.to, r.current, r.format)
}

func (r *Range) CurrentValue() string {
	if r.current < 0 || r.current > (r.to-r.from) {
		return fmt.Sprintf("invalid range position: %d", r.current)
	}
	return fmt.Sprintf(r.format, r.from+r.current)
}

func (r *Range) CurrentValueAsInt() int32 {
	if r.current < 0 || r.current > (r.to-r.from) {
		return -1
	}
	return r.from + r.current
}

func (r *Range) To() int32 {
	return r.to
}

func (r *Range) HasNext() bool {
	if r.current < 0 {
		return true
	}

	return r.current < (r.to - r.from)
}

func (r *Range) Next() (string, error) {
	r.current++
	return r.CurrentValue(), nil
}

type RangeMap map[string]*Range

func (rm RangeMap) NextVal(n string) (string, error) {
	if r, ok := rm[n]; ok {
		if r.HasNext() {
			v, _ := r.Next()
			return v, nil
		} else {
			return "", errors.New("range exhausted: " + r.CurrentValue())
		}
	}

	return "", errors.New("invalid range name: " + n)
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
