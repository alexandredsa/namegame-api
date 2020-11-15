package repositories

import (
	"api.namegame.com/domains"
)

type RoomStateRepository struct {
	RoomStates map[string]domains.RoomState
}

func (r RoomStateRepository) FindByRoomCode(roomCode string) domains.RoomState {
	roomState := r.RoomStates[roomCode]
	return roomState
}
