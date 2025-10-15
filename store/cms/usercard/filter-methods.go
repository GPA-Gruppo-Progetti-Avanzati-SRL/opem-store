package usercard

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func FilterGoInfo() string {
	i := fmt.Sprintf("tpm_morphia query filter support generated for %s package on %s", "author", time.Now().String())
	return i
}

//----- oId of type object-id
//----- oId - object-id -  [oId]

// AndOIdEqTo No Remarks
func (ca *Criteria) AndOIdEqTo(oId bson.ObjectID) *Criteria {

	if oId == bson.NilObjectID {
		return ca
	}

	mName := fmt.Sprintf(OID)
	c := func() bson.E { return bson.E{Key: mName, Value: oId} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndOIdIn(p []bson.ObjectID) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(OID)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

//----- userId of type string
//----- userId - string -  [userId]

// AndUserIdEqTo No Remarks
func (ca *Criteria) AndUserIdEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(USERID)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

// AndUserIdIsNullOrUnset No Remarks
func (ca *Criteria) AndUserIdIsNullOrUnset() *Criteria {

	mName := fmt.Sprintf(USERID)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndUserIdIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(USERID)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

//----- objType of type string
//----- objType - string -  [objType]

// AndObjTypeEqTo No Remarks
func (ca *Criteria) AndObjTypeEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(OBJTYPE)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

// AndObjTypeIsNullOrUnset No Remarks
func (ca *Criteria) AndObjTypeIsNullOrUnset() *Criteria {

	mName := fmt.Sprintf(OBJTYPE)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndObjTypeIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(OBJTYPE)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}
