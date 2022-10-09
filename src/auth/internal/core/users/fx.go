package users

import (
	"github.com/popoffvg/async-arch/auth/internal/ports"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		New,
		ProvidePort,
	),
)

func ProvidePort(s *UserService) ports.UserService {
	return s
}
