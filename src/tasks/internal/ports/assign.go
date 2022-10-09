package ports

import (
	"context"
)

type (
	Assigner interface {
		PlanAssign(ctx context.Context) error
	}
)
