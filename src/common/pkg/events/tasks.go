package events

type TaskCUD struct {
	Event

	ID       string `json:"id"`
	Title    string `json:"title"`
	Status   string `json:"status"`
	Assigned string `json:"assigned"`
}
