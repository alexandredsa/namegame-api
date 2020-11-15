package consume

import (
	"encoding/json"

	"api.namegame.com/domains"
	"api.namegame.com/socket/events"

	"api.namegame.com/repositories"
	"api.namegame.com/socket/data"
	"api.namegame.com/socket/data/messages"
	socketio "github.com/googollee/go-socket.io"
)

type RoomCreate struct {
	RoomStateRepository  repositories.RoomStateRepository
	ScoreboardRepository repositories.ScoreboardRepository
}

func (r RoomCreate) Bind(GameContext data.GameContext) (string, func(socketio.Conn, string)) {
	return events.ROOM_CREATE_EVENT, func(s socketio.Conn, msg string) {
		roomCreateMessage := messages.RoomCreate{}
		err := json.Unmarshal([]byte(msg), roomCreateMessage)

		if err != nil {
			panic(err)
		}

		roomState := domains.RoomState{}
		roomState.UpdateCode()

		userScores := make([]domains.UserScore, 0)
		userScores = append(userScores, domains.UserScore{Score: 0,
			User: domains.User{Name: roomCreateMessage.Username}})
		scoreboard := domains.Scoreboard{RoomCode: roomState.Code, UserScores: userScores}

		r.RoomStateRepository.Add(roomState)
		r.ScoreboardRepository.Add(scoreboard)

		roomStateMessageBytes, err := json.Marshal(roomState)

		if err != nil {
			panic(err)
		}

		s.Emit(events.ROOM_STATE_EVENT, string(roomStateMessageBytes))
	}
}
