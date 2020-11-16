package repositories

import (
	"api.namegame.com/domains"
)

type ScoreboardRepository struct {
	Scoreboards map[string]domains.Scoreboard
}

func (s ScoreboardRepository) UpdateUserScoreState(roomCode string, fcmToken string, state string) {
	scoreboard := s.Scoreboards[roomCode]
	for _, userScore := range scoreboard.UserScores {
		if userScore.User.FCMToken == fcmToken {
			userScore.User.State = state
			return
		}
	}
}
func (s ScoreboardRepository) UpdateUserScorePoints(roomCode string, fcmToken string, pointsToAdd int) {
	scoreboard := s.Scoreboards[roomCode]
	for _, userScore := range scoreboard.UserScores {
		if userScore.User.FCMToken == fcmToken {
			userScore.Score = userScore.Score + pointsToAdd
			return
		}
	}
}

func (s ScoreboardRepository) FindByRoomCode(roomCode string) domains.Scoreboard {
	scoreboard := s.Scoreboards[roomCode]
	return scoreboard
}

func (s ScoreboardRepository) Add(scoreboard domains.Scoreboard) {
	s.Scoreboards[scoreboard.RoomCode] = scoreboard
}
