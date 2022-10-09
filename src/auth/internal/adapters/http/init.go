package http

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/popoffvg/async-arch/auth/internal/adapters/http/gen"
	"github.com/popoffvg/async-arch/auth/internal/adapters/http/handlers"
	"github.com/popoffvg/async-arch/auth/internal/ports"
	"github.com/popoffvg/async-arch/common/pkg/logger"
)

type (
	Adapter struct {
		s *http.Server
		l net.Listener
	}
)

func New(cfg *Config, l logger.Logger, userService ports.UserService, authService ports.Auth) (*Adapter, error) {
	li, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return nil, fmt.Errorf("start server failed: %w", err)
	}

	g := gin.Default()

	type handlerConainer struct {
		handlers.Auth
		handlers.Users
	}
	hc := handlerConainer{
		Auth: handlers.Auth{
			Usecases: authService,
		},
		Users: handlers.Users{
			Usecases: userService,
		},
	}
	gen.RegisterHandlersWithOptions(g, &hc, gen.GinServerOptions{
		Middlewares: []gen.MiddlewareFunc{
			gen.MiddlewareFunc(ginzap.Ginzap(l.Desugar(), time.RFC3339, true)),
			gen.MiddlewareFunc(ginzap.RecoveryWithZap(l.Desugar(), true)),
		},
	})
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: g,
	}
	return &Adapter{
		l: li,
		s: &server,
	}, nil
}

func (a *Adapter) Start() error {
	return a.s.Serve(a.l)
}

func (a *Adapter) Stop(ctx context.Context) error {
	return a.s.Shutdown(ctx)
}
