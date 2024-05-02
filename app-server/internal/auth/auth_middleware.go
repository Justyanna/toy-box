package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

func AuthMiddleware(secret_byte []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := c.MustGet("logger").(*zap.SugaredLogger)
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			return
		}

		tokenString := authHeaderParts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return secret_byte, nil
		})
		if err != nil {
			logger.Errorw("Header token parsing failed", "error", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userID, ok := claims["userID"].(string)
			if !ok || userID == "" {
				logger.Errorw("Claims userID not found", "error", errors.New("userID not found"))
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			}
			c.Set("userID", userID)
			c.Next()
		} else {
			logger.Errorw("Claims failed", "error", errors.New("userID not found"))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}
	}
}
