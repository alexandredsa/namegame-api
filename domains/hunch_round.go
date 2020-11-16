package domains

type UserHunch struct {
	User  User `json:"user"`
	Hunch int  `json:"hunch"`
}

type HunchRound struct {
	Name        int         `json:"name"`
	Answer      int         `json:"answer"`
	UserHunches []UserHunch `json:"hunches"`
}
