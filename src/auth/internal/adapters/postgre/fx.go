package postgre

import (
	"github.com/popoffvg/async-arch/auth/internal/ports"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewConfig,
		New,
		ProvidePort,
	),
)

func ProvidePort(a *Adapter) ports.UserStorage {
	return a
}
