package sequence

import (
	"context"
	"errors"
	"time"

	model "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/model"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-mongo-common/util"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func FindByPk(collection *mongo.Collection, domain, site, bid string, mustFind bool, findOptions *options.FindOneOptionsBuilder) (*model.Sequence, bool, error) {
	const semLogContext = "sequence::find-by-pk"
	evtTraceLog := log.Trace()
	evtErrLog := log.Error()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ent := model.Sequence{}

	f := Filter{}
	f.Or().AndDomainEqTo(domain).AndSiteEqTo(site).AndBidEqTo(bid).AndEtEqTo(EntityType)
	fd := f.Build()
	evtTraceLog = evtTraceLog.Str("filter", util.MustToExtendedJsonString(fd, false, false))
	err := collection.FindOne(ctx, fd, findOptions).Decode(&ent)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		evtErrLog.Err(err).Msg(semLogContext)
		return nil, false, err
	} else {
		if err != nil {
			if mustFind {
				evtTraceLog.Msg(semLogContext + " document not found")
				return nil, false, err
			}
			evtTraceLog.Msg(semLogContext + " document not found but allowed")
			return nil, false, nil
		} else {
			evtTraceLog.Msg(semLogContext + " document found")
		}
	}

	return &ent, true, nil
}

func Find(collection *mongo.Collection, f *Filter, withCount bool, findOptions *options.FindOptionsBuilder) (model.SequenceQueryResult, error) {
	const semLogContext = "sequence::find"
	fd := f.Build()
	evtTraceLog := log.Trace().Str("filter", util.MustToExtendedJsonString(fd, false, false))
	evtErrLog := log.Error().Str("filter", util.MustToExtendedJsonString(fd, false, false))
	evtTraceLog.Msg(semLogContext)

	qr := model.SequenceQueryResult{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if withCount {
		countDocsOptions := options.Count()
		nr, err := collection.CountDocuments(ctx, fd, countDocsOptions)
		if err != nil {
			evtErrLog.Err(err).Msg(semLogContext)
			return qr, err
		}
		qr.Records = int(nr)
	}

	cur, err := collection.Find(ctx, fd, findOptions)
	if err != nil {
		evtErrLog.Err(err).Msg(semLogContext)
		return qr, err
	}

	for cur.Next(context.Background()) {
		dto := model.Sequence{}
		err = cur.Decode(&dto)
		if err != nil {
			return qr, err
		}
		qr.Data = append(qr.Data, dto)
	}

	if cur.Err() != nil {
		return qr, cur.Err()
	}

	return qr, nil
}
