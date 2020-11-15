package consume

import (
	"api.namegame.com/socket/data"
	socketio "github.com/googollee/go-socket.io"
)

type RoomJoin struct{}

func (r *RoomJoin) Bind(GameContext data.GameContext) (string, func(socketio.Conn, string)) {
	return "ROOM_JOIN", func(s socketio.Conn, msg string) {

	}
}
