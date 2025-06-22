package controllers

import (
	"go-hospital-app/config"
	"go-hospital-app/models"
	"go-hospital-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	config.DB.Where("username = ? AND password = ?", input.Username, input.Password).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	token, _ := utils.GenerateToken(user.Username, user.Role)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
