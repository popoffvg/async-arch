package postgre

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/popoffvg/async-arch/auth/internal/ports"
	"github.com/popoffvg/async-arch/common/pkg/logger"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/caarlos0/env"
	"github.com/popoffvg/async-arch/auth/internal/ent"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type (
	Config struct {
		URI      string `env:"PG_URI" envDefault:"postgresql://postgres:password@0.0.0.0:5432/postgres?sslmode=disable"`
		Password string `env:"PG_PASSWORD" envDefault:"password"`
	}

	Adapter struct {
		users ent.UserClient
		mb    ports.EventRegistrator
		l     logger.Logger
	}
)

func New(cfg *Config, mb ports.EventRegistrator, log logger.Logger) (*Adapter, error) {
	db, err := sql.Open("pgx", cfg.URI)
	if err != nil {
		return nil, fmt.Errorf("db create failed: %w", err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	cEnt := ent.NewClient(ent.Driver(drv))
	if err := cEnt.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	return &Adapter{
		users: *cEnt.User,
		mb:    mb,
		l:     log,
	}, nil
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
