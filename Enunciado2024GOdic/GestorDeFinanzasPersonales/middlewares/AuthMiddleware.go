package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("x-caller-auth")

		if authHeader != "true" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "No autorizado",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
