package domains

type User struct {
	Name     string `json:"name"`
	FCMToken string `json:"-"`
	State    string `json:"state"`
}
