package kv

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type SortOptions struct {
	document bson.D
}

func NewSortOptions(opts ...SortOption) SortOptions {
	sops := SortOptions{}
	for _, o := range opts {
		o(&sops)
	}
	return sops
}

func (pops SortOptions) Build() bson.D {
	return pops.document
}

type SortOption func(sops *SortOptions)

func WithSortByFieldAsc(fn string) SortOption {
	return func(sops *SortOptions) {
		sops.document = append(sops.document, bson.E{Key: fn, Value: 1})
	}
}

func WithSortByFieldDesc(fn string) SortOption {
	return func(sops *SortOptions) {
		sops.document = append(sops.document, bson.E{Key: fn, Value: -1})
	}
}

func CriteriaGoInfo() string {
	i := fmt.Sprintf("tpm_morphia query filter support generated for %s package on %s", "author", time.Now().String())
	return i
}

type ProjectionOptions struct {
	document bson.D
}

type ProjectionOption func(sops *ProjectionOptions)

func (pops ProjectionOptions) Build() bson.D {
	return pops.document
}

func WithNoId() ProjectionOption {
	return func(sops *ProjectionOptions) {
		sops.document = append(sops.document, bson.E{Key: "_id", Value: 0})
	}
}

const (
	Text                   string = "$text"
	Regex                  string = "$regex"
	TextLanguage           string = "$language"
	TextCaseSensitive             = "$caseSensitive"
	TextDiacriticSensitive        = "$diacriticSensitive"
	TextMeta                      = "$meta"
	TextMetaTextScore             = "textScore"
)

type Criterion func() bson.E
type Criteria []Criterion

type Filter struct {
	listOfCriteria []Criteria
}

func (f *Filter) Or() *Criteria {
	ca := make(Criteria, 0)
	if len(f.listOfCriteria) == 0 {
		f.listOfCriteria = make([]Criteria, 0, 5)
	}
	f.listOfCriteria = append(f.listOfCriteria, ca)
	return &f.listOfCriteria[len(f.listOfCriteria)-1]
}

var emptyFilter = bson.D{}

func (f *Filter) Build() bson.D {
	if len(f.listOfCriteria) == 0 {
		return emptyFilter
	}
	docA := bson.A{}
	for _, cas := range f.listOfCriteria {
		doc := bson.D{}
		for _, c := range cas {
			doc = append(doc, c())
		}
		docA = append(docA, doc)
	}
	if len(docA) == 1 {
		return docA[0].(bson.D)
	}
	return bson.D{{"$or", docA}}
}
