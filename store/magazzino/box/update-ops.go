package box

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/file/file"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-mongo-common/util"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
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

func Update4ReintegroMagazzino(coll *mongo.Collection, domain, site, bidMagazzino, bidBox string, who string, valueRange commons.ValueRange) error {
	const semLogContext = "box::update-4-reintegro-magazzino"

	f := Filter{}
	f.Or().AndDomainEqTo(domain).AndSiteEqTo(site).AndEtEqTo(EntityType).AndBidMagazzinoEqTo(bidMagazzino).AndBidEqTo(bidBox)
	fd := f.Build()
	log.Trace().Str("update-4-reintegro-magazzino-filter", util.MustToExtendedJsonString(fd, false, false)).Msg(semLogContext)

	evt := commons.Event{
		Type: EventReintegroMagazzino,
		Who:  who,
		Text: valueRange.String(),
		When: bson.NewDateTimeFromTime(time.Now()),
	}

	updOpts := []UpdateOption{
		UpdateWithAddEvent(evt),
		UpdateWithStatusStatus(StatusInAttesaDiInvio),
		UpdateWithSysInfoModifiedAt(),
		UpdateWithCard_bids_range(&valueRange),
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

func Update4RichiestaCarte(coll *mongo.Collection, domain, site, bidMagazzino, bidBox string, who string, files []file.File) error {
	const semLogContext = "box::update-4-richiesta_carte"

	f := Filter{}
	f.Or().AndDomainEqTo(domain).AndSiteEqTo(site).AndEtEqTo(EntityType).AndBidMagazzinoEqTo(bidMagazzino).AndBidEqTo(bidBox)
	fd := f.Build()
	log.Trace().Str("update-4-richiesta-carte-filter", util.MustToExtendedJsonString(fd, false, false)).Msg(semLogContext)

	var sb strings.Builder
	for _, f := range files {
		sb.WriteString(fmt.Sprintf("[%s]: %s ", f.BlobBucket, f.BlobKey))
	}

	evt := commons.Event{
		Type: EventRichiestaCarte,
		Who:  who,
		Text: sb.String(),
		When: bson.NewDateTimeFromTime(time.Now()),
	}

	updOpts := []UpdateOption{
		UpdateWithAddEvent(evt),
		UpdateWithStatusStatus(StatusInAttesaDiConferma),
		UpdateWithSysInfoModifiedAt(),
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

func Update4RendicontazioneRichiestaCarte(coll *mongo.Collection, domain, site, bidMagazzino, bidBox string, who string, fl *file.File) error {
	const semLogContext = "box::update-4-rendicontazione-richiesta_carte"

	f := Filter{}
	f.Or().AndDomainEqTo(domain).AndSiteEqTo(site).AndEtEqTo(EntityType).AndBidMagazzinoEqTo(bidMagazzino).AndBidEqTo(bidBox)
	fd := f.Build()
	log.Trace().Str("update-4-rendicontazione-richiesta-carte-filter", util.MustToExtendedJsonString(fd, false, false)).Msg(semLogContext)

	evt := commons.Event{
		Type: EventRichiestaCarte,
		Who:  who,
		Text: fmt.Sprintf("[%s]: %s ", fl.BlobBucket, fl.BlobKey),
		When: bson.NewDateTimeFromTime(time.Now()),
	}

	updOpts := []UpdateOption{
		UpdateWithAddEvent(evt),
		UpdateWithStatusStatus(StatusConfermatoDisponibilePerLaConsegna),
		UpdateWithSysInfoModifiedAt(),
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
