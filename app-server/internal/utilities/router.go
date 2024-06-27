package utilities

import (
	"app-server/internal/auth"
	"app-server/internal/mock"
	"app-server/internal/room"
	"app-server/internal/ws"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func ServerRouter(database *sql.DB, clientManager *ws.ClientManager) *gin.Engine {
	r := gin.Default()
	r.RedirectTrailingSlash = true
	secret_byte := getSecretKey()
	addLogger(r)

	// Mock endpoints
	r.GET("/health_check", mock.HealthCheckHandler)
	r.GET("/mock_token", func(c *gin.Context) {
		mock.MockTokenHandler(c, secret_byte)
	})

	// Endpoints with JWT middleware
	auth_endpoints := r.Group("/api", auth.AuthMiddleware(secret_byte))
	auth_endpoints.GET("/token_check", mock.TokenCheckHandler)

	room.RegisterRoutes(r, clientManager)

	return r
}
