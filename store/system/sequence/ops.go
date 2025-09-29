package sequence

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NextVal(collection *mongo.Collection, nextValOpts ...NextValOption) (int, error) {
	const semLogContext = "sequence::next-val"

	opts := NextValOptions{CreateIfMissing: true, Increment: 1}
	for _, o := range nextValOpts {
		o(&opts)
	}

	if opts.SeqId == "" {
		panic(fmt.Errorf("sequence missing core params - id: %s", opts.SeqId))
	}

	f := Filter{}
	f.Or().AndBidEqTo(opts.SeqId).AndEtEqTo(EntityType)
	fd := f.Build()

	updOpts := []UpdateOption{UpdateWithIncrementValue(opts.Increment), UpdateWith_bid(opts.SeqId), UpdateWith_et(EntityType)}
	upd := GetUpdateDocumentFromOptions(updOpts...)
	findOneAndUpdateOpts := options.FindOneAndUpdate()
	if opts.CreateIfMissing {
		findOneAndUpdateOpts.SetUpsert(true)
	}
	findOneAndUpdateOpts.SetReturnDocument(options.After)
	res := collection.FindOneAndUpdate(context.Background(), fd, upd.Build(), findOneAndUpdateOpts)
	if res.Err() != nil {
		log.Error().Err(res.Err()).Msg(semLogContext)
		return 0, res.Err()
	}

	var seq Sequence
	err := res.Decode(&seq)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return 0, err
	}

	return int(seq.Value), nil
}
