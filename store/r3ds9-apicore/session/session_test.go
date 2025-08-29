package session_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/r3ds9-apicore/session"
	"testing"
	"time"
)

const (
	targetDomain = "cvf"
)

func TestCache(t *testing.T) {

	r := session.NewCacheResolver(collection)

	session.NewCache(5*time.Second, 10*time.Minute)
	d, ok := session.GetFromCache(r, targetDomain)
	t.Log(d, ok)

	t.Log("briefly sleeping....")
	time.Sleep(2 * time.Second)
	d, ok = session.GetFromCache(r, targetDomain)
	t.Log(d, ok)

	t.Log("sleeping longer....")
	time.Sleep(10 * time.Second)

	d, ok = session.GetFromCache(r, targetDomain)
	t.Log(d, ok)
}
