package r3ds9_apigtw

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/r3ds9-apicore/domain"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/r3ds9-apicore/session"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/r3ds9-apicore/site"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/r3ds9-apicore/user"
	"time"
)

const (
	ExpiryDuration       = 5 * time.Minute
	PurgeExpiredDuration = 10 * time.Minute
)

func InitStore(expiryDuration time.Duration, purgeExpiryDuration time.Duration) error {

	if expiryDuration == 0 {
		expiryDuration = ExpiryDuration
	}

	if purgeExpiryDuration == 0 {
		purgeExpiryDuration = PurgeExpiredDuration
	}

	domain.NewCache(expiryDuration, purgeExpiryDuration)
	site.NewCache(expiryDuration, purgeExpiryDuration)
	user.NewCache(expiryDuration, purgeExpiryDuration)
	session.NewCache(expiryDuration, purgeExpiryDuration)
	return nil
}
