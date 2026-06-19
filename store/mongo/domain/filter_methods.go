package domain

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func FilterMethodsGoInfo() string {
	i := fmt.Sprintf("tpm_morphia query filter support generated for %s package on %s", "author", time.Now().String())
	return i
}

func (ca *Criteria) AndTextSearch(ssearch string) *Criteria {
	if ssearch == "" {
		return ca
	}
	c := func() bson.E {
		const TextOperator = "$text"
		return bson.E{Key: TextOperator, Value: bson.E{Key: "$search", Value: ssearch}}
	}
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndOIdEqTo(oId bson.ObjectID) *Criteria {
	if oId == bson.NilObjectID {
		return ca
	}
	mName := fmt.Sprintf(OIdFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: oId} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndOIdIn(p []bson.ObjectID) *Criteria {
	if len(p) == 0 {
		return ca
	}
	mName := fmt.Sprintf(OIdFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndBidEqTo(p string) *Criteria {
	if p == "" {
		return ca
	}
	mName := fmt.Sprintf(BidFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndBidIsNullOrUnset() *Criteria {
	mName := fmt.Sprintf(BidFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndBidIn(p []string) *Criteria {
	if len(p) == 0 {
		return ca
	}
	mName := fmt.Sprintf(BidFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndCodeNotEqTo(p string) *Criteria {
	if p == "" {
		return ca
	}
	mName := fmt.Sprintf(BidFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$ne", p}}} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndEtEqTo(p string) *Criteria {
	if p == "" {
		return ca
	}
	mName := fmt.Sprintf(EtFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndEtIsNullOrUnset() *Criteria {
	mName := fmt.Sprintf(EtFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndEtIn(p []string) *Criteria {
	if len(p) == 0 {
		return ca
	}
	mName := fmt.Sprintf(EtFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}
