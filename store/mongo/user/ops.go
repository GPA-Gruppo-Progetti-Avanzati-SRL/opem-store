package user

import (
	"context"
	"errors"
	"time"

	model "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func FindByNickname(collection *mongo.Collection, nickname string, mustFind bool, findOptions *options.FindOneOptionsBuilder) (*model.User, error) {
	const semLogContext = "mdb-user::find-by-nickname"

	log.Trace().Str("nickname", nickname).Msg(semLogContext)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ent := model.User{}

	f := Filter{}
	f.Or().AndNicknameEqTo(nickname)
	err := collection.FindOne(ctx, f.Build(), findOptions).Decode(&ent)
	if err != nil {
		isNotFound := errors.Is(err, mongo.ErrNoDocuments)
		if !isNotFound || mustFind {
			log.Error().Err(err).Msg(semLogContext)
			return nil, err
		}
		log.Trace().Str("nickname", nickname).Msgf("%s - document not found", semLogContext)
		return nil, nil
	}

	return &ent, nil
}

func FindByHexOid(collection *mongo.Collection, userId string, mustFind bool, findOptions *options.FindOneOptionsBuilder) (*model.User, error) {

	const SemLogContext = "opem-core/user/find-by-object-id"
	log.Trace().Str("userId", userId).Msg(SemLogContext)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ent := model.User{}

	f := Filter{}
	f.Or().AndHexOIdEqTo(userId)
	err := collection.FindOne(ctx, f.Build(), findOptions).Decode(&ent)
	if err != nil && (err != mongo.ErrNoDocuments || (err == mongo.ErrNoDocuments && mustFind)) {
		log.Error().Err(err).Msg(SemLogContext)
		return nil, err
	} else {
		if err != nil {
			log.Trace().Str("userId", userId).Msgf("%s - document not found", SemLogContext)
			return nil, err
		} else {
			log.Trace().Str("userId", userId).Msgf("%s - document found", SemLogContext)
		}
	}

	return &ent, nil
}
