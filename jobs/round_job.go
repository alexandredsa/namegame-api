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
	RoomStateEmitter     emitters.RoomState
}

func (r RoundJob) FinishRound(roomState domains.RoomState) {
	sleepTime := roomState.Round.EndsAt - int32(time.Now().Unix())
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)

	bestHunch := r.HunchRoundRepository.CalculateBestHunch(roomState.Code)
	roomState.Round.Winner = bestHunch
	r.ScoreboardRepository.UpdateUserScorePoints(roomState.Code, bestHunch.User.FCMToken, 1)
	r.RoomStateEmitter.Run(roomState.Code)
}
