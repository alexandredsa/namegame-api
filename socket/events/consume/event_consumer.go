package consume

import (
	"api.namegame.com/socket/data"
	socketio "github.com/googollee/go-socket.io"
)

type EventConsumer interface {
	Bind(GameContext data.GameContext) (string, func(socketio.Conn, string))
}
