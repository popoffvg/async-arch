package ports

import (
	"context"

	"github.com/popoffvg/async-arch/auth/internal/core/models"
)

type UserService interface {
	SaveOrUpdate(ctx context.Context, m models.User) (models.User, error)
	GetUsers(ctx context.Context, offset, limit int) ([]*models.User, error) // TODO: offset
	Delete(ctx context.Context, id int) error
}
