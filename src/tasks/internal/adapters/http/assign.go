package http

import (
	"github.com/gin-gonic/gin"
	"github.com/popoffvg/async-arch/tasks/internal/ports"
	"net/http"
)

type assignHandler struct {
	assigner ports.Assigner
}

func (a *assignHandler) Assign(c *gin.Context) {
	err := a.assigner.PlanAssign(c)
	if err != nil {
		c.String(http.StatusInternalServerError, "unexpected error")
		return
	}

	c.Status(http.StatusOK)
}
