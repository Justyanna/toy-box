package mock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *BaseMockHandler) TokenCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
