package kafka

import (
	"context"
	"github.com/popoffvg/async-arch/tasks/internal/ports"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewConfig),
	fx.Provide(New),
	fx.Provide(toAssigner),
	fx.Invoke(registerHook),
)

func toAssigner(a *Adapter) ports.Assigner {
	return a
}

func registerHook(lifecicle fx.Lifecycle, a *Adapter) {
	lifecicle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) (err error) {
				startRead(a)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				for _, c := range a.cancels {
					c()
				}

				a.rAssign.Close()
				a.rUsers.Close()
				a.rUsersRoleChanged.Close()

				return nil
			},
		},
	)
}
