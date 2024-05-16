package mock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *BaseMockHandler) HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
