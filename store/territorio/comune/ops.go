package comune

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindByBid(collection *mongo.Collection, bid string, mustFind bool, findOptions *options.FindOneOptions) (*Comune, bool, error) {
	const semLogContext = "comune::find-by-bid"

	log.Trace().Str("_bid", bid).Msg(semLogContext)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ent := Comune{}

	f := Filter{}
	f.Or().AndEtEqTo(EntityType).AndBidEqTo(bid)
	err := collection.FindOne(ctx, f.Build(), findOptions).Decode(&ent)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		log.Error().Err(err).Msg(semLogContext)
		return &ent, false, err
	} else {
		if err != nil {
			if mustFind {
				log.Trace().Str("_bid", bid).Msg(semLogContext + " document not found")
				return &ent, false, err
			}

			log.Trace().Str("_bid", bid).Msg(semLogContext + " document not found but allowed")
			ent.Bid = bid
			return &ent, false, nil
		} else {
			log.Trace().Str("_bid", bid).Msg(semLogContext + " document found")
		}
	}

	return &ent, true, nil
}

func FindByNazioneAndProvinciaAndStatus(collection *mongo.Collection, code_nazione_uic string, provincia string, status string, withCount bool, findOptions *options.FindOptions) (QueryResult, error) {
	const semLogContext = "comune::find-by-code_nazione_uic-provincia-status"
	log.Trace().Str("cod_nazione_uic", code_nazione_uic).Str("provincia", provincia).Str("status", status).Msg(semLogContext)
	f := Filter{}
	f.Or().AndEtEqTo(EntityType).AndCodeUicNazioneEqTo(code_nazione_uic).AndCodeProvinciaEqTo(provincia).AndStatusEqTo(status)
	return Find(collection, &f, withCount, findOptions)
}

func Find(collection *mongo.Collection, f *Filter, withCount bool, findOptions *options.FindOptions) (QueryResult, error) {
	const semLogContext = "comune::find"

	qr := QueryResult{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fd := f.Build()

	if withCount {
		countDocsOptions := options.CountOptions{}
		nr, err := collection.CountDocuments(ctx, fd, &countDocsOptions)
		if err != nil {
			log.Error().Err(err).Msg(semLogContext)
			return qr, err
		}

		qr.Records = int(nr)
	}

	cur, err := collection.Find(ctx, f.Build(), findOptions)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
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
