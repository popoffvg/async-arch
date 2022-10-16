package ports

import "context"

type (
	Accounting interface {
		Assign(ctx context.Context, taskID string) error
		DoneTask(ctx context.Context, taskID string) error
	}
)
