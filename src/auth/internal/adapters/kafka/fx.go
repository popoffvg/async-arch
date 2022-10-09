package kafka

import (
	"github.com/popoffvg/async-arch/auth/internal/ports"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewConfig),
	fx.Provide(New),
	fx.Provide(toEventRegistrator),
)

func toEventRegistrator(a *Adapter) ports.EventRegistrator {
	return a
}
