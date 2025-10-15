package site_test

import (
	"context"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

const (
	instance       = "mongodb://localhost:27017"
	dbName         = "r3ds9"
	collectionName = "site"
)

var collection *mongo.Collection

func TestMain(m *testing.M) {
	client, err := mongo.Connect(options.Client().ApplyURI(instance))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}

	database := client.Database(dbName)
	collection = database.Collection(collectionName)

	exitVal := m.Run()
	os.Exit(exitVal)
}
