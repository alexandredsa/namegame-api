package repositories

import (
	"math"

	"api.namegame.com/domains"
)

type HunchRoundRepository struct {
	HunchRounds map[string]domains.HunchRound
}

func (h HunchRoundRepository) AddUserHunch(user domains.User, roomCode string, hunch int) {
	userHunch := domains.UserHunch{User: user, Hunch: hunch}
	hunchRound := h.HunchRounds[roomCode]
	hunchRound.UserHunches = append(hunchRound.UserHunches, userHunch)
	h.HunchRounds[roomCode] = hunchRound
}

func (h HunchRoundRepository) CalculateBestHunch(roomCode string) domains.UserHunch {
	var bestUserHunch domains.UserHunch
	hunchRound := h.HunchRounds[roomCode]
	for _, userHunch := range hunchRound.UserHunches {
		if bestUserHunch.User.FCMToken == "" {
			bestUserHunch = userHunch
			continue
		}

		thisDiff := math.Abs(float64(hunchRound.Answer - userHunch.Hunch))
		bestCurrentDiff := math.Abs(float64(hunchRound.Answer - bestUserHunch.Hunch))

		if thisDiff < bestCurrentDiff {
			bestUserHunch = userHunch
		}
	}

	return bestUserHunch
}
