package messages

type RoomJoin struct {
	Username string `json:"username"`
	RoomCode string `json:"room_code"`
}
