package middleware

import (
	"go-project-akhir-ezra/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Request does not contain an access token"})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateJWT(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)

		c.Next()
	}
}
