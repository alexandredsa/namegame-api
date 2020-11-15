package domains

import (
	"math/rand"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct {
	Name   string
	Answer int
}

type Round struct {
	Current  int8
	Max      int8
	Question Question
	EndsAt   primitive.Timestamp
}

type Winner struct {
	User  UserState
	Hunch int
}

type RoomState struct {
	Code  string
	Round Round
}

func (rs RoomState) UpdateCode() {
	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 6)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	rs.Code = string(b)
}
