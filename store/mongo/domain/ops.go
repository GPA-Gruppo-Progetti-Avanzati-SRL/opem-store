package domain

import (
	"context"
	"time"

	model "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func FindByCode(collection *mongo.Collection, code string, mustFind bool, findOptions *options.FindOneOptionsBuilder) (*model.Domain, bool, error) {
	const semLogContext = "mdb-domain::find-by-code"

	log.Trace().Str("domain", code).Msg(semLogContext)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ent := model.Domain{}

	f := Filter{}
	f.Or().AndBidEqTo(code).AndEtEqTo(EntityType)
	err := collection.FindOne(ctx, f.Build(), findOptions).Decode(&ent)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Error().Err(err).Msg(semLogContext)
		return nil, false, err
	} else {
		if err != nil {
			if mustFind {
				log.Trace().Str("domain", code).Msg(semLogContext + " document not found")
				return nil, false, err
			}
			log.Trace().Str("domain", code).Msg(semLogContext + " document not found but allowed")
			ent.Bid = code
			return nil, false, nil
		} else {
			log.Trace().Str("domain", code).Msg(semLogContext + " document found")
		}
	}

	return &ent, true, nil
}
