package events

const (
	CUDTypeCreate = "CREATE"
	CUDTypeUpdate = "UPDATE"
	CUDTypeDelete = "DELETE"
)

type CUDType string

type (
	UserCUD struct {
		EventType CUDType `json:"type"`
		ID        string  `json:"id"`
		Login     string  `json:"login"`
	}

	UserRoleChanged struct {
		ID    string   `json:"id"`
		Login string   `json:"login"`
		Roles []string `json:"roles"`
	}
)
