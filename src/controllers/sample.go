package controllers

import (
	"net/http"

	"../database"
	"../models"

	"github.com/gin-gonic/gin"
)

var db = database.GetDatabase()

type SampleController struct{}

// List shows a list of samples
// @Summary List shows a list of samples
// @Tags Sample
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {array} models.SampleModel
// @Router /samples/ [get]
func (h SampleController) List(c *gin.Context) {
	var models []models.SampleModel
	err := db.Find(&models).Error
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
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
// @Param id path int true "Sample ID"
// @Router /samples/{id} [get]
func (h SampleController) Read(c *gin.Context) {
	var model models.SampleModel
	id := c.Params.ByName("id")
	err := db.First(&model, id).Error
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
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
// @Param model body models.SampleModel true "Sample Model"
// @Router /samples/ [post]
func (h SampleController) Create(c *gin.Context) {
	var model models.SampleModel
	c.BindJSON(&model)
	err := db.Create(&model).Error
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, model)
}

// Update updates a sample model
// @Summary Update updates a sample model
// @Tags Sample
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {object} models.SampleModel
// @Param model body models.SampleModel true "Sample Model"
// @Router /samples/ [put]
func (h SampleController) Update(c *gin.Context) {
	var model models.SampleModel
	c.BindJSON(&model)
	err := db.Update(&model).Error
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, model)
}

// Delete removes a sample by id
// @Summary Delete removes a sample by id
// @Tags Sample
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200
// @Param id path int true "Sample ID"
// @Router /samples/{id} [delete]
func (h SampleController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var model models.SampleModel
	err := db.Where("id = ?", id).Delete(&model).Error
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, gin.H{})
}
