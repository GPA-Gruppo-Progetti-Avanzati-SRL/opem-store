package prodotto

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func findByDomainAndSiteAndStatus(collection *mongo.Collection, domain string, site string, status string, findOptions *options.FindOptionsBuilder) ([]Prodotto, error) {
	const semLogContext = "prodotto::find-by-domain-site-status"

	log.Trace().Str("domain", domain).Str("site", site).Str("status", status).Msg(semLogContext)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var ents []Prodotto

	f := Filter{}
	f.Or().AndEtEqTo(EntityType).AndDomainEqTo(domain).AndSiteIn([]string{site, "*"}).AndStatusEqTo(status)
	crs, err := collection.Find(ctx, f.Build(), findOptions)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return nil, err
	}

	err = crs.All(ctx, &ents)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return nil, err
	}

	return ents, nil
}
