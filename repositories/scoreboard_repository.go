package repositories

import (
	"api.namegame.com/domains"
)

type ScoreboardRepository struct {
	Scoreboards map[string]domains.Scoreboard
}

func (s ScoreboardRepository) FindByRoomCode(roomCode string) domains.Scoreboard {
	scoreboard := s.Scoreboards[roomCode]
	return scoreboard
}

func (s ScoreboardRepository) Add(scoreboard domains.Scoreboard) {
	s.Scoreboards[scoreboard.RoomCode] = scoreboard
}
