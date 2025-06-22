package routes

import (
	"go-hospital-app/controllers"
	"go-hospital-app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/login", controllers.Login)

	auth := router.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/patients", controllers.CreatePatient)
		auth.GET("/patients", controllers.GetPatients)
		auth.PUT("/patients/:id", controllers.UpdatePatient)
		auth.DELETE("/patients/:id", controllers.DeletePatient)
	}
}
