package domains

import (
	"math/rand"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Round struct {
	Current  int8
	Max      int8
	Question Question
	Winner   UserHunch
	EndsAt   primitive.Timestamp
}

type RoomState struct {
	Code  string
	Round Round
}

func (rs RoomState) GenerateCode() string {
	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 6)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
