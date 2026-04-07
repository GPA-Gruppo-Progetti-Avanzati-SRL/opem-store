package acl

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// CacheResolver è l'interfaccia che il cache usa per caricare da DB in caso di miss.
type CacheResolver interface {
	Retrieve(ctx context.Context, domain, site string, roles []string) (interface{}, error)
}

// CacheResolverFunc implementa CacheResolver come funzione.
type CacheResolverFunc func(ctx context.Context, domain, site string, roles []string) (interface{}, error)

func (f CacheResolverFunc) Retrieve(ctx context.Context, domain, site string, roles []string) (interface{}, error) {
	return f(ctx, domain, site, roles)
}

var aclCache *cache.Cache

// NewCache inizializza (o re-inizializza) la cache ACL package-level.
// Chiamata da store-mongo::New con gli stessi TTL delle altre cache.
func NewCache(expDuration time.Duration, purgeInterval time.Duration) {
	const semLogContext = "opem-store/acl/new-cache"
	log.Info().Dur("expiry", expDuration).Dur("purge", purgeInterval).Msg(semLogContext)
	aclCache = cache.New(expDuration, purgeInterval)
}

// NewCacheResolver restituisce un CacheResolverFunc che risolve le capabilities
// direttamente da MongoDB tramite ResolveCapabilities.
func NewCacheResolver(coll *mongo.Collection) CacheResolverFunc {
	const semLogContext = "opem-store/acl/new-cache-resolver"
	return func(ctx context.Context, domain, site string, roles []string) (interface{}, error) {
		result, err := ResolveCapabilities(ctx, coll, domain, site, roles)
		if err != nil {
			log.Error().Err(err).Str("domain", domain).Str("site", site).
				Msg(semLogContext + " - resolve error")
			return nil, err
		}
		return result, nil
	}
}

// GetFromCache cerca le capabilities nella cache; in caso di miss chiama il resolver,
// salva il risultato e lo ritorna. Segue il pattern identico a user, site, domain.
func GetFromCache(resolver CacheResolver, ctx context.Context, domain, site string, roles []string) (map[string]map[string][]CapDef, bool) {
	const semLogContext = "opem-store/acl/get-from-cache"

	k := CacheKey(domain, site, roles)
	item, ok := aclCache.Get(k)
	if !ok {
		log.Debug().Str("k", k).Msg(semLogContext + " - cache miss")
		var err error
		item, err = resolver.Retrieve(ctx, domain, site, roles)
		if err != nil {
			log.Error().Err(err).Str("k", k).Msg(semLogContext + " - resolver error")
			return nil, false
		}
		aclCache.Set(k, item, cache.DefaultExpiration)
	}

	if item == nil {
		return nil, false
	}
	return item.(map[string]map[string][]CapDef), true
}

// CacheKey genera una chiave deterministica per (domain, site, roles).
// I ruoli vengono ordinati così la chiave è invariante rispetto al loro ordine.
func CacheKey(domain, site string, roles []string) string {
	sorted := make([]string, len(roles))
	copy(sorted, roles)
	sort.Strings(sorted)
	return fmt.Sprintf("%s#%s#%s", domain, site, strings.Join(sorted, ","))
}

// InvalidateCapabilities rimuove dalla cache una entry specifica.
// Da chiamare quando cambiano i dati ACL per domain+site+roles.
func InvalidateCapabilities(domain, site string, roles []string) {
	if aclCache != nil {
		aclCache.Delete(CacheKey(domain, site, roles))
	}
}

// ── Cache per ResolveCapabilitiesPerSite ──────────────────────────────────────

// perSiteCacheKey genera la chiave per la cache per-site.
// Usa il prefisso "__persite__" per non collidere con le chiavi site-agnostiche
// (che usano site="" → "domain##roles").
func perSiteCacheKey(domain string, roles []string) string {
	return CacheKey(domain, "__persite__", roles)
}

// NewSiteAwareCacheResolver restituisce un CacheResolverFunc che risolve le
// capabilities per site tramite ResolveCapabilitiesPerSite.
func NewSiteAwareCacheResolver(coll *mongo.Collection) CacheResolverFunc {
	const semLogContext = "opem-store/acl/new-site-aware-cache-resolver"
	return func(ctx context.Context, domain, _ string, roles []string) (interface{}, error) {
		result, err := ResolveCapabilitiesPerSite(ctx, coll, domain, roles)
		if err != nil {
			log.Error().Err(err).Str("domain", domain).
				Msg(semLogContext + " - resolve error")
			return nil, err
		}
		return result, nil
	}
}

// GetSiteAwareFromCache cerca le capabilities per-site nella cache.
// In caso di miss chiama il resolver (NewSiteAwareCacheResolver), salva e ritorna.
// La chiave usa "__persite__" per isolarsi dalla cache site-agnostica.
func GetSiteAwareFromCache(resolver CacheResolver, ctx context.Context, domain string, roles []string) (map[string]map[string]map[string][]CapDef, bool) {
	const semLogContext = "opem-store/acl/get-site-aware-from-cache"

	k := perSiteCacheKey(domain, roles)
	item, ok := aclCache.Get(k)
	if !ok {
		log.Debug().Str("k", k).Msg(semLogContext + " - cache miss")
		var err error
		item, err = resolver.Retrieve(ctx, domain, "__persite__", roles)
		if err != nil {
			log.Error().Err(err).Str("k", k).Msg(semLogContext + " - resolver error")
			return nil, false
		}
		aclCache.Set(k, item, cache.DefaultExpiration)
	}

	if item == nil {
		return nil, false
	}
	return item.(map[string]map[string]map[string][]CapDef), true
}

// InvalidateCapabilitiesPerSite rimuove dalla cache la entry per-site
// per domain+roles. Da chiamare insieme a InvalidateCapabilities quando
// cambiano i dati ACL.
func InvalidateCapabilitiesPerSite(domain string, roles []string) {
	if aclCache != nil {
		aclCache.Delete(perSiteCacheKey(domain, roles))
	}
}

