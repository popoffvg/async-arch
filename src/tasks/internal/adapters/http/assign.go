package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *Adapter) Assign(c *gin.Context) {
	err := a.assigner.PlanAssign(c)
	if err != nil {
		c.String(http.StatusInternalServerError, "unexpected error")
		return
	}

	c.Status(http.StatusOK)
}
