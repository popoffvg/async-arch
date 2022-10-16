package models

import "github.com/popoffvg/async-arch/auth/internal/ent"

type User = ent.User

type VerifyResponse struct {
	AccessToken  string
	RefreshToken string
	User         User
}
