package store_test

import (
	"context"
	"testing"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/system/sequence"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-mongo-common/mongolks"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	sequenceDomain = "test"
	sequenceSite   = "test"
	sequenceBid    = "test"
)

func TestSequence(t *testing.T) {

	coll, err := mongolks.GetCollection(context.Background(), "default", sequence.CollectionId)
	require.NoError(t, err)

	rng, err := sequence.NextRange(coll, sequenceDomain, sequenceSite, sequence.WithIncrement(10), sequence.WithCreateIfMissing(true), sequence.WithSeqId(sequenceBid))
	require.NoError(t, err)
	t.Log("range: ", rng)

	seq, _, err := sequence.FindByPk(coll, sequenceDomain, sequenceSite, sequenceBid, true, options.FindOne())
	require.NoError(t, err)
	t.Log("seq: ", seq)

	t.Log("next", rng.Next())
	err = sequence.ReleaseSequenceRange(coll, sequenceDomain, sequenceSite, sequenceBid, rng)
	require.NoError(t, err)

	seq, _, err = sequence.FindByPk(coll, sequenceDomain, sequenceSite, sequenceBid, true, options.FindOne())
	require.NoError(t, err)
	t.Log("seq: ", seq)

	rng, err = sequence.NextRange(coll, sequenceDomain, sequenceSite, sequence.WithIncrement(1), sequence.WithCreateIfMissing(true), sequence.WithSeqId(sequenceBid))
	require.NoError(t, err)
	t.Log("range: ", rng)
}
