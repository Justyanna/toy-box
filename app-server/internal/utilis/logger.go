package utilis

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func addLogger(r *gin.Engine) {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()

	// Add Zap-suggar logger
	r.Use(func(c *gin.Context) {
		c.Set("logger", logger)
		c.Next()
	})
}
