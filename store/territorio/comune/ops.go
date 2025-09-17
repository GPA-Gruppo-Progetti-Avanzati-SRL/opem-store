package comune

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func findByNazioneAndProvinciaAndStatus(collection *mongo.Collection, code_nazione_uic string, provincia string, status string, findOptions *options.FindOptions) ([]Comune, error) {
	const semLogContext = "comune::find-by-code_nazione_uic-provincia-status"

	log.Trace().Str("cod_nazione_uic", code_nazione_uic).Str("provincia", provincia).Str("status", status).Msg(semLogContext)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var ents []Comune

	f := Filter{}
	f.Or().AndEtEqTo(EntityType).AndCodeUicNazioneEqTo(code_nazione_uic).AndCodeProvinciaEqTo(provincia).AndStatusEqTo(status)
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
