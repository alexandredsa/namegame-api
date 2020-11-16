package emitters

import (
	"firebase.google.com/go/messaging"
)

type RoomState struct {
	FirebaseClient *messaging.Client
}

func (r RoomState) Run(RoomCode string) (err error) {
	return err
}
