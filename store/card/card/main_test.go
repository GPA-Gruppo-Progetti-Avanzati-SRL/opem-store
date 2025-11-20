package card_test

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-mongo-common/mongolks"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	CardsCollectionId   = "card"
	CardsCollectionName = "opem_cards"

	HostEnvVarName             = "MONGODB_HOST_ENV"
	DbEnvVarName               = "MONGODB_DB_ENV"
	UserEnvVarName             = "MONGODB_USR_ENV"
	PwdEnvVarName              = "MONGODB_PWD_ENV"
	AuthSourceEnvVarName       = "MONGODB_AUTHSRC_ENV"
	SkpVerifEnvVarName         = "MONGODB_SKPVER_ENV"
	SecurityProtocolEnvVarName = "MONGODB_SECURITY_PROTOCOL"
)

var cfg = mongolks.Config{
	Name:          "default",
	Host:          os.Getenv(HostEnvVarName),
	DbName:        os.Getenv(DbEnvVarName),
	User:          os.Getenv(UserEnvVarName),
	Pwd:           os.Getenv(PwdEnvVarName),
	AuthMechanism: "",
	AuthSource:    os.Getenv(AuthSourceEnvVarName),
	Pool: mongolks.PoolConfig{
		MinConn: 1,
		MaxConn: 20,
		//MaxWaitQueueSize:      1000,
		ConnectTimeout:        3 * time.Second,
		MaxConnectionIdleTime: 30000,
		//MaxConnectionLifeTime: 6000000,
	},
	//BulkWriteOrdered: true,
	WriteConcern: "majority",
	//WriteTimeout:     "120s",
	Collections: []mongolks.CollectionCfg{
		{
			Id:   CardsCollectionId,
			Name: CardsCollectionName,
		},
	},
	SecurityProtocol: os.Getenv(SecurityProtocolEnvVarName),
	TLS:              mongolks.TLSConfig{SkipVerify: false},
}

func TestMain(m *testing.M) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	skp, err := strconv.ParseBool(os.Getenv(SkpVerifEnvVarName))
	if err != nil {
		panic(err)
	}

	cfg.TLS.SkipVerify = skp

	_, err = mongolks.Initialize([]mongolks.Config{cfg})
	if err != nil {
		panic(err)
	}

	exitVal := m.Run()
	os.Exit(exitVal)
}
