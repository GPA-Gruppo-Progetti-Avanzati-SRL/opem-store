package sequence

import (
	"errors"
	"fmt"
)

const (
	CollectionId = "sequence"
	EntityType   = "sequence"

	BoxSequenceBid   = "seq-mag"
	CardSequenceId   = "seq-card"
	PersonSequenceId = "seq-pers"
)

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
		}
		return "", errors.New("range exhausted: " + r.CurrentValue())
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
