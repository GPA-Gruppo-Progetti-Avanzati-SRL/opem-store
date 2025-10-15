package focalpoint

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func findByDomainAndSiteAndStatus(collection *mongo.Collection, domain string, site string, status string, focalPointBid string, findOptions *options.FindOptionsBuilder) ([]FocalPoint, error) {
	const semLogContext = "focal-point::find-by-domain-site-status-with-bid-constraint"

	log.Trace().Str("domain", domain).Str("site", site).Str("status", status).Msg(semLogContext)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var ents []FocalPoint

	f := Filter{}
	f.Or().AndEtEqTo(EntityType).AndDomainEqTo(domain).AndSiteEqTo(site).AndStatusEqTo(status).AndBidEqTo(focalPointBid)
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
