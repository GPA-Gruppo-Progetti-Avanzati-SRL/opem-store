package sequence

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func SprintfNextVal(collection *mongo.Collection, domain, site string, nextValOpts ...NextValOption) (string, error) {
	const semLogContext = "sequence::formatted-next-val"
	rng, err := NextRange(collection, domain, site, nextValOpts...)
	if err != nil {
		return "", err
	}

	if rng.from != rng.to {
		log.Warn().Interface("range", rng).Msg(semLogContext + " - call not appropriate for ranges")
	}

	return fmt.Sprintf(rng.format, rng.to), nil
}

func NextRange(collection *mongo.Collection, domain, site string, nextValOpts ...NextValOption) (Range, error) {
	const semLogContext = "sequence::next-val"

	opts := NextValOptions{CreateIfMissing: true, Increment: 1}
	for _, o := range nextValOpts {
		o(&opts)
	}

	if opts.SeqId == "" {
		panic(fmt.Errorf("sequence missing core params - id: %s", opts.SeqId))
	}

	f := Filter{}
	f.Or().AndBidEqTo(opts.SeqId).AndEtEqTo(EntityType).AndDomainEqTo(domain).AndSiteEqTo(site)
	fd := f.Build()

	updOpts := []UpdateOption{UpdateWithIncrementValue(opts.Increment), UpdateWith_bid(opts.SeqId), UpdateWith_et(EntityType), UpdateWithDomain(domain), UpdateWithSite(site)}
	upd := GetUpdateDocumentFromOptions(updOpts...)
	findOneAndUpdateOpts := options.FindOneAndUpdate()
	if opts.CreateIfMissing {
		findOneAndUpdateOpts.SetUpsert(true)
	}
	findOneAndUpdateOpts.SetReturnDocument(options.After)
	res := collection.FindOneAndUpdate(context.Background(), fd, upd.Build(), findOneAndUpdateOpts)
	if res.Err() != nil {
		log.Error().Err(res.Err()).Msg(semLogContext)
		return Range{}, res.Err()
	}

	var seq Sequence
	err := res.Decode(&seq)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return Range{}, err
	}

	return Range{from: seq.Value - opts.Increment + 1, to: seq.Value, format: seq.FormatSpecifier(), current: -1}, nil
}
