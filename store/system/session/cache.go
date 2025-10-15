package session

import (
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CacheResolver interface {
	Retrieve(k string) (interface{}, error)
}

type CacheResolverFunc func(k string) (interface{}, error)

func (f CacheResolverFunc) Retrieve(k string) (interface{}, error) {
	return f(k)
}

var sessionCache *cache.Cache

func NewCache(expDuration time.Duration, purgeInterval time.Duration) {
	const semLogContext = "r3ds9-mongodb/session/new-cache"
	log.Info().Dur("expiry", expDuration).Dur("purge", purgeInterval).Msg(semLogContext)
	sessionCache = cache.New(expDuration, purgeInterval)
}

func NewCacheResolver(coll *mongo.Collection) CacheResolverFunc {
	const SemLogContext = "r3ds9-mongodb/session/new-cache-resolver"
	return func(k string) (interface{}, error) {

		// Use the mustFind false for easier handling of DOS cache misses.... Dunno how to handle.... tbd.
		ent, err := FindBySId(coll, k, false, nil)
		if err != nil {
			log.Error().Err(err).Msg(SemLogContext)
			return nil, err
		}

		if ent == nil {
			return nil, mongo.ErrNoDocuments
		}

		return ent, nil
	}
}

func InvalidateSession(code string) {
	sessionCache.Delete(code)
}

func GetFromCache(resolver CacheResolver, code string) (*Session, bool) {

	const SemLogContext = "r3ds9-mongodb/session/get-session-from-cache"

	var err error
	item, ok := sessionCache.Get(code)
	if !ok {
		log.Warn().Str("k", code).Msg(SemLogContext + " cache miss")
		item, err = resolver.Retrieve(code)
		if err != nil {
			return nil, false
		}

		sessionCache.Set(code, item, cache.DefaultExpiration)
	}

	return item.(*Session), true
}
