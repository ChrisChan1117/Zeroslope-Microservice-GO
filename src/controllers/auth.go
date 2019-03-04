package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

// Login generates a JWT token
func (h AuthController) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"token": "TOKENHERE"})
}
