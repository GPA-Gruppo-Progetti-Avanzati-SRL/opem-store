package kv

import (
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func FilterMethodsGoInfo() string {
	return fmt.Sprintf("tpm_morphia query filter support generated for kv package")
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

func (ca *Criteria) AndBidEqTo(p string) *Criteria {
	if p == "" {
		return ca
	}
	mName := fmt.Sprintf(BidFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
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

func (ca *Criteria) AndScopeEqTo(p string) *Criteria {
	if p == "" {
		return ca
	}
	mName := fmt.Sprintf(ScopeFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndScopeIsNullOrUnset() *Criteria {
	mName := fmt.Sprintf(ScopeFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndScopeIn(p []string) *Criteria {
	if len(p) == 0 {
		return ca
	}
	mName := fmt.Sprintf(ScopeFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
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

func (ca *Criteria) AndEtIn(p []string) *Criteria {
	if len(p) == 0 {
		return ca
	}
	mName := fmt.Sprintf(EtFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndCategoryEqTo(p string) *Criteria {
	if p == "" {
		return ca
	}
	mName := fmt.Sprintf(CategoryFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndCategoryIsNullOrUnset() *Criteria {
	mName := fmt.Sprintf(CategoryFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndCategoryIn(p []string) *Criteria {
	if len(p) == 0 {
		return ca
	}
	mName := fmt.Sprintf(CategoryFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}
