package domain

import (
	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type CacheResolver interface {
	Retrieve(k string) (interface{}, error)
}

type CacheResolverFunc func(k string) (interface{}, error)

func (f CacheResolverFunc) Retrieve(k string) (interface{}, error) {
	return f(k)
}

var domainCache *cache.Cache

func NewCache(expDuration time.Duration, purgeInterval time.Duration) {
	const semLogContext = "r3ds9-mongodb/domain/new-cache"
	log.Info().Dur("expiry", expDuration).Dur("purge", purgeInterval).Msg(semLogContext)
	domainCache = cache.New(expDuration, purgeInterval)
}

func NewCacheResolver(coll *mongo.Collection) CacheResolverFunc {
	const SemLogContext = "r3ds9-mongodb/domain/new-cache-resolver"
	return func(k string) (interface{}, error) {

		// Use the mustFind false for easier handling of DOS cache misses.... Dunno how to handle.... tbd.
		ent, ok, err := FindByCode(coll, k, false, nil)
		if err != nil {
			log.Error().Err(err).Msg(SemLogContext)
			return nil, err
		}

		if !ok {
			return nil, mongo.ErrNoDocuments
		}

		return ent, nil
	}
}

func GetFromCache(resolver CacheResolver, code string) (*Domain, bool) {

	const SemLogContext = "r3ds9-mongodb/domain/get-domain-from-cache"

	var err error
	item, ok := domainCache.Get(code)
	if !ok {
		log.Warn().Str("k", code).Msg(SemLogContext + " cache miss")
		item, err = resolver.Retrieve(code)
		if err != nil {
			return nil, false
		}

		domainCache.Set(code, item, cache.DefaultExpiration)
	}

	return item.(*Domain), true
}
