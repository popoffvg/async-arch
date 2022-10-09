package kafka

import (
	"context"
	"github.com/popoffvg/async-arch/common/pkg/events"
)

func (a *Adapter) ProduceCUDEvent(ctx context.Context, e events.UserCUD) error {
	return writeMessage(ctx, a.wUsers, e)
}

func (a *Adapter) ProduceRolesChangedEvent(ctx context.Context, e events.UserRoleChanged) error {
	return writeMessage(ctx, a.wUsersRoleChanged, e)
}
