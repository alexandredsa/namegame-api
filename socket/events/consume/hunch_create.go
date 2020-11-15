package consume

import (
	"api.namegame.com/socket/data"
	"api.namegame.com/socket/events"
	socketio "github.com/googollee/go-socket.io"
)

type HunchCreate struct{}

func (h *HunchCreate) Bind(GameContext data.GameContext) (string, func(socketio.Conn, string)) {
	return events.HUNCH_CREATE_EVENT, func(socketio.Conn, string) {

	}
}
