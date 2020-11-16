package services

import (
	"api.namegame.com/domains"
	"api.namegame.com/messaging/emitters"
	"api.namegame.com/repositories"
)

type RoomService struct {
	RoomStateRepository  repositories.RoomStateRepository
	ScoreboardRepository repositories.ScoreboardRepository
	RoomStateEmitter     emitters.RoomState
	ScoreboardEmitter    emitters.Scoreboard
}

func (r RoomService) Create(fcmToken string, username string) (domains.RoomState, domains.Scoreboard) {
	roomState := domains.RoomState{}
	roomState.UpdateCode()

	userScores := make([]domains.UserScore, 0)
	userScores = append(userScores, domains.UserScore{Score: 0,
		User: domains.User{Name: username}})
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
			User:  domains.User{Name: username}})

	roomState := r.RoomStateRepository.FindByRoomCode(roomCode)
	r.ScoreboardRepository.Add(scoreboard)
	r.RoomStateEmitter.Run(roomCode)
	return roomState, scoreboard

}

func (r RoomService) UpdatePlayerState(fcmToken string, state string, roomCode string) (err error) {
	r.ScoreboardRepository.UpdateUserScoreState(roomCode, fcmToken, state)
	r.RoomStateEmitter.Run(roomCode)
	r.ScoreboardEmitter.Run(roomCode)
	return err
}
