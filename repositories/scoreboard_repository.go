package repositories

import (
	"api.namegame.com/domains"
)

type ScoreboardRepository struct {
	Scoreboards map[string]domains.Scoreboard
}

func (s ScoreboardRepository) UpdateUserScoreState(roomCode string, fcmToken string, state string) (isRoomReady bool) {
	scoreboard := s.Scoreboards[roomCode]
	isRoomReady = true
	for i := 0; i < len(scoreboard.UserScores); i++ {
		if scoreboard.UserScores[i].User.FCMToken == fcmToken {
			scoreboard.UserScores[i].User.State = state

			if !isRoomReady {
				return false
			}
		}

		if scoreboard.UserScores[i].User.State == "PENDING" {
			isRoomReady = false
		}
	}

	return isRoomReady
}

func (s ScoreboardRepository) ResetUserScoreState(roomCode string) {
	scoreboard := s.Scoreboards[roomCode]
	for i := 0; i < len(scoreboard.UserScores); i++ {
		scoreboard.UserScores[i].User.State = "PENDING"
	}
}

func (s ScoreboardRepository) UpdateUserScorePoints(roomCode string, fcmToken string, pointsToAdd int) {
	scoreboard := s.Scoreboards[roomCode]
	for i := 0; i < len(scoreboard.UserScores); i++ {
		if scoreboard.UserScores[i].User.FCMToken == fcmToken {
			currentScore := scoreboard.UserScores[i].Score
			scoreboard.UserScores[i].Score = currentScore + pointsToAdd
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
