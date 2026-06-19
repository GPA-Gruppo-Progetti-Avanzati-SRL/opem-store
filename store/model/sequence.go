package model

type Sequence struct {
	OId    string `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Bid    string `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Et     string `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Domain string `json:"domain,omitempty" bson:"domain,omitempty" yaml:"domain,omitempty"`
	Site   string `json:"site,omitempty" bson:"site,omitempty" yaml:"site,omitempty"`
	Value  int32  `json:"value,omitempty" bson:"value,omitempty" yaml:"value,omitempty"`
	Format string `json:"format,omitempty" bson:"format,omitempty" yaml:"format,omitempty"`
	Prefix string `json:"prefix,omitempty" bson:"prefix,omitempty" yaml:"prefix,omitempty"`
}

func (s Sequence) IsZero() bool {
	return s.OId == "" && s.Bid == "" && s.Et == "" && s.Domain == "" && s.Site == "" && s.Value == 0 && s.Format == "" && s.Prefix == ""
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

type SequenceQueryResult struct {
	Records int        `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []Sequence `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}
