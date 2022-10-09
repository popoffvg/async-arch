package http

import (
	"fmt"
	"github.com/caarlos0/env"
)

type Config struct {
	Port int `env:"HTTP_PORT" envDefault:"3000"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("parse http adapter config failed: %w", err)
	}

	return &cfg, nil
}
