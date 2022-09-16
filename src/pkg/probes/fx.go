package probes

import (
	"context"
	"time"

	"go.uber.org/fx"
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

func NewConfig() (*Config, error) {
	cfg, err := new(Config).Load()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func registerHook(lifecycle fx.Lifecycle, p *Probes) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				startTimeout := 3 * time.Second
				errCh := make(chan error)
				go func() {
					err := p.Start()
					errCh <- err
				}()

				select {
				case err :=<-errCh:
					return err
				case <-time.After(startTimeout):
				}

				return nil
			},
			OnStop: func(ctx context.Context) error {
				return p.Stop(ctx)
			},
		},
	)
}
