package dtos

import "api.namegame.com/domains"

type GameData struct {
	Room       domains.RoomState  `json:"room"`
	Scoreboard domains.Scoreboard `json:"scoreboard"`
}
