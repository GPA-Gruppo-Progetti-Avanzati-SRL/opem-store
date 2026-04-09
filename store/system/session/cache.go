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
	const semLogContext = "opem-mongodb/session/new-cache"
	log.Info().Dur("expiry", expDuration).Dur("purge", purgeInterval).Msg(semLogContext)
	sessionCache = cache.New(expDuration, purgeInterval)
}

func NewCacheResolver(coll *mongo.Collection) CacheResolverFunc {
	const SemLogContext = "opem-mongodb/session/new-cache-resolver"
	return func(k string) (interface{}, error) {
		ent, err := FindBySId(coll, k, false, nil)
		if err != nil {
			log.Error().Err(err).Msg(SemLogContext)
			return nil, err
		}
		// Non trovato: non è un errore, è una normale assenza
		if ent == nil {
			return nil, nil
		}
		return ent, nil
	}
}

func InvalidateSession(code string) {
	sessionCache.Delete(code)
}

func GetFromCache(resolver CacheResolver, code string) (*Session, bool) {
	const SemLogContext = "opem-mongodb/session/get-session-from-cache"

	item, ok := sessionCache.Get(code)
	if !ok {
		log.Debug().Str("code", code).Msg(SemLogContext + " - cache miss, querying db")
		var err error
		item, err = resolver.Retrieve(code)
		if err != nil {
			log.Error().Err(err).Str("code", code).Msg(SemLogContext + " - db error")
			return nil, false
		}
		if item == nil {
			log.Debug().Str("code", code).Msg(SemLogContext + " - not found in db")
			return nil, false
		}
		sessionCache.Set(code, item, cache.DefaultExpiration)
	}

	return item.(*Session), true
}
