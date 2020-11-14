package consume

import (
	"api.namegame.com/socket"
	socketio "github.com/googollee/go-socket.io"
)

type RoomJoin struct{}

func (r *RoomJoin) Bind(GameContext socket.GameContext) (string, func(socketio.Conn, string)) {
	return "ROOM_JOIN", func(socketio.Conn, string) {

	}
}
