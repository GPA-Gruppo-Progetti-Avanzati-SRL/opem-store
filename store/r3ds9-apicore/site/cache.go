package site

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type CacheResolver interface {
	Retrieve(d, s string) (interface{}, error)
}

type CacheResolverFunc func(d, s string) (interface{}, error)

func (f CacheResolverFunc) Retrieve(d, s string) (interface{}, error) {
	return f(d, s)
}

var domainCache *cache.Cache

func NewCache(expDuration time.Duration, purgeInterval time.Duration) {
	const semLogContext = "r3ds9-mongodb/site/new-cache"
	log.Info().Dur("expiry", expDuration).Dur("purge", purgeInterval).Msg(semLogContext)
	domainCache = cache.New(expDuration, purgeInterval)
}

func NewCacheResolver(coll *mongo.Collection) CacheResolverFunc {
	const SemLogContext = "r3ds9-mongodb/site/new-cache-resolver"
	return func(d, s string) (interface{}, error) {

		// Use the mustFind false for easier handling of DOS cache misses.... Dunno how to handle.... tbd.
		ent, ok, err := FindByDomainAndCode(coll, d, s, false, nil)
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

func GetFromCache(resolver CacheResolver, domain, code string) (*Site, bool) {

	const SemLogContext = "r3ds9-mongodb/site/get-site-from-cache"
	var err error

	k := CacheKey(domain, code)
	item, ok := domainCache.Get(k)
	if !ok {
		log.Warn().Str("k", k).Str("domain", domain).Str("site", code).Msgf(SemLogContext + " cache miss")
		item, err = resolver.Retrieve(domain, code)
		if err != nil {
			log.Error().Err(err).Msg(SemLogContext)
			return nil, false
		}

		domainCache.Set(k, item, cache.DefaultExpiration)
	}

	return item.(*Site), true
}

func CacheKey(d, s string) string {
	return fmt.Sprintf("%s#%s", d, s)
}
