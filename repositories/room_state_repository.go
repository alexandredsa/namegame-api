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

func (r RoomStateRepository) Add(roomState domains.RoomState) {
	r.RoomStates[roomState.Code] = roomState
}
