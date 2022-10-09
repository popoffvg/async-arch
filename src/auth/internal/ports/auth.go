package ports

import (
	"context"

	"github.com/popoffvg/async-arch/auth/internal/core/models"
)

type (
	Auth interface {
		Login(ctx context.Context, login, password string) (string, error)
		Verify(ctx context.Context, token string) (models.VerifyResponse, error)
	}

	PasswordManager interface {
		Encode(password []byte) ([]byte, error)
	}
)
