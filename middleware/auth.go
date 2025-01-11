package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	secretKey := "secret-123"
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token missing"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
		if err != nil || !isTokenValid(token)  {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Next()
	}
}
func isTokenValid(token *jwt.Token) bool{
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expiration, ok := claims["exp"].(float64)
		if !ok {
			return false
		}

		expirationTime := time.Unix(int64(expiration), 0)
		currentTime := time.Now()

		if currentTime.Before(expirationTime) {
			return true
		} else {
			return false
		}
	}
	return false
}
