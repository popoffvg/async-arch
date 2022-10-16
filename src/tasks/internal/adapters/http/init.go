package http

import (
	"context"
	"errors"
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/popoffvg/async-arch/common/pkg/logger"
	"github.com/popoffvg/async-arch/tasks/ent/ogent"
	"github.com/popoffvg/async-arch/tasks/internal/core/tasks"
	"github.com/popoffvg/async-arch/tasks/internal/ports"
	"go.uber.org/zap"
	"net"
	"net/http"
	"time"

	"github.com/popoffvg/async-arch/tasks/ent"
)

type Adapter struct {
	s *http.Server
	l net.Listener

	tasks    *tasks.Service
	assigner ports.Assigner
}
type handler struct {
	*ogent.OgentHandler
}

type CtxKeys string

const (
	isRequestProcessed CtxKeys = "is_request_processed"
)

func New(client *ent.Client, cfg *Config, log logger.Logger, assigner ports.Assigner) (*Adapter, error) {
	handler, err := ogent.NewServer(handler{
		OgentHandler: ogent.NewOgentHandler(client),
	})
	if err != nil {
		return nil, fmt.Errorf("creating http adapter failed: %w", err)
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return nil, fmt.Errorf("server creating failed: %w", err)
	}

	g := gin.Default()
	g.Use(ginzap.Ginzap(log.Desugar().WithOptions(zap.AddCallerSkip(1)), time.RFC3339, true))
	g.Use(ginzap.RecoveryWithZap(log.Desugar().WithOptions(zap.AddCallerSkip(1)), true))

	assignHandler := &assignHandler{
		assigner: assigner,
	}
	g.POST("/tasks/assign", assignHandler.Assign)
	g.Any("/tasks", func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	})

	s := http.Server{
		Handler: g,
	}

	return &Adapter{
		s:        &s,
		l:        l,
		assigner: assigner,
	}, nil
}

func (a *Adapter) Start() (err error) {
	go func() {
		err = a.s.Serve(a.l)
	}()

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}

func (a *Adapter) Stop(ctx context.Context) error {
	return a.s.Shutdown(ctx)
}
