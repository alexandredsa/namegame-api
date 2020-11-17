package jobs

import (
	"time"

	"api.namegame.com/domains"
	"api.namegame.com/messaging/emitters"
	"api.namegame.com/repositories"
)

type RoundJob struct {
	HunchRoundRepository repositories.HunchRoundRepository
	ScoreboardRepository repositories.ScoreboardRepository
	RoomStateRepository  repositories.RoomStateRepository
	RoomStateEmitter     emitters.RoomState
}

func (r RoundJob) FinishRound(roomState domains.RoomState) {
	// sleepTime := roomState.Round.EndsAt - int32(time.Now().Unix())
	// time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	time.Sleep(30 * time.Second)

	bestHunch := r.HunchRoundRepository.CalculateBestHunch(roomState.Code)
	roomState.Round.Winner = bestHunch
	r.RoomStateRepository.Add(roomState)
	r.ScoreboardRepository.UpdateUserScorePoints(roomState.Code, bestHunch.User.FCMToken, 1)
	r.ScoreboardRepository.ResetUserScoreState(roomState.Code)
	r.RoomStateEmitter.Run(roomState.Code)
}
