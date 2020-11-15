package consume

import (
	"api.namegame.com/socket/data"
	socketio "github.com/googollee/go-socket.io"
)

type HunchCreate struct{}

func (h *HunchCreate) Bind(GameContext data.GameContext) (string, func(socketio.Conn, string)) {
	return "HUNCH_CREATE", func(socketio.Conn, string) {

	}
}
