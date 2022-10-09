package http

import (
	"context"
	"go.uber.org/fx"
	"time"
)

var Module = fx.Options(
	fx.Provide(
		NewConfig,
		New,
	),
	fx.Invoke(
		registerHook,
	),
)

func registerHook(lifecicle fx.Lifecycle, s *Adapter) {
	lifecicle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) (err error) {
				go func() {
					err = s.Start()
				}()
				time.Sleep(3 * time.Second) // TODO: normal fx shutdown
				return err
			},
			OnStop: func(ctx context.Context) error {
				return s.Stop(ctx)
			},
		},
	)
}
