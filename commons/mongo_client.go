package commons

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func (m MongoClient) GetDatabase(mongoURI string, databaseName string) *mongo.Database {

	if m.DB != nil {
		return m.DB
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	if err != nil {
		panic(err)
	}

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	m.Client = client
	m.DB = client.Database(databaseName)
	return m.DB
}
