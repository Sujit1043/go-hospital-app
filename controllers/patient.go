package controllers

import (
	"go-hospital-app/config"
	"go-hospital-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePatient(c *gin.Context) {
	role := c.MustGet("role").(string)
	if role != "receptionist" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&patient)
	c.JSON(http.StatusOK, patient)
}

func GetPatients(c *gin.Context) {
	var patients []models.Patient
	config.DB.Find(&patients)
	c.JSON(http.StatusOK, patients)
}

func UpdatePatient(c *gin.Context) {
	var patient models.Patient
	id := c.Param("id")
	if err := config.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.BindJSON(&patient)
	config.DB.Save(&patient)
	c.JSON(http.StatusOK, patient)
}

func DeletePatient(c *gin.Context) {
	role := c.MustGet("role").(string)
	if role != "receptionist" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only receptionist can delete"})
		return
	}
	id := c.Param("id")
	config.DB.Delete(&models.Patient{}, id)
	c.JSON(http.StatusOK, gin.H{"msg": "Deleted"})
}
