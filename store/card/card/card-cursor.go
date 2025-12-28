package card

import (
	"context"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-mongo-common/mongolks"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Cursor struct {
	coll *mongo.Collection
	crs  *mongo.Cursor
}

func NewCursor(c *mongo.Collection, filter Filter, useView bool) (Cursor, error) {
	const semLogContext = "card::new-cursor"

	lks, err := mongolks.GetLinkedService(context.Background(), "default")
	if err != nil {
		return Cursor{}, err
	}

	var crs *mongo.Cursor
	if !useView {
		crs, err = c.Find(context.Background(), filter)
	} else {
		crs, err = cursorByAggregationView(c, lks.GetCollectionsCfg(), &filter, nil)
	}

	if err != nil {
		return Cursor{}, err
	}

	return Cursor{
		coll: c,
		crs:  crs,
	}, nil
}

func (cf *Cursor) Next() (*Card, error) {
	const semLogContext = "card-cursor::next"
	if cf.crs == nil {
		return nil, nil
	}

	if cf.crs.Next(context.Background()) {
		dto := Card{}
		err := cf.crs.Decode(&dto)
		if err != nil {
			return nil, err
		}

		return &dto, nil
	}

	if cf.crs.Err() != nil {
		err := cf.crs.Err()
		cf.crs = nil
		log.Error().Err(err).Msg(semLogContext)
		return nil, err
	}

	log.Info().Msg(semLogContext + " - cursor empty")

	return nil, nil
}
