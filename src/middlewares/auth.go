package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"../config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates that you have a valid JWT token
func AuthMiddleware(cfg config.Configuration) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		// Ensure we have a token in the header
		if len(authHeader) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Parse and validate the token, check for Bearer and remove.
		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwt.Parse(tokenString, keyLookupFn)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		} else if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func keyLookupFn(token *jwt.Token) (interface{}, error) {
	// Check for expected signing method.
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	// TODO: Load from config
	return []byte("A14E45A7-D02B-4ADA-94BC-66DCBFD3181E"), nil
}
