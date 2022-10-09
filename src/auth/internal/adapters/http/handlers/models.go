package handlers

type (
	User struct {
		ID     string   `json:"id"`
		Login  string   `json:"login"`
		Scopes []string `json:"scopes"`
	}
)
