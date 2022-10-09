package auth

import (
	"github.com/popoffvg/async-arch/auth/internal/ports"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewConfig,
		New,
		ProvidePort,
		ProvidePasswordManager,
	),
)

func ProvidePort(a *Auth) ports.Auth {
	return a
}

func ProvidePasswordManager(a *Auth) ports.PasswordManager {
	return a
}
