package consume

import (
	"api.namegame.com/socket/data"
	"api.namegame.com/socket/events"
	socketio "github.com/googollee/go-socket.io"
)

type PlayerStateUpdate struct{}

func (p PlayerStateUpdate) Bind(GameContext data.GameContext) (string, func(socketio.Conn, string)) {
	return events.PLAYER_STATE_UPDATE_EVENT, func(s socketio.Conn, msg string) {

	}
}
