package file

import (
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

func FilterMethodsGoInfo() string {
	i := fmt.Sprintf("tpm_morphia query filter support generated for %s package on %s", "author", time.Now().String())
	return i
}

// to be able to succesfully call this method you have to define a text index on the collection. The $text operator has some additional fields that are not supported yet.
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

/*
 * filter-object-id template: oId
 */

// AndOIdEqTo No Remarks
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

// @tpm-schematics:start-region("o-id-field-filter-section")
// @tpm-schematics:end-region("o-id-field-filter-section")

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
