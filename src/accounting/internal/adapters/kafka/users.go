package kafka

import (
	"context"
	"github.com/popoffvg/async-arch/common/pkg/events"
)

func (a *Adapter) ReadUserCUD(ctx context.Context) {
	for {
		defer func() {
			if r := recover(); r != nil {
				a.log.Errorf("panic: %v", r)
			}
		}()

		select {
		case <-ctx.Done():
			return
		default:

		}

		_, event, err := readMessage[events.UserCUD](ctx, a.rUsers)
		if err != nil {
			a.log.Errorf("read message from %s failed %s", topicUsers, err.Error())
			continue
		}

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
