package emitters

import (
	"context"
	"encoding/json"
	"fmt"

	"api.namegame.com/repositories"
	"firebase.google.com/go/messaging"
)

type RoomState struct {
	FirebaseClient       *messaging.Client
	RoomStateRepository  repositories.RoomStateRepository
	ScoreboardRepository repositories.ScoreboardRepository
}

func (r RoomState) Run(roomCode string) (err error) {
	ctx := context.Background()
	roomState := r.RoomStateRepository.FindByRoomCode(roomCode)
	roomStateJSON, err := json.Marshal(roomState)

	scoreboard := r.ScoreboardRepository.FindByRoomCode(roomCode)

	if err != nil {
		panic(err)
	}

	for _, userScore := range scoreboard.UserScores {
		fmt.Println(fmt.Sprintf("Sending %v to token %v", string(roomStateJSON), userScore.User.FCMToken))
		go r.FirebaseClient.Send(ctx, &messaging.Message{
			Data: map[string]string{
				"json_data":    string(roomStateJSON),
				"message_type": "ROOM_STATE",
			},
			Token: userScore.User.FCMToken,
		})
	}

	return err
}
