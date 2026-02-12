package comune

import (
	"context"
	"errors"
	"time"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-mongo-common/util"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

// FindByPk ...
// @tpm-schematics:start-region("find-by-pk-signature-section")
func FindByPk(collection *mongo.Collection, bid string, mustFind bool, findOptions *options.FindOneOptionsBuilder) (*Comune, bool, error) {
	// @tpm-schematics:end-region("find-by-pk-signature-section")
	const semLogContext = "comune::find-by-pk"
	// @tpm-schematics:start-region("log-event-section")
	evtTraceLog := log.Trace()
	evtErrLog := log.Error()
	// @tpm-schematics:end-region("log-event-section")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ent := Comune{}

	f := Filter{}
	// @tpm-schematics:start-region("filter-section")
	f.Or().AndEtEqTo(EntityType).AndBidEqTo(bid)
	// @tpm-schematics:end-region("filter-section")
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

func Find(collection *mongo.Collection, f *Filter, withCount bool, findOptions *options.FindOptionsBuilder) (QueryResult, error) {
	const semLogContext = "comune::find"
	fd := f.Build()
	evtTraceLog := log.Trace().Str("filter", util.MustToExtendedJsonString(fd, false, false))
	evtErrLog := log.Error().Str("filter", util.MustToExtendedJsonString(fd, false, false))
	evtTraceLog.Msg(semLogContext)

	qr := QueryResult{}

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
		dto := Comune{}
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

// @tpm-schematics:start-region("bottom-file-section")

func FindByNazioneAndProvinciaAndStatus(collection *mongo.Collection, code_nazione_uic string, provincia string, status string, withCount bool, findOptions *options.FindOptionsBuilder) (QueryResult, error) {
	const semLogContext = "comune::find-by-code_nazione_uic-provincia-status"
	log.Trace().Str("cod_nazione_uic", code_nazione_uic).Str("provincia", provincia).Str("status", status).Msg(semLogContext)
	f := Filter{}
	f.Or().AndEtEqTo(EntityType).AndNazioneEqTo(code_nazione_uic).AndProvinciaEqTo(provincia).AndStatusEqTo(status)
	return Find(collection, &f, withCount, findOptions)
}

// @tpm-schematics:end-region("bottom-file-section")
