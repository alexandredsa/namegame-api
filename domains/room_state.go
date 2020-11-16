package domains

import (
	"math/rand"
)

type Round struct {
	Current  int8      `json:"current"`
	Max      int8      `json:"max"`
	Question Question  `json:"question"`
	Winner   UserHunch `json:"winner"`
	EndsAt   int32     `json:"ends_at"`
}

type RoomState struct {
	Code  string `json:"code"`
	Round Round  `json:"round"`
}

func (rs RoomState) GenerateCode() string {
	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 6)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
