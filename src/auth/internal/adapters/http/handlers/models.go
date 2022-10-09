package handlers

type (
	User struct {
		ID     int      `json:"id"`
		Login  string   `json:"login"`
		Scopes []string `json:"scopes"`
	}
)
