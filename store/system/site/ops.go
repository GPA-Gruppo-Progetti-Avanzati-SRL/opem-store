package site

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func Find(collection *mongo.Collection, f Filter, findOptions *options.FindOptions, withCount bool) (QueryResult, error) {
	const semLogContext = "mdb-site::find"

	qr := QueryResult{}

	log.Trace().Msg(semLogContext)

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
		dto := Site{}
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

func FindByDomainAndCode(collection *mongo.Collection, domain, code string, mustFind bool, findOptions *options.FindOneOptions) (*Site, bool, error) {
	const semLogContext = "mdb-site::find-by-domain-and-code"

	log.Trace().Str("domain", code).Msg(semLogContext)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ent := Site{}

	f := Filter{}
	f.Or().AndCodeEqTo(code).AndDomainEqTo(domain)
	err := collection.FindOne(ctx, f.Build(), findOptions).Decode(&ent)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Error().Err(err).Msg("site find operation")
		return &ent, false, err
	} else {
		if err != nil {
			if mustFind {
				log.Trace().Str("site", code).Str("domain", domain).Msg(semLogContext + " document not found")
				return &ent, false, err
			}

			log.Trace().Str("site", code).Str("domain", domain).Msg(semLogContext + " document not found but allowed")
			ent.Code = code
			return &ent, false, nil
		} else {
			log.Trace().Str("site", code).Str("domain", domain).Msg(semLogContext + " document found")
		}
	}

	return &ent, true, nil
}
