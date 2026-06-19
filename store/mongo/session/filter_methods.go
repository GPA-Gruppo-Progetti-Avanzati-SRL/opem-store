package session

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
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

func (ca *Criteria) AndUseridEqTo(p string) *Criteria {
	if p == "" {
		return ca
	}
	mName := fmt.Sprintf(UseridFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

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

func (ca *Criteria) AndNicknameEqTo(p string) *Criteria {
	if p == "" {
		return ca
	}
	mName := fmt.Sprintf(NicknameFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

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

func (ca *Criteria) AndRemoteAddrEqTo(p string) *Criteria {
	if p == "" {
		return ca
	}
	mName := fmt.Sprintf(RemoteAddrFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndRemoteAddrIsNullOrUnset() *Criteria {
	mName := fmt.Sprintf(RemoteAddrFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndRemoteAddrIn(p []string) *Criteria {
	if len(p) == 0 {
		return ca
	}
	mName := fmt.Sprintf(RemoteAddrFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndSessionIdEqTo(sid string) *Criteria {
	const SemLogContext = "opem-core/session/and-hex-oid-eq-to"
	if sid == "" {
		return ca
	}
	oId, err := bson.ObjectIDFromHex(sid)
	if err != nil {
		log.Error().Err(err).Msg(SemLogContext)
	}
	mName := fmt.Sprintf(OIdFieldName)
	c := func() bson.E { return bson.E{Key: mName, Value: oId} }
	*ca = append(*ca, c)
	return ca
}
