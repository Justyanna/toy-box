package routes

import (
	"app-server/internal/auth"
	"app-server/internal/mock"
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ServerRouter(database *sql.DB) *gin.Engine {
	secret_key := os.Getenv("SECRET_KEY")
	secret_byte := []byte(secret_key)

	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	mock_handler := mock.NewBaseMockHandler(database)

	r := gin.Default()
	// Add Zap-suggar logger
	r.Use(func(c *gin.Context) {
		c.Set("logger", sugar)
		c.Next()
	})

	// Mock endpoints
	r.GET("/health_check", mock_handler.HealthCheckHandler)
	r.GET("/mock_token", func(c *gin.Context) {
		mock_handler.MockTokenHandler(c, secret_byte)
	})

	// Endpoints with JWT middleware
	auth_endpoints := r.Group("/api", auth.AuthMiddleware(secret_byte))
	auth_endpoints.GET("/token_check", mock_handler.TokenCheckHandler)
	return r
}
