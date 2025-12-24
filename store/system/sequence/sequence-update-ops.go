package sequence

import (
	"context"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-mongo-common/util"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func ReleaseSequenceRange(coll *mongo.Collection, domain, site, bid string, rng Range) error {
	const semLogContext = "sequence::release-sequence-range"

	if !rng.HasNext() {
		log.Info().Str("range", rng.String()).Msg(semLogContext + " - no item to release")
		return nil
	}

	val := rng.CurrentValueAsInt()

	f := Filter{}
	f.Or().AndDomainEqTo(domain).AndSiteEqTo(site).AndEtEqTo(EntityType).AndValueEqTo(rng.To())
	fd := f.Build()
	log.Trace().Str("release-sequence-range-filter", util.MustToExtendedJsonString(fd, false, false)).Msg(semLogContext)

	updOpts := []UpdateOption{
		UpdateWithValue(val),
	}

	updDoc := GetUpdateDocumentFromOptions(updOpts...)
	ud := updDoc.Build()
	log.Trace().Str("release-sequence-range-document", util.MustToExtendedJsonString(ud, false, false)).Msg(semLogContext)

	resp, err := coll.UpdateOne(context.Background(), fd, ud)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return err
	}

	if resp.MatchedCount == 0 {
		log.Warn().Int32("value", val).Msg(semLogContext + " - didn't match any documents")
	}

	return nil
}
