package probes

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Config struct {
	Port int `env:"PROBES_PORT" envDefault:"8081"`
}

func (c *Config) Load() (*Config, error) {
	if err := env.Parse(c); err != nil {
		return nil, fmt.Errorf("read env for probes failed: %w", err)
	}

	return c, nil
}
