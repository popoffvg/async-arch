package main

import (
	"context"
	"github.com/popoffvg/async-arch/accounting/internal/adapters/kafka"
	"github.com/popoffvg/async-arch/accounting/internal/adapters/postgre"
	"github.com/popoffvg/async-arch/common/pkg/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/popoffvg/async-arch/accounting/internal/ent/runtime"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		logger.Module,
		kafka.Module,
		postgre.Module,
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
