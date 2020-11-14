package domains

type UserState string

const (
	PENDING UserState = "PENDING"
	READY   UserState = "READY"
)

type User struct {
	Name  string
	state UserState
}
