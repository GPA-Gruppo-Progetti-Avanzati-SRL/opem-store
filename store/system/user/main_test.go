package user_test

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"testing"
	"time"
)

const (
	instance       = "mongodb://localhost:27017"
	dbName         = "r3ds9"
	collectionName = "user"
)

var collection *mongo.Collection

func TestMain(m *testing.M) {
	client, err := mongo.NewClient(options.Client().ApplyURI(instance))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	database := client.Database(dbName)
	collection = database.Collection(collectionName)

	exitVal := m.Run()
	os.Exit(exitVal)
}
