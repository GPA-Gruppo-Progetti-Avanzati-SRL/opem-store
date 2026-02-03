package card_test

import (
	"context"
	"testing"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/card/card"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-mongo-common/mongolks"
	"github.com/stretchr/testify/require"
)

func TestCard(t *testing.T) {

	coll, err := mongolks.GetCollection(context.Background(), "default", CardsCollectionId)
	require.NoError(t, err)

	cardRanges := []commons.ValueRange{
		{
			From: "00679822",
			To:   "00679823",
		},
		{
			From: "NOT",
			To:   "00679823",
		},
		{
			From: "00687145",
			To:   "00687147",
		},
	}

	fndr, err := card.NewCardFinderByBidRange(coll, "card", "10000", cardRanges)
	require.NoError(t, err)

	crd, err := fndr.Next()
	for err == nil && crd != nil {
		t.Log("found: ", crd.Bid)
		crd, err = fndr.Next()
	}

	require.NoError(t, err)
}
