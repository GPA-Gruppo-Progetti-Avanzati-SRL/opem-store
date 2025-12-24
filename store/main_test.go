package store_test

import (
	"os"
	"testing"
	"time"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/system/sequence"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-mongo-common/mongolks"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const ()

var cfg = mongolks.Config{
	Name:          "default",
	Host:          os.Getenv("MONGODB_HOST_ENV"),
	DbName:        os.Getenv("MONGODB_DB_ENV"),
	User:          os.Getenv("MONGODB_USR_ENV"),
	Pwd:           os.Getenv("MONGODB_PWD_ENV"),
	AuthMechanism: "",
	AuthSource:    os.Getenv("MONGODB_AUTHSRC_ENV"),
	Pool: mongolks.PoolConfig{
		MinConn:               1,
		MaxConn:               20,
		ConnectTimeout:        1 * time.Second,
		MaxConnectionIdleTime: 30 * time.Second,
	},
	//BulkWriteOrdered: true,
	WriteConcern: "majority",
	//WriteTimeout: "120s",
	Collections: []mongolks.CollectionCfg{
		{
			Id:   sequence.CollectionId,
			Name: "opem_system",
		},
	},
	SecurityProtocol: os.Getenv("MONGODB_SECURITY_PROTOCOL"),
	TLS:              mongolks.TLSConfig{SkipVerify: os.Getenv("MONGODB_SKPVER_ENV") == "true"},
}

func TestMain(m *testing.M) {
	_, err := mongolks.Initialize([]mongolks.Config{cfg})
	if err != nil {
		panic(err)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	exitVal := m.Run()
	os.Exit(exitVal)
}
