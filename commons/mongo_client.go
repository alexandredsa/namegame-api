package commons

import (
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

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	if err != nil {
		panic(err)
	}

	m.Client = client
	m.DB = client.Database(databaseName)
	return m.DB
}
