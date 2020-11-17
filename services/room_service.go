package services

import (
	"time"

	"api.namegame.com/domains"
	"api.namegame.com/jobs"
	"api.namegame.com/messaging/emitters"
	"api.namegame.com/repositories"
)

type RoomService struct {
	HunchRoundRepository     repositories.HunchRoundRepository
	NameStatisticsRepository repositories.NameStatisticsRepository
	RoomStateRepository      repositories.RoomStateRepository
	ScoreboardRepository     repositories.ScoreboardRepository
	RoomStateEmitter         emitters.RoomState
	ScoreboardEmitter        emitters.Scoreboard
	RoundJob                 jobs.RoundJob
}

func (r RoomService) GetByRoomCode(roomCode string) (domains.RoomState, domains.Scoreboard) {
	roomState := r.RoomStateRepository.FindByRoomCode(roomCode)
	scoreboard := r.ScoreboardRepository.FindByRoomCode(roomCode)
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
			break
		}
	}

	r.HunchRoundRepository.AddUserHunch(user, roomCode, hunch)
	return err
}

func (r RoomService) CreateNextRound(roomCode string) domains.RoomState {
	roomState := r.RoomStateRepository.FindByRoomCode(roomCode)
	round := domains.Round{}
	if roomState.Round.Current == 0 {
		round.Current = 1
		round.Max = 10
	} else {
		round.Current = roomState.Round.Current + 1
	}

	nameStatistics := r.NameStatisticsRepository.Shuffle()
	round.Question = domains.
		Question{Name: nameStatistics.Name,
		Answer: int(nameStatistics.Total)}

	endsAt := time.Now().Local().Add(time.Second * time.Duration(30))
	round.EndsAt = int32(endsAt.Unix())
	roomState.Round = round
	r.RoomStateRepository.Add(roomState)
	return roomState
}

func (r RoomService) UpdatePlayerState(fcmToken string, state string, roomCode string) (err error) {
	isRoomReady := r.ScoreboardRepository.UpdateUserScoreState(roomCode, fcmToken, state)
	r.ScoreboardEmitter.Run(roomCode)

	if isRoomReady {
		roomState := r.CreateNextRound(roomCode)
		go r.RoundJob.FinishRound(roomState)
	}

	r.RoomStateEmitter.Run(roomCode)

	return err
}
