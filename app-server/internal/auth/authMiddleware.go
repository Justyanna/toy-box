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
		authHeaderParts := strings.Split(authHeader, " ")

		if !isAuthHeaderPresent(authHeader, c) {
			return
		}

		if !isAuthHeaderValid(authHeaderParts, c) {
			return
		}

		tokenString := authHeaderParts[1]
		token, err := parseToken(tokenString, secret_byte)
		if err != nil {
			logger.Errorw("Header token parsing failed", "error", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}

		userID, err := extractClaims(logger, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}
		c.Set("userID", userID)
		c.Next()
	}
}

func isAuthHeaderPresent(authHeader string, c *gin.Context) bool {
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
		return false
	}
	return true
}

func isAuthHeaderValid(authHeaderParts []string, c *gin.Context) bool {
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
		return false
	}
	return true
}

func parseToken(tokenString string, secret_byte []byte) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secret_byte, nil
	})
}

func extractClaims(logger *zap.SugaredLogger, token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		logger.Errorw("Claims failed", "error", errors.New("claims assertion failed"))
		return "", errors.New("claims failed")
	}
	userID, ok := claims["userID"].(string)
	if !ok || userID == "" {
		logger.Errorw("Claims userID not found", "error", errors.New("userID not found"))
		return "", errors.New("userID not found")
	}
	return userID, nil
}
