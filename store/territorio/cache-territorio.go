package territorio

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/territorio/comune"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/territorio/nazione"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/territorio/provincia"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-mongo-common/mongolks"
	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type CacheResolver interface {
	Retrieve(k string) (interface{}, error)
}

type CacheResolverFunc func(k string) (interface{}, error)

func (f CacheResolverFunc) Retrieve(k string) (interface{}, error) {
	return f(k)
}

var theCache *cache.Cache
var theCacheResolverFunc CacheResolverFunc

func NewCache(expDuration time.Duration, purgeInterval time.Duration) error {
	const semLogContext = "territorio::new-cache"
	log.Info().Dur("expiry", expDuration).Dur("purge", purgeInterval).Msg(semLogContext)
	theCache = cache.New(expDuration, purgeInterval)

	coll, err := mongolks.GetCollection(context.Background(), "default", comune.CollectionId)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return err
	} else {
		theCacheResolverFunc = NewCacheResolver(coll)
	}

	return nil
}

func NewCacheResolver(coll *mongo.Collection) CacheResolverFunc {
	const semLogContext = "territorio::new-cache-resolver"
	return func(k string) (interface{}, error) {
		var err error

		keyComponents := strings.Split(k, "::")
		if len(keyComponents) != 2 {
			err = errors.New("unsupported key")
			log.Error().Err(err).Str("k", k).Msgf(semLogContext)
			return nil, err
		}

		var ent interface{}
		switch keyComponents[0] {
		case comune.EntityType:
			ent, _, err = comune.FindByPk(coll, keyComponents[1], true, options.FindOne())
		case provincia.EntityType:
			ent, _, err = provincia.FindByPk(coll, keyComponents[1], true, options.FindOne())
		case nazione.EntityType:
			ent, _, err = nazione.FindByPk(coll, keyComponents[1], true, options.FindOne())
		}

		if err != nil {
			log.Error().Err(err).Msg(semLogContext)
			return nil, err
		}

		if ent == nil {
			return nil, mongo.ErrNoDocuments
		}

		return ent, nil
	}
}

func GetFromCache(resolver CacheResolver, et, bid string) (interface{}, bool) {

	const semLogContext = "territorio::get-from-cache"

	var err error
	key := fmt.Sprintf("%s::%s", et, bid)
	item, ok := theCache.Get(key)
	if !ok {
		log.Warn().Str("k", key).Msgf(semLogContext + " cache miss")
		item, err = resolver.Retrieve(key)
		if err != nil {
			return nil, false
		}

		theCache.Set(key, item, cache.DefaultExpiration)
	}

	return item, true
}

func GetNomeComuneFromCache(bid string) (string, bool) {
	const semLogContext = "territorio::get-nome-comune-from-cache"
	cachedObj, ok := GetFromCache(theCacheResolverFunc, comune.EntityType, bid)
	if ok {
		if comune, ok := cachedObj.(*comune.Comune); ok {
			return comune.Name, true
		} else {
			err := errors.New("error in resolving cache")
			log.Error().Err(err).Str("cached-obj-type", fmt.Sprintf("%T", cachedObj)).Msg(semLogContext)
		}
	}

	return "", false
}
