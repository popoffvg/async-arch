package postgre

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/popoffvg/async-arch/accounting/internal/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/caarlos0/env"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type (
	Config struct {
		URI      string `env:"PG_URI" envDefault:"postgresql://postgres:password@0.0.0.0:5432/accounting?sslmode=disable"`
		Password string `env:"PG_PASSWORD" envDefault:"password"`
	}
)

func New(cfg *Config) (*ent.Client, error) {
	db, err := sql.Open("pgx", cfg.URI)
	if err != nil {
		return nil, fmt.Errorf("db create failed: %w", err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	cEnt := ent.NewClient(ent.Driver(drv))

	if err := cEnt.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	return cEnt, nil
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
