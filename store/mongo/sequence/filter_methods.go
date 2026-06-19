package sequence

import (
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func FilterMethodsGoInfo() string {
	return fmt.Sprintf("tpm_morphia query filter support generated for sequence package")
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

func (ca *Criteria) AndDomainEqTo(p string) *Criteria {
	if p == "" {
		return ca
	}
	mName := fmt.Sprintf(DomainFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndDomainIn(p []string) *Criteria {
	if len(p) == 0 {
		return ca
	}
	mName := fmt.Sprintf(DomainFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndSiteEqTo(p string) *Criteria {
	if p == "" {
		return ca
	}
	mName := fmt.Sprintf(SiteFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndSiteIn(p []string) *Criteria {
	if len(p) == 0 {
		return ca
	}
	mName := fmt.Sprintf(SiteFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndValueLte(p int32) *Criteria {
	if p == 0 {
		return ca
	}
	mName := fmt.Sprintf(ValueFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$lte", p}}} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndValueEqTo(p int32) *Criteria {
	if p == 0 {
		return ca
	}
	mName := fmt.Sprintf(ValueFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

