package ports

import (
	"context"

	"github.com/popoffvg/async-arch/auth/internal/core/models"
)

type UserStorage interface {
	GetUser(ctx context.Context, login string) (models.User, error)
	CreateOrUpdate(ctx context.Context, u models.User) (models.User, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*models.User, error)
}
