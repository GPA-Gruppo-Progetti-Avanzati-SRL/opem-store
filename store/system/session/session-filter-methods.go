package session

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

func FilterMethodsGoInfo() string {
	i := fmt.Sprintf("tpm_morphia query filter support generated for %s package on %s", "author", time.Now().String())
	return i
}

/*
 * filter-object-id template: oId
 */

// AndOIdEqTo No Remarks
func (ca *Criteria) AndOIdEqTo(oId primitive.ObjectID) *Criteria {

	if oId == primitive.NilObjectID {
		return ca
	}

	mName := fmt.Sprintf(OIdFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: oId} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndOIdIn(p []primitive.ObjectID) *Criteria {

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

/*
 * filter-string template: userid
 */

// AndUseridEqTo No Remarks
func (ca *Criteria) AndUseridEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(UseridFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

// AndUseridIsNullOrUnset No Remarks
func (ca *Criteria) AndUseridIsNullOrUnset() *Criteria {

	mName := fmt.Sprintf(UseridFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndUseridIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(UseridFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

// @tpm-schematics:start-region("userid-field-filter-section")
// @tpm-schematics:end-region("userid-field-filter-section")

/*
 * filter-string template: nickname
 */

// AndNicknameEqTo No Remarks
func (ca *Criteria) AndNicknameEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(NicknameFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

// AndNicknameIsNullOrUnset No Remarks
func (ca *Criteria) AndNicknameIsNullOrUnset() *Criteria {

	mName := fmt.Sprintf(NicknameFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndNicknameIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(NicknameFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

// @tpm-schematics:start-region("nickname-field-filter-section")
// @tpm-schematics:end-region("nickname-field-filter-section")

/*
 * filter-string template: remoteaddr
 */

// AndRemoteaddrEqTo No Remarks
func (ca *Criteria) AndRemoteaddrEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(RemoteaddrFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

// AndRemoteaddrIsNullOrUnset No Remarks
func (ca *Criteria) AndRemoteaddrIsNullOrUnset() *Criteria {

	mName := fmt.Sprintf(RemoteaddrFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndRemoteaddrIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(RemoteaddrFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

// @tpm-schematics:start-region("remoteaddr-field-filter-section")
// @tpm-schematics:end-region("remoteaddr-field-filter-section")

// @tpm-schematics:start-region("bottom-file-section")

// AndSessionIdEqTo No Remarks
func (ca *Criteria) AndSessionIdEqTo(sid string) *Criteria {

	const SemLogContext = "r3ds9-core/session/and-hex-oid-eq-to"
	if sid == "" {
		return ca
	}

	oId, err := primitive.ObjectIDFromHex(sid)
	if err != nil {
		log.Error().Err(err).Msg(SemLogContext)
	}

	mName := fmt.Sprintf(OIdFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: oId} }
	*ca = append(*ca, c)
	return ca
}

// @tpm-schematics:end-region("bottom-file-section")
