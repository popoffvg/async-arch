package events

type (
	UserCUD struct {
		Event

		ID    string `json:"id"`
		Login string `json:"login"`
	}

	UserRoleChanged struct {
		Event

		ID    string   `json:"id"`
		Login string   `json:"login"`
		Roles []string `json:"roles"`
	}
)
