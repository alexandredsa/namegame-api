package consume

import (
	"api.namegame.com/socket/data"
	socketio "github.com/googollee/go-socket.io"
)

type RoomCreate struct{}

func (r *RoomCreate) Bind(GameContext data.GameContext) (string, func(socketio.Conn, string)) {
	return "ROOM_CREATE", func(socketio.Conn, string) {

	}
}
