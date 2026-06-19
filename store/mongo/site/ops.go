package site

import (
	"context"
	"time"

	model "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func FindByDomainAndCode(collection *mongo.Collection, domain, code string, mustFind bool, findOptions *options.FindOneOptionsBuilder) (*model.Site, bool, error) {
	const semLogContext = "mdb-site::find-by-domain-and-code"

	log.Trace().Str("domain", code).Msg(semLogContext)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ent := model.Site{}

	f := Filter{}
	f.Or().AndBidEqTo(code).AndEtEqTo(EntityType).AndDomainEqTo(domain)
	err := collection.FindOne(ctx, f.Build(), findOptions).Decode(&ent)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Error().Err(err).Msg("site find operation")
		return nil, false, err
	} else {
		if err != nil {
			if mustFind {
				log.Trace().Str("site", code).Str("domain", domain).Msg(semLogContext + " document not found")
				return nil, false, err
			}
			log.Trace().Str("site", code).Str("domain", domain).Msg(semLogContext + " document not found but allowed")
			ent.Bid = code
			return nil, false, nil
		} else {
			log.Trace().Str("site", code).Str("domain", domain).Msg(semLogContext + " document found")
		}
	}

	return &ent, true, nil
}
