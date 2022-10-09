package ports

import (
	"context"
	"github.com/popoffvg/async-arch/common/pkg/events"

	"github.com/popoffvg/async-arch/auth/internal/core/models"
)

type UserService interface {
	SaveOrUpdate(ctx context.Context, m models.User) (models.User, error)
	GetUsers(ctx context.Context, offset, limit int) ([]*models.User, error) // TODO: offset
	Delete(ctx context.Context, id string) error
}

type EventRegistrator interface {
	ProduceCUDEvent(ctx context.Context, e events.UserCUD) error
	ProduceRolesChangedEvent(ctx context.Context, e events.UserRoleChanged) error
}
