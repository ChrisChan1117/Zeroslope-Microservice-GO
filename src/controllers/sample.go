package controllers

import (
	"net/http"

	"github.com/cfryerdev/zeroslope-golang/src/database"
	"github.com/cfryerdev/zeroslope-golang/src/models"

	"github.com/gin-gonic/gin"
)

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
	var m []models.SampleModel
	var db = database.GetDatabase()
	if err := db.Find(&m).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, m)
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
	var m models.SampleModel
	id := c.Params.ByName("id")
	var db = database.GetDatabase()
	if err := db.First(&m, id).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, m)
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
	var m models.SampleModel
	c.BindJSON(&m)
	var db = database.GetDatabase()
	if err := db.Create(&m).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, m)
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
	var m models.SampleModel
	c.BindJSON(&m)
	var db = database.GetDatabase()
	if err := db.Update(&m).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, m)
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
	var m models.SampleModel
	var db = database.GetDatabase()
	if err := db.Where("id = ?", id).Delete(&m).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, gin.H{})
}
