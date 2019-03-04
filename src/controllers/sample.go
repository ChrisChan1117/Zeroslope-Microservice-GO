package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SampleController struct{}

// List shows a list of samples
func (h SampleController) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// Read returns a single sample by id
func (h SampleController) Read(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// Create stores a sample model
func (h SampleController) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// Update updates a sample model
func (h SampleController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// Delete removes a sample by id
func (h SampleController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
