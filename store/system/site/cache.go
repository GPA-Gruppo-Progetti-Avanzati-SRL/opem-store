package site

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
	const semLogContext = "opem-mongodb/site/new-cache"
	log.Info().Dur("expiry", expDuration).Dur("purge", purgeInterval).Msg(semLogContext)
	domainCache = cache.New(expDuration, purgeInterval)
}

func NewCacheResolver(coll *mongo.Collection) CacheResolverFunc {
	const SemLogContext = "opem-mongodb/site/new-cache-resolver"
	return func(d, s string) (interface{}, error) {

		// Use the mustFind false for easier handling of DOS cache misses.... Dunno how to handle.... tbd.
		ent, ok, err := FindByDomainAndCode(coll, d, s, false, nil)
		if err != nil {
			log.Error().Err(err).Msg(SemLogContext)
			return nil, err
		}

		// Documento non trovato: ritorna (nil, nil) — non è un errore,
		// è una normale assenza. GetFromCache gestisce il nil senza loggare ERR.
		if !ok {
			return nil, nil
		}

		return ent, nil
	}
}

func GetFromCache(resolver CacheResolver, domain, code string) (*Site, bool) {

	const SemLogContext = "opem-mongodb/site/get-site-from-cache"

	k := CacheKey(domain, code)
	item, ok := domainCache.Get(k)
	if !ok {
		// Cache miss è normale: la prima richiesta o dopo scadenza TTL.
		log.Debug().Str("domain", domain).Str("site", code).Msg(SemLogContext + " - cache miss, querying db")
		var err error
		item, err = resolver.Retrieve(domain, code)
		if err != nil {
			// Errore reale di connessione / query MongoDB
			log.Error().Err(err).Str("domain", domain).Str("site", code).Msg(SemLogContext + " - db error")
			return nil, false
		}
		if item == nil {
			// Documento non presente in MongoDB — non è un errore
			log.Debug().Str("domain", domain).Str("site", code).Msg(SemLogContext + " - not found in db")
			return nil, false
		}
		domainCache.Set(k, item, cache.DefaultExpiration)
	}

	return item.(*Site), true
}

func CacheKey(d, s string) string {
	return fmt.Sprintf("%s#%s", d, s)
}
