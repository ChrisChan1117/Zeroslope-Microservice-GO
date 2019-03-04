package middlewares

import (
	"strings"

	"../utilities"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates that you have a valid JWT token
func AuthMiddleware(cfg utilities.Configuration) gin.HandlerFunc {
	return func(c *gin.Context) {

		reqKey := c.Request.Header.Get("X-Auth-Key")
		reqSecret := c.Request.Header.Get("X-Auth-Secret")

		var key string
		var secret string
		if key = cfg.JWT.Key; len(strings.TrimSpace(key)) == 0 {
			c.AbortWithStatus(500)
		}
		if secret = cfg.JWT.Secret; len(strings.TrimSpace(secret)) == 0 {
			c.AbortWithStatus(401)
		}
		if key != reqKey || secret != reqSecret {
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}
