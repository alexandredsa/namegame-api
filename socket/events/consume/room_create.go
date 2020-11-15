package consume

import (
	"encoding/json"

	"api.namegame.com/domains"

	"api.namegame.com/repositories"
	"api.namegame.com/socket/data"
	"api.namegame.com/socket/data/messages"
	socketio "github.com/googollee/go-socket.io"
)

type RoomCreate struct {
	RoomStateRepository  repositories.RoomStateRepository
	ScoreboardRepository repositories.ScoreboardRepository
}

func (r *RoomCreate) Bind(GameContext data.GameContext) (string, func(socketio.Conn, string)) {
	return "ROOM_CREATE", func(s socketio.Conn, msg string) {
		roomCreateMessage := messages.RoomCreate{}
		err := json.Unmarshal([]byte(msg), roomCreateMessage)

		if err != nil {
			panic(err)
		}

		roomState := domains.RoomState{}
		roomState.UpdateCode()
		r.RoomStateRepository.Add(roomState)
	}
}
