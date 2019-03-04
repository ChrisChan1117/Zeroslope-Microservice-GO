package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

// Status will return 200 if the service is working
// @Summary Server Health
// @Description Returns a message if the service is working correctly
// @Tags Health
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"The service is working."
// @Router /health [get]
func (h HealthController) Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "The service is working."})
}
