package probes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New(cfg *Config) *Probes {

	r := gin.Default()
	r.GET("/healthz", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})
	r.GET("/ready", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})
	r.GET("/startup", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: r.Handler(),
	}

	return &Probes{
		server: &server,
	}
}

type Probes struct {
	server *http.Server
}

func (p *Probes) Start() error {
	return p.server.ListenAndServe()
}

func (p *Probes) Stop(ctx context.Context) error {
	return p.server.Shutdown(ctx)
}
