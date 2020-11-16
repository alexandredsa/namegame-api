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
	collection := n.DB.Collection("names")
	ctx := context.Background()
	sampleStage := bson.D{{"$sample", bson.D{{"size", 1}}}}
	nameDataCursor, err := collection.Aggregate(ctx, mongo.Pipeline{sampleStage})

	if err != nil {
		panic(err)
	}

	var nameStatistics []domains.NameStatistics
	nameDataCursor.All(ctx, &nameStatistics)
	nameStatistic := nameStatistics[0]
	nameStatistic.Total = nameStatistic.Until1930 +
		nameStatistic.Until1940 + nameStatistic.Until1950 +
		nameStatistic.Until1960 + nameStatistic.Until1970 +
		nameStatistic.Until1980 + nameStatistic.Until1990 +
		nameStatistic.Until2000 + nameStatistic.Until2010
	return nameStatistic
}
