package services

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ServiceRouterInit initializes the routes related to services.
func ServiceRouterInit(router *gin.RouterGroup, db *gorm.DB) {

	// Initialize database instance for services
	baseInstance := Database{DB: db}

	// Automigrate / Update service-related table
	NewServiceRepository(db)

	// Private
	services := router.Group("/services/:companyID")
	{

		// POST endpoint to create a new service
		services.POST("", baseInstance.CreateService)

		// GET endpoint to retrieve all services for a specific company
		services.GET("", baseInstance.ReadServices)

		// GET endpoint to retrieve a list of services for a specific company
		services.GET("/list", baseInstance.ReadServicesList)

		// GET endpoint to retrieve the count of services for a specific company
		services.GET("/count", baseInstance.ReadServicesCount)

		// GET endpoint to retrieve details of a specific service
		services.GET("/:ID", baseInstance.ReadService)

		// PUT endpoint to update details of a specific service
		services.PUT("/:ID", baseInstance.UpdateService)

		// DELETE endpoint to delete a specific service
		services.DELETE("/:ID", baseInstance.DeleteService)
	}
}
