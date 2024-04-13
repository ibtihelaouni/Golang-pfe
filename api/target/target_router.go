package target

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TargetRouterInit initializes the routes related to targets.
func TargetRouterInit(router *gin.RouterGroup, db *gorm.DB) {

	// Initialize database instance for targets
	baseInstance := Database{DB: db}

	// Automigrate / Update target-related table
	NewTargetRepository(db)

	// Private
	targets := router.Group("/targets/:companyID")
	{

		// POST endpoint to create a new target
		targets.POST("", baseInstance.CreateTarget)

		// GET endpoint to retrieve all targets for a specific company
		targets.GET("", baseInstance.ReadTargets)

		// GET endpoint to retrieve a list of targets for a specific company
		targets.GET("/list", baseInstance.ReadTargetList)

		// GET endpoint to retrieve the count of targets for a specific company
		targets.GET("/count", baseInstance.ReadTargetCount)

		// GET endpoint to retrieve details of a specific target
		targets.GET("/:ID", baseInstance.ReadTarget)

		// PUT endpoint to update details of a specific target
		targets.PUT("/:ID", baseInstance.UpdateTarget)

		// DELETE endpoint to delete a specific targets
		targets.DELETE("/:ID", baseInstance.DeleteTarget)
	}
}
