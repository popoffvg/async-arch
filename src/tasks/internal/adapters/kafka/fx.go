package kafka

import (
	"context"
	"github.com/popoffvg/async-arch/tasks/internal/ports"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewConfig),
	fx.Provide(New),
	fx.Provide(toAssigner),
	fx.Invoke(startAssigner),
)

func toAssigner(a *Adapter) ports.Assigner {
	return a
}

func startAssigner(a *Adapter) {
	a.ProduceAssign(context.Background())
}
