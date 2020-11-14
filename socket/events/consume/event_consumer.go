package consume

import (
	"api.namegame.com/socket"
	socketio "github.com/googollee/go-socket.io"
)

type EventConsumer interface {
	Bind(GameContext socket.GameContext) (string, func(socketio.Conn, string))
}
