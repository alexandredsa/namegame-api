package consume

import (
	"api.namegame.com/socket"
	socketio "github.com/googollee/go-socket.io"
)

type HunchCreate struct{}

func (h *HunchCreate) Bind(GameContext socket.GameContext) (string, func(socketio.Conn, string)) {
	return "HUNCH_CREATE", func(socketio.Conn, string) {

	}
}
