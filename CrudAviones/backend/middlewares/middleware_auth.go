package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MiddlewareAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("x-is-authentication")

		if token != "xur-2225-vcx-8900-aie" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
			c.Abort()
			return
		}
		c.Next()
	}
}
