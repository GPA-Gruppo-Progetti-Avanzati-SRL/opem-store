package model

import (
	"crypto/rand"
	"encoding/hex"
)

// newOID genera un identificatore opaco compatibile con il formato MongoDB ObjectID
// (12 byte random → 24 char hex). Usato dai driver statici/YAML quando _id è assente.
func newOID() string {
	var b [12]byte
	if _, err := rand.Read(b[:]); err != nil {
		panic("model: failed to generate OID: " + err.Error())
	}
	return hex.EncodeToString(b[:])
}
