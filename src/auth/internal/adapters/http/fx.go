package http

import (
	"context"
	"time"

	"github.com/caarlos0/env"
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

type Config struct {
	Port int `env:"HTTP_PORT" envDefault:"3000"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

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
