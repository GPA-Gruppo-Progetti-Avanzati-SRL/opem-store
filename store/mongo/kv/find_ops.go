package kv

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

func FindByPk(collection *mongo.Collection, mustFind bool, findOptions *options.FindOneOptionsBuilder) (*model.KeyValuePackage, bool, error) {
	const semLogContext = "kv::find-by-pk"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ent := model.KeyValuePackage{}

	f := Filter{}
	fd := f.Build()
	log.Trace().Str("filter", util.MustToExtendedJsonString(fd, false, false)).Msg(semLogContext)
	err := collection.FindOne(ctx, fd, findOptions).Decode(&ent)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		log.Error().Err(err).Msg(semLogContext)
		return nil, false, err
	} else {
		if err != nil {
			if mustFind {
				log.Trace().Msg(semLogContext + " document not found")
				return nil, false, err
			}
			log.Trace().Msg(semLogContext + " document not found but allowed")
			return nil, false, nil
		} else {
			log.Trace().Msg(semLogContext + " document found")
		}
	}

	return &ent, true, nil
}

func Find(collection *mongo.Collection, f *Filter, withCount bool, findOptions *options.FindOptionsBuilder) (model.KVQueryResult, error) {
	const semLogContext = "kv::find"
	fd := f.Build()
	log.Trace().Str("filter", util.MustToExtendedJsonString(fd, false, false)).Msg(semLogContext)

	qr := model.KVQueryResult{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if withCount {
		countDocsOptions := options.Count()
		nr, err := collection.CountDocuments(ctx, fd, countDocsOptions)
		if err != nil {
			log.Error().Err(err).Msg(semLogContext)
			return qr, err
		}
		qr.Records = int(nr)
	}

	cur, err := collection.Find(ctx, fd, findOptions)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return qr, err
	}

	for cur.Next(context.Background()) {
		dto := model.KeyValuePackage{}
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
