package sequence

import (
	"context"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-mongo-common/mongolks"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type BufferedRange struct {
	coll         *mongo.Collection
	domain, site string
	opts         []NextValOption
	rng          Range
}

func NewBufferedRange(collection *mongo.Collection, domain, site string, nextValOpts ...NextValOption) *BufferedRange {
	return &BufferedRange{coll: collection, domain: domain, site: site, opts: nextValOpts}
}

func (br *BufferedRange) Next() (string, error) {
	const semLogContext = "buffered-range::next"

	if br.rng.HasNext() {
		return br.rng.Next()
	}

	coll, err := mongolks.GetCollection(context.Background(), "default", CollectionId)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return "", err
	}

	br.rng, err = NextRange(coll, br.domain, br.site, br.opts...)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return "", err
	}

	return br.rng.Next()
}

func (br *BufferedRange) Release() error {
	const semLogContext = "buffered-range::release"

	options := NextValOptions{CreateIfMissing: true, Increment: 1}
	for _, o := range br.opts {
		o(&options)
	}

	return ReleaseSequenceRange(br.coll, br.domain, br.site, options.SeqId, br.rng)
}
