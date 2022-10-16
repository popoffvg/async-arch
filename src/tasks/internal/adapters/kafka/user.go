package kafka

import (
	"context"
	"errors"
	"github.com/popoffvg/async-arch/common/pkg/events"
	"github.com/popoffvg/async-arch/tasks/ent"
)

const roleUser = "USER"

func (a *Adapter) ReadUserCUD(ctx context.Context) {
	a.log.Info("Start read USER topic")
	for {
		select {
		case <-ctx.Done():
			return
		default:

		}

		event, err := readMessage[events.UserCUD](ctx, a.rUsers)
		if err != nil {
			a.log.Errorf("read message from %s failed %s", topicUsers, err.Error())
			continue
		}

		// TODO: mb normal service layer
		switch event.EventType {
		case events.CUDTypeCreate:
			_, err = a.orm.User.Create().
				SetID(event.ID).
				SetLogin(event.Login).
				Save(ctx)
		case events.CUDTypeUpdate:
			err = a.orm.User.UpdateOneID(event.ID).
				SetLogin(event.Login).
				Exec(ctx)
		case events.CUDTypeDelete:
			err = a.orm.User.DeleteOneID(event.ID).Exec(ctx)
		}
		if err != nil {
			a.log.Errorf("process CUD user event %#v failed: %s", event, err.Error())
		}
	}
}

func (a *Adapter) readRolesChanged(ctx context.Context) {
	a.log.Info("Start read ROLES topic")
	for {
		select {
		case <-ctx.Done():
			return
		default:

		}

		event, err := readMessage[events.UserRoleChanged](ctx, a.rUsers)
		if err != nil {
			a.log.Errorf("read message from %s failed %s", topicUsers, err.Error())
			continue
		}

		if len(event.Roles) > 1 || notContains(event.Roles, roleUser) {
			err = a.orm.User.DeleteOneID(event.ID).Exec(ctx)
			if err != nil {
				if errors.Is(err, &ent.NotFoundError{}) {
					continue
				}
				a.log.Errorf("process message %#v failed: %s", event, err.Error())
				continue
			}
		}

		if len(event.Roles) == 1 && event.Roles[0] == roleUser {
			err = a.orm.User.Create().
				SetID(event.ID).
				SetLogin(event.Login).
				Exec(ctx)
			if err != nil {
				a.log.Errorf("process message %#v failed: %s", event, err.Error())
				continue
			}
		}
	}
}

func notContains(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}

	}
	return false
}
