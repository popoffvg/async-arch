package main

import (
	"context"
	"github.com/popoffvg/async-arch/auth/internal/adapters/kafka"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/popoffvg/async-arch/auth/internal/adapters/http"
	"github.com/popoffvg/async-arch/auth/internal/adapters/postgre"
	"github.com/popoffvg/async-arch/auth/internal/core/auth"
	"github.com/popoffvg/async-arch/auth/internal/core/users"
	"github.com/popoffvg/async-arch/common/pkg/logger"
	"github.com/popoffvg/async-arch/common/pkg/probes"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		probes.Module,
		logger.Module,
		http.Module,
		auth.Module,
		users.Module,
		postgre.Module,
		kafka.Module,
	)

	startCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	<-startCtx.Done()

	stopCtx, cancelStop := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelStop()

	err := app.Stop(stopCtx)
	if err != nil {
		log.Fatalf("App stop failed: %s", err.Error())
	}
}
