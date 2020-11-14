package domains

import "go.mongodb.org/mongo-driver/bson/primitive"

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
