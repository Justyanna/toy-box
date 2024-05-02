package routes

import (
	"app-server/internal/handlers"
	"app-server/internal/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ServerRouter() *gin.Engine {
	secret_key := os.Getenv("SECRET_KEY")
	if secret_key == "" {
		secret_key = "test"
	}
	secret_byte := []byte(secret_key)

	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	r := gin.Default()
	// Add Zap-suggar logger to via middleware
	r.Use(func(c *gin.Context) {
		c.Set("logger", sugar)
		c.Next()
	})
	r.GET("/health_check", handlers.HealthCheckHandler)
	r.GET("/mock_token", func(c *gin.Context) {
		handlers.MockTokenHandler(c, secret_byte)
	})

	// Add JWT token middleware
	auth_endpoints := r.Group("/api", middleware.AuthMiddleware(secret_byte))
	auth_endpoints.GET("/token_check", handlers.TokenCheckHandler)
	return r
}
