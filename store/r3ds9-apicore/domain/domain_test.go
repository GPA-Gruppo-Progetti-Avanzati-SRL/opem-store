package domain_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/r3ds9-apicore/domain"
	"testing"
	"time"
)

const (
	targetDomain = "cvf"
)

func TestCache(t *testing.T) {

	r := domain.NewCacheResolver(collection)

	domain.NewCache(5*time.Second, 10*time.Minute)
	d, ok := domain.GetFromCache(r, targetDomain)
	t.Log(d, ok)

	t.Log("briefly sleeping....")
	time.Sleep(2 * time.Second)
	d, ok = domain.GetFromCache(r, targetDomain)
	t.Log(d, ok)

	t.Log("sleeping longer....")
	time.Sleep(10 * time.Second)

	d, ok = domain.GetFromCache(r, targetDomain)
	t.Log(d, ok)
}
