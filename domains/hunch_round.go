package domains

type UserHunch struct {
	User  User
	Hunch int
}

type HunchRound struct {
	Name        int
	Answer      int
	UserHunches []UserHunch
}
