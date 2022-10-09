package users

import (
	"context"

	"github.com/popoffvg/async-arch/auth/internal/core/models"
	"github.com/popoffvg/async-arch/auth/internal/ports"
)

type UserService struct {
	userStorage ports.UserStorage
	auth        ports.PasswordManager
}

func New(userStorage ports.UserStorage, pm ports.PasswordManager) *UserService {
	return &UserService{
		userStorage: userStorage,
		auth:        pm,
	}
}

func (us *UserService) SaveOrUpdate(ctx context.Context, m models.User) (_ models.User, err error) {
	m.Password, err = us.auth.Encode(m.Password)
	if err != nil {
		return m, err
	}
	return us.userStorage.CreateOrUpdate(ctx, m)
}

func (us *UserService) GetUsers(ctx context.Context, offset, limit int) ([]*models.User, error) {
	return us.userStorage.List(ctx)
}

func (us *UserService) Delete(ctx context.Context, id int) error {
	return us.userStorage.Delete(ctx, id)
}
