package provincia

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func findByServizioAndNumeroAndNdgAndGruppoAndTokenId(collection *mongo.Collection, code_nazione_uic string, status string, findOptions *options.FindOptions) ([]Provincia, error) {
	const semLogContext = "provincia::find-by-code_nazione_uic"

	log.Trace().Str("cod_nazione_uic", code_nazione_uic).Str("status", status).Msg(semLogContext)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var ents []Provincia

	f := Filter{}
	f.Or().AndEtEqTo(EntityType).AndCodeEqTo(code_nazione_uic).AndStatusEqTo(status)
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
