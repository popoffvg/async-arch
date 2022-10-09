package models

import "github.com/popoffvg/async-arch/auth/ent"

type User = ent.User

type VerifyResponse struct {
	AccessToken  string
	RefreshToken string
	User         User
}
