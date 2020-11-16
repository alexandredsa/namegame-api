package dtos

import "api.namegame.com/domains"

type GameData struct {
	Room       domains.RoomState
	Scoreboard domains.Scoreboard
}
