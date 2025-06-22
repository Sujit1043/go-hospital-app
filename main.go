package main

import (
	"go-hospital-app/config"
	"go-hospital-app/models"
	"go-hospital-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	config.DB.AutoMigrate(&models.User{}, &models.Patient{})

	// Seed users
	config.DB.Create(&models.User{Username: "doc1", Password: "pass", Role: "doctor"})
	config.DB.Create(&models.User{Username: "recept1", Password: "pass", Role: "receptionist"})

	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")
}
