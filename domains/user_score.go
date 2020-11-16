package domains

type UserScore struct {
	User  User `json:"user"`
	Score int  `json:"score"`
}
