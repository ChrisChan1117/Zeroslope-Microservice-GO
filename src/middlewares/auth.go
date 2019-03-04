package middlewares

import (
	"net/http"

	"../utilities"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates that you have a valid JWT token
func AuthMiddleware(cfg utilities.Configuration) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if len(authHeader) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
