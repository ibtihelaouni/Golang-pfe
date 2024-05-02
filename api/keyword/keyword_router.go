package keyword

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CompanyRouterInit initializes the routes related to companies.
func KeywordRouterInit(router *gin.RouterGroup, db *gorm.DB) {

	// Initialize database instance
	baseInstance := Database{DB: db}

	// Automigrate / Update table
	NewKeywordRepository(db)

	// Private
	keywords := router.Group("/keywords/:serviceID")
	{

		// POST endpoint to create a new keyword
		keywords.POST("", baseInstance.CreateKeyword)

		// GET endpoint to retrieve all keywords for a specific keyword
		keywords.GET("", baseInstance.ReadKeywords)

		// GET endpoint to retrieve details of a specific keyword
		keywords.GET("/:ID", baseInstance.ReadKeyword)

		// PUT endpoint to update details of a specific keyword
		keywords.PUT("/:ID", baseInstance.UpdateKeyword)

		// DELETE endpoint to delete a specific keyword
		keywords.DELETE("/:ID", baseInstance.DeleteKeyword)
	}
}
