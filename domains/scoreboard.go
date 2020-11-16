package domains

type Scoreboard struct {
	UserScores []UserScore `json:"user_scores"`
	RoomCode   string      `json:"room_code"`
}
