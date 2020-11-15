package consume

import (
	"api.namegame.com/socket/data"
	socketio "github.com/googollee/go-socket.io"
)

type PlayerStateUpdate struct{}

func (p *PlayerStateUpdate) Bind(GameContext data.GameContext) (string, func(socketio.Conn, string)) {
	return "PLAYER_STATE_UPDATE", func(s socketio.Conn, msg string) {

	}
}
