package kafka

import (
	"context"
	"github.com/popoffvg/async-arch/accounting/internal/ent/task"
	"github.com/popoffvg/async-arch/common/pkg/events"
)

type EventCtxKey int

const (
	EventIDKey EventCtxKey = iota
)

func (a *Adapter) ReadTaskCUD(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:

		}

		ctx, event, err := readMessage[events.TaskCUD](ctx, a.rTasks)
		if err != nil {
			a.log.Errorf("read message from %s failed %s", topicUsers, err.Error())
			continue
		}

		switch event.EventType {
		case events.CUDTypeCreate:
			_, err = a.orm.Task.Create().
				SetID(event.ID).
				SetTitle(event.Title).
				SetAssignedID(event.Assigned).
				SetStatus(task.Status(event.Status)).
				Save(ctx)
		case events.CUDTypeUpdate:
			err = a.orm.Task.UpdateOneID(event.ID).
				SetTitle(event.Title).
				SetAssignedID(event.Assigned).
				SetStatus(task.Status(event.Status)).
				Exec(ctx)
		case events.CUDTypeDelete:
			err = a.orm.Task.DeleteOneID(event.ID).Exec(ctx)
		}
		if err != nil {
			a.log.Errorf("process CUD user event %#v failed: %s", event, err.Error())
		}
	}
}

func (a *Adapter) ReadTaskAssign(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:

		}

		ctx, event, err := readMessage[events.TaskCUD](ctx, a.rAssign)
		if err != nil {
			a.log.Errorf("read message from %s failed %s", topicUsers, err.Error())
			continue
		}

		exist, err := a.orm.Task.Query().Where(task.ID(event.ID)).Exist(ctx)
		if err != nil {
			a.log.Errorf("process event %#v failed: %s", event, err.Error())
			continue
		}
		if !exist {
			err = a.orm.Task.Create().SetID(event.ID).SetAssignedID(event.Assigned).Exec(ctx)
		} else {
			err = a.orm.Task.UpdateOneID(event.ID).SetAssignedID(event.Assigned).Exec(ctx)
		}

		if err != nil {
			a.log.Errorf("process event %#v failed: %s", event, err.Error())
		}
	}
}
