package consume

import (
	"api.namegame.com/socket"
	socketio "github.com/googollee/go-socket.io"
)

type PlayerStateUpdate struct{}

func (p *PlayerStateUpdate) Bind(GameContext socket.GameContext) (string, func(socketio.Conn, string)) {
	return "PLAYER_STATE_UPDATE", func(socketio.Conn, string) {

	}
}
