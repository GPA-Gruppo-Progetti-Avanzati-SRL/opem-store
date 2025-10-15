package box

import (
	"context"
	"errors"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-mongo-common/util"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func AddNote(coll *mongo.Collection, domain, site, bidMagazzino, bidBox string, aNote commons.Note) error {
	const semLogContext = "box::add-note"

	f := Filter{}
	f.Or().AndDomainEqTo(domain).AndSiteEqTo(site).AndEtEqTo(EntityType).AndBidMagazzinoEqTo(bidMagazzino).AndBidEqTo(bidBox)
	fd := f.Build()
	log.Trace().Str("add-note-filter", util.MustToExtendedJsonString(fd, false, false)).Msg(semLogContext)

	updOpts := []UpdateOption{
		UpdateWithAddNote(aNote),
	}
	updDoc := GetUpdateDocumentFromOptions(updOpts...)
	ud := updDoc.Build()
	log.Trace().Str("add-note-update-document", util.MustToExtendedJsonString(ud, false, false)).Msg(semLogContext)

	resp, err := coll.UpdateOne(context.Background(), fd, ud)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return err
	}

	if resp.ModifiedCount == 0 {
		err = errors.New("no document updated matched")
		log.Error().Err(err).Msg(semLogContext)
		return err
	}

	return nil
}
