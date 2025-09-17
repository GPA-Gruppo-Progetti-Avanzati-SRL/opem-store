package file

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

func (ca *Criteria) AndHexOIdEqTo(oId string) *Criteria {
	const semLogContext = "mongo-file::and-hex-oid-eq-to"
	if oId == "" {
		return ca
	}

	objId, err := primitive.ObjectIDFromHex(oId)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return ca
	}
	mName := fmt.Sprintf(OIdFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: objId} }
	*ca = append(*ca, c)
	return ca
}

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
