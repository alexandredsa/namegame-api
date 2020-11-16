package domains

import (
	"math/rand"
)

type Round struct {
	Current  int8
	Max      int8
	Question Question
	Winner   UserHunch
	EndsAt   int32
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
