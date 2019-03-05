package controllers

import (
	"net/http"

	"../database"
	models "../models"

	"github.com/gin-gonic/gin"
)

type SampleController struct{}

var db = database.GetDatabase()

// List shows a list of samples
// @Summary
// @Tags Sample
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {array} models.SampleModel
// @Router /samples [get]
func (h SampleController) List(c *gin.Context) {
	var models []models.SampleModel
	err := db.Find(&models).Error
	if err != nil {
		c.AbortWithStatus(500)
	}
	c.JSON(http.StatusOK, models)
}

// Read returns a single sample by id
// @Summary Read returns a single sample by id
// @Tags Sample
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {object} models.SampleModel
// @Router /samples/:id [get]
func (h SampleController) Read(c *gin.Context) {
	var model models.SampleModel
	err := db.First(&model, c.Param("id")).Error
	if err != nil {
		c.AbortWithStatus(500)
	}
	c.JSON(http.StatusOK, model)
}

// Create stores a sample model
// @Summary Create stores a sample model
// @Tags Sample
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {object} models.SampleModel
// @Router /samples [post]
func (h SampleController) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// Update updates a sample model
// @Summary Update updates a sample model
// @Tags Sample
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Produce  json
// @Success 200 {object} models.SampleModel
// @Router /samples [put]
func (h SampleController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// Delete removes a sample by id
// @Summary Delete removes a sample by id
// @Tags Sample
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200
// @Router /samples [delete]
func (h SampleController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
