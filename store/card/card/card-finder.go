package card

import (
	"context"
	"errors"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-mongo-common/util"
	"github.com/rs/zerolog/log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type ByBidRangesFinder struct {
	domain          string
	site            string
	cardRanges      []commons.ValueRange
	currentRangeNdx int

	coll *mongo.Collection
	crs  *mongo.Cursor
}

func NewCardFinder(c *mongo.Collection, domain, site string, r []commons.ValueRange) (ByBidRangesFinder, error) {
	if len(r) == 0 {
		return ByBidRangesFinder{}, errors.New("no card ranges provided")
	}

	return ByBidRangesFinder{
		domain:          domain,
		site:            site,
		coll:            c,
		cardRanges:      r,
		currentRangeNdx: -1,
		crs:             nil,
	}, nil
}

func (cf *ByBidRangesFinder) hasNextRange() bool {
	return cf.currentRangeNdx+1 < len(cf.cardRanges)
}

func (cf *ByBidRangesFinder) nextRange() commons.ValueRange {
	if cf.currentRangeNdx+1 >= len(cf.cardRanges) {
		return commons.ValueRange{}
	}

	cf.currentRangeNdx++
	return cf.cardRanges[cf.currentRangeNdx]
}

func (cf *ByBidRangesFinder) Next() (*Card, error) {

	crd, err := cf.cursorNext()
	if crd == nil {
		crd, err = cf.executeQuery()
	}

	if err != nil {
		return nil, err
	}

	return crd, nil
}

func (cf *ByBidRangesFinder) cursorNext() (*Card, error) {
	const semLogContext = "card-finder::cursor-next"

	var err error
	if cf.crs == nil {
		log.Info().Msg(semLogContext + " - cursor exhausted")
		return nil, nil
	}

	if cf.crs.Next(context.Background()) {
		dto := Card{}
		err = cf.crs.Decode(&dto)
		if err != nil {
			return nil, err
		}

		return &dto, nil
	}

	cf.crs = nil
	return nil, nil
}

func (cf *ByBidRangesFinder) executeQuery() (*Card, error) {
	const semLogContext = "card-finder::execute-query"

	for cf.hasNextRange() {
		f := Filter{}
		f.Or().AndDomainEqTo(cf.domain).AndSiteEqTo(cf.site).AndEtEqTo(EntityType).AndBidInRange(cf.nextRange())
		fd := f.Build()
		log.Info().Str("filter", util.MustToExtendedJsonString(fd, false, false)).Msg(semLogContext)

		fo := options.Find()
		cur, err := cf.coll.Find(context.Background(), fd, fo)
		if err != nil {
			log.Error().Err(err).Msg(semLogContext)
			return nil, err
		}

		if cur.Next(context.Background()) {
			cf.crs = cur

			dto := Card{}
			err = cur.Decode(&dto)
			if err != nil {
				return nil, err
			}

			return &dto, nil
		} else {
			log.Info().Msg(semLogContext + " - cursor empty")
		}
	}

	cf.crs = nil
	return nil, nil
}
