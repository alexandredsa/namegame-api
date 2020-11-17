package repositories

import (
	"context"

	"api.namegame.com/domains"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NameStatisticsRepository struct {
	DB *mongo.Database
}

func (n *NameStatisticsRepository) Shuffle() domains.NameStatistics {
	collection := n.DB.Collection("name_statistics")
	ctx := context.Background()
	sampleStage := bson.D{{"$sample", bson.D{{"size", 1}}}}
	nameDataCursor, err := collection.Aggregate(ctx, mongo.Pipeline{sampleStage})

	if err != nil {
		panic(err)
	}

	var nameStatistics []domains.NameStatistics
	if err = nameDataCursor.All(ctx, &nameStatistics); err != nil {
		panic(err)
	}

	nameStatistic := nameStatistics[0]
	return nameStatistic
}
