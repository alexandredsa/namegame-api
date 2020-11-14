package produce

import (
	"api.namegame.com/socket"
	socketio "github.com/googollee/go-socket.io"
)

type Scoreboard struct {
	GameContext socket.GameContext
}

func (s *Scoreboard) Emit(c *socketio.Conn) {

}
