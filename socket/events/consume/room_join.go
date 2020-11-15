package consume

import (
	"api.namegame.com/socket/data"
	"api.namegame.com/socket/events"
	socketio "github.com/googollee/go-socket.io"
)

type RoomJoin struct{}

func (r RoomJoin) Bind(GameContext data.GameContext) (string, func(socketio.Conn, string)) {
	return events.ROOM_JOIN_EVENT, func(s socketio.Conn, msg string) {

	}
}
