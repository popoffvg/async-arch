package postgre

import (
	"context"
	"errors"
	"github.com/popoffvg/async-arch/auth/internal/core/models"
	"github.com/popoffvg/async-arch/auth/internal/ent"
	"github.com/popoffvg/async-arch/auth/internal/ent/user"
	"github.com/popoffvg/async-arch/common/pkg/events"
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
	if m.ID == "" {
		isCreate = true
	}
	a.l.Infof("user UPDATE/CREATE %#v", m)

	if isCreate {
		eUser, err = a.users.Create().
			SetLogin(m.Login).
			SetPassword(m.Password).
			SetScopes(m.Scopes).
			Save(ctx)

		if err != nil {
			return models.User{}, err
		}
		a.onCreate(ctx, eUser)
	} else {
		t := ent.User(m)
		eUser, err = a.users.UpdateOne(&t).Save(ctx)
		if err != nil {
			return models.User{}, err
		}
		a.onUpdate(ctx, eUser)
	}

	if err != nil {
		return u, err
	}
	return *eUser, nil
}

func (a *Adapter) onCreate(ctx context.Context, user *ent.User) {
	err := a.mb.ProduceCUDEvent(ctx, events.UserCUD{
		Event: events.Event{
			EventType: events.CUDTypeCreate,
		},
		ID:    user.ID,
		Login: user.Login,
	})
	if err != nil {
		a.l.Errorf("produce CREATE event failed: %s for id %s", err.Error(), user.ID)
	}
}

func (a *Adapter) onUpdate(ctx context.Context, user *ent.User) {
	err := a.mb.ProduceCUDEvent(ctx, events.UserCUD{
		Event: events.Event{
			EventType: events.CUDTypeCreate,
		},
		ID:    user.ID,
		Login: user.Login,
	})
	if err != nil {
		a.l.Errorf("produce UPDATE event failed: %s for id %s", err.Error(), user.ID)
	}
}

func (a *Adapter) onDelete(ctx context.Context, id string) {
	err := a.mb.ProduceCUDEvent(ctx, events.UserCUD{
		Event: events.Event{
			EventType: events.CUDTypeCreate,
		},
		ID: id,
	})
	if err != nil {
		a.l.Errorf("produce DELETE event failed: %s for id %s", err.Error(), id)
	}
}

func (a *Adapter) Delete(ctx context.Context, id string) error {
	err := a.users.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}

	a.onDelete(nil, id)
	return nil
}

func (a *Adapter) List(ctx context.Context) ([]*models.User, error) {
	return a.users.Query().All(ctx)
}
