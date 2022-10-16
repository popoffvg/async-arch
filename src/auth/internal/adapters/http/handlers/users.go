package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/popoffvg/async-arch/auth/internal/adapters/http/gen"
	"github.com/popoffvg/async-arch/auth/internal/core/models"
	"github.com/popoffvg/async-arch/auth/internal/ports"
	"net/http"
)

type Users struct {
	Usecases ports.UserService
}

// (GET /users)
func (u *Users) GetUsers(c *gin.Context) {
	var (
		err error
	)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	type Params struct {
		Offset int
		Limit  int
	}
	var p Params
	err = c.BindQuery(&p)
	if err != nil {
		c.Error(err)
		return
	}
	users, err := u.Usecases.GetUsers(c, p.Offset, p.Limit)
	if err != nil {
		c.Error(err)
		return
	}

	payload := make([]*User, 0, len(users))
	for _, v := range users {
		payload = append(payload, convertToProjection(*v))
	}
	c.JSON(http.StatusOK, payload)
}

// (PATCH /users)
// TODO: make method for create user
func (u *Users) PatchUsers(c *gin.Context) {
	var (
		p   gen.User
		err error
		m   models.User
	)
	err = c.Bind(&p)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	m.Login = p.Login
	m.Password = []byte(p.Password)
	m.Scopes = "USER"
	m, err = u.Usecases.SaveOrUpdate(c, m)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, convertToProjection(m))
}

// (POST /users)
func (u *Users) PostUsers(c *gin.Context) {
	u.PatchUsers(c)
}

func convertToProjection(m models.User) *User {
	return &User{
		ID:     m.ID,
		Login:  m.Login,
		Scopes: []string{m.Scopes.String()},
	}
}
