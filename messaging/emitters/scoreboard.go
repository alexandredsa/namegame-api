package emitters

import (
	"context"
	"encoding/json"
	"fmt"

	"api.namegame.com/repositories"
	"firebase.google.com/go/messaging"
)

type Scoreboard struct {
	FirebaseClient       *messaging.Client
	ScoreboardRepository repositories.ScoreboardRepository
}

func (s Scoreboard) Run(roomCode string) error {
	ctx := context.Background()
	scoreboard := s.ScoreboardRepository.FindByRoomCode(roomCode)
	scoreboardJSON, err := json.Marshal(scoreboard)

	if err != nil {
		panic(err)
	}

	for _, userScore := range scoreboard.UserScores {
		fmt.Println(fmt.Sprintf("Sending %v to token %v", string(scoreboardJSON), userScore.User.FCMToken))
		go s.FirebaseClient.Send(ctx, &messaging.Message{
			Data: map[string]string{
				"json_data":    string(scoreboardJSON),
				"message_type": "SCOREBOARD",
			},
			Token: userScore.User.FCMToken,
		})
	}

	return err
}
