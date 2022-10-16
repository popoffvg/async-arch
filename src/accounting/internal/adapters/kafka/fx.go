package kafka

import (
	"context"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewConfig),
	fx.Provide(New),
	fx.Invoke(registerHook),
)

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

				return nil
			},
		},
	)
}
