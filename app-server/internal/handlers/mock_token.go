package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

func MockTokenHandler(c *gin.Context, secret_byte []byte) {
	logger := c.MustGet("logger").(*zap.SugaredLogger)

	expirationTime := time.Now().Add(12 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": "bar",
		"exp":    expirationTime.Unix(),
	})
	tokenString, err := token.SignedString(secret_byte)
	if err != nil {
		logger.Errorw("An error occurred", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred"})
	}
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
