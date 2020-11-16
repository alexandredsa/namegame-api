package services

import (
	"time"

	"api.namegame.com/domains"
	"api.namegame.com/messaging/emitters"
	"api.namegame.com/repositories"
)

type RoomService struct {
	HunchRoundRepository repositories.HunchRoundRepository
	RoomStateRepository  repositories.RoomStateRepository
	ScoreboardRepository repositories.ScoreboardRepository
	RoomStateEmitter     emitters.RoomState
	ScoreboardEmitter    emitters.Scoreboard
}

func (r RoomService) GetByRoomCode(roomCode string) (domains.RoomState, domains.Scoreboard) {
	roomState := r.RoomStateRepository.FindByRoomCode(roomCode)
	scoreboard := r.ScoreboardRepository.FindByRoomCode(roomCode)

	if roomState.Round.EndsAt >= int32(time.Now().Unix()) {
		bestHunch := r.HunchRoundRepository.CalculateBestHunch(roomCode)
		roomState.Round.Winner = bestHunch
		r.ScoreboardRepository.UpdateUserScorePoints(roomCode, bestHunch.User.FCMToken, 1)
	}

	return roomState, scoreboard
}

func (r RoomService) Create(fcmToken string, username string) (domains.RoomState, domains.Scoreboard) {
	roomState := domains.RoomState{}
	roomState.Code = roomState.GenerateCode()

	userScores := make([]domains.UserScore, 0)
	userScores = append(userScores, domains.UserScore{Score: 0,
		User: domains.User{Name: username, FCMToken: fcmToken, State: "PENDING"}})
	scoreboard := domains.Scoreboard{RoomCode: roomState.Code, UserScores: userScores}

	r.RoomStateRepository.Add(roomState)
	r.ScoreboardRepository.Add(scoreboard)

	return roomState, scoreboard

}

func (r RoomService) Join(fcmToken string, username string, roomCode string) (domains.RoomState, domains.Scoreboard) {
	scoreboard := r.ScoreboardRepository.FindByRoomCode(roomCode)
	scoreboard.UserScores = append(scoreboard.UserScores,
		domains.UserScore{
			Score: 0,
			User:  domains.User{Name: username, FCMToken: fcmToken, State: "PENDING"}})

	roomState := r.RoomStateRepository.FindByRoomCode(roomCode)
	r.ScoreboardRepository.Add(scoreboard)
	r.RoomStateEmitter.Run(roomCode)
	return roomState, scoreboard
}

func (r RoomService) HunchCreate(fcmToken string, roomCode string, hunch int) (err error) {
	scoreboard := r.ScoreboardRepository.FindByRoomCode(roomCode)
	var user domains.User

	for _, userScore := range scoreboard.UserScores {
		if userScore.User.FCMToken == fcmToken {
			user = userScore.User
			return
		}
	}

	r.HunchRoundRepository.AddUserHunch(user, roomCode, hunch)
	return err
}

func (r RoomService) UpdatePlayerState(fcmToken string, state string, roomCode string) (err error) {
	r.ScoreboardRepository.UpdateUserScoreState(roomCode, fcmToken, state)
	r.RoomStateEmitter.Run(roomCode)
	r.ScoreboardEmitter.Run(roomCode)
	return err
}
