package s3event

import (
	"context"
	"errors"
	"time"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-mongo-common/util"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func UpdateStatus(coll *mongo.Collection, bid string, status commons.StatusCodeTextPair) error {
	const semLogContext = "object-store-event::update-status"

	f := Filter{}
	f.Or().AndEtEqTo(EntityType).AndBidEqTo(bid)
	fd := f.Build()
	log.Trace().Str("update-status-filter", util.MustToExtendedJsonString(fd, false, false)).Msg(semLogContext)

	updOpts := []UpdateOption{
		UpdateWith_status(&status),
		UpdateWith_rip(bson.NewDateTimeFromTime(time.Now())),
	}
	updDoc := GetUpdateDocumentFromOptions(updOpts...)
	ud := updDoc.Build()
	log.Trace().Str("update-status", util.MustToExtendedJsonString(ud, false, false)).Msg(semLogContext)

	resp, err := coll.UpdateOne(context.Background(), fd, ud)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return err
	}

	if resp.ModifiedCount == 0 {
		err = errors.New("no document updated matched")
		log.Error().Err(err).Msg(semLogContext)
		return err
	}

	return nil
}
