package handlers

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/popoffvg/async-arch/auth/internal/ports"
)

var (
	bearerPattern = regexp.MustCompile("Bearer (.+)")
)

type (
	Auth struct {
		Usecases ports.Auth
	}

	Headers struct {
		Auth string `header:"Authorization"`
	}
)

func (a *Auth) PostLogin(c *gin.Context) {
	login, password, ok := c.Request.BasicAuth()
	if !ok {
		c.Header("WWW-Authenticate", "basic")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := a.Usecases.Login(c, login, password)
	if err != nil {
		c.Error(err)
		return
	}

	c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	c.JSON(http.StatusOK, http.NoBody)
}

// (POST /verify)
func (a *Auth) PostVerify(c *gin.Context) {

	var h Headers
	err := c.ShouldBindHeader(&h)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	token, err := h.BearerToken()
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	data, err := a.Usecases.Verify(c, token)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *Headers) BearerToken() (string, error) {
	groups := bearerPattern.FindStringSubmatch(h.Auth)
	if len(groups) < 2 {
		return "", fmt.Errorf("token not valid")
	}

	return groups[1], nil
}
