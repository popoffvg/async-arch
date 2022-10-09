package postgre

import (
	"context"
	"errors"

	"github.com/popoffvg/async-arch/auth/ent"
	"github.com/popoffvg/async-arch/auth/ent/user"
	"github.com/popoffvg/async-arch/auth/internal/core/models"
)

func (a *Adapter) GetUser(ctx context.Context, login string) (u models.User, err error) {
	var eUser *ent.User
	eUser, err = a.users.Query().
		Where(user.Login(login)).
		Only(ctx)
	if err != nil {
		if errors.Is(err, &ent.NotFoundError{}) {
			return models.User(*eUser), models.ErrNotFound
		}
		return u, err
	}

	return *eUser, nil
}

func (a *Adapter) CreateOrUpdate(ctx context.Context, m models.User) (u models.User, err error) {
	var (
		eUser    *ent.User
		isCreate bool
	)
	if m.ID == 0 {
		isCreate = true
	}

	if isCreate {
		eUser, err = a.users.Create().
			SetLogin(m.Login).
			SetPassword(m.Password).
			SetScopes(m.Scopes).
			Save(ctx)
	} else {
		t := ent.User(m)
		eUser, err = a.users.UpdateOne(&t).Save(ctx)
	}

	if err != nil {
		return u, err
	}
	return models.User(*eUser), nil
}

func (a *Adapter) Delete(ctx context.Context, id int) error {
	return a.users.DeleteOneID(id).Exec(ctx)
}

func (a *Adapter) List(ctx context.Context) ([]*models.User, error) {
	return a.users.Query().All(ctx)
}
