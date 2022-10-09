package auth

import (
	"bytes"
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/caarlos0/env"
	"github.com/golang-jwt/jwt/v4"
	"github.com/popoffvg/async-arch/auth/internal/core/models"
	"github.com/popoffvg/async-arch/auth/internal/ports"
	"github.com/xdg-go/pbkdf2"
)

type (
	Config struct {
		AccessTokenExpired time.Duration `env:"AUTH_JWT_EXPIRED" envDefault:"24h"`
		Secret             string        `env:"AUTH_SECRET" envDefault:"test"`
		Salt               string        `env:"AUTH_SALT" envDefault:"test"`
	}

	Auth struct {
		cfg         *Config
		userStorage ports.UserStorage
	}
)

func New(cfg *Config, storage ports.UserStorage) *Auth {
	return &Auth{
		cfg:         cfg,
		userStorage: storage,
	}
}

func (a *Auth) Login(ctx context.Context, login string, password string) (string, error) {
	user, err := a.userStorage.GetUser(ctx, login)
	if err != nil {
		return "", err
	}
	cryptPass, _ := a.Encode([]byte(password))
	if !bytes.Equal(cryptPass, user.Password) {
		return "", models.ErrForbidden
	}

	access, err := a.generateToken(login, a.cfg.AccessTokenExpired)
	if err != nil {
		return "", err
	}

	return access, nil
}

func (a *Auth) Verify(ctx context.Context, token string) (r models.VerifyResponse, err error) {

	var login string
	login, err = a.verifyToken(token)
	if err != nil && !errors.Is(err, models.ErrTokenExpired) {
		return r, err
	}

	if err != nil {
		return r, err
	}

	r.User, err = a.userStorage.GetUser(ctx, login)
	if err != nil {
		return r, err
	}

	r.AccessToken, err = a.generateToken(login, a.cfg.AccessTokenExpired)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (a *Auth) generateToken(login string, expiredIn time.Duration) (string, error) {
	mySigningKey := a.cfg.Secret

	now := time.Now()
	// Create the Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(now.Add(expiredIn)),
		Issuer:    login,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		return "", fmt.Errorf("generate token failed: %w", err)
	}
	return ss, nil
}

func (a *Auth) verifyToken(tokenString string) (string, error) {
	var claims jwt.RegisteredClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.cfg.Secret), nil
	})
	if !token.Valid {
		return "", fmt.Errorf("parse token unexpected error: %w", err)
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return "", models.ErrTokenInvalid
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		// Token is either expired or not active yet
		return "", models.ErrTokenExpired
	}

	return claims.Issuer, nil
}

func (a *Auth) Encode(password []byte) ([]byte, error) {
	dk := pbkdf2.Key(password, []byte(a.cfg.Salt), 4096, 32, sha1.New)
	return dk, nil
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
