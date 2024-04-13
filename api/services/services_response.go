package services

import (
	"time"

	"github.com/google/uuid"
)

// @Description	ServicesIn represents the input structure for creating a new service.
type ServicesIn struct {
	Name        string    `json:"name" binding:"required,min=3,max=30"`         // name is the  name of the service. It is required and should be between 3 and 30 characters.
	Description string    `json:"description" binding:"required,min=3,max=35"`  // description is the description of the service. It is required and should be between 3 and 35 characters.
	CompanyID   uuid.UUID `json:"companyID" binding:"required"`                 // CompanyID is the unique identifier for the company associated with the service. It is required.
	KeywordName string    `json:"keywordName" binding:"required,min=2,max=255"` // KeywordName is the name of the keyword's service. It is required and should be between 2 and 255 characters.
} //@name ServicesIn

// @Description	ServicesPagination represents the paginated list of Services.
type ServicesPagination struct {
	Items      []ServicesTable `json:"items"`      // Items is a slice containing individual service details.
	Page       uint            `json:"page"`       // Page is the current page number in the pagination.
	Limit      uint            `json:"limit"`      // Limit is the maximum number of items per page in the pagination.
	TotalCount uint            `json:"totalCount"` // TotalCount is the total number of services in the entire list.
} //@name ServicesPagination

// @Description	ServicesTable represents a single service entry in a table.
type ServicesTable struct {
	ID          uuid.UUID `json:"id"`                                          // ID is the unique identifier for the user.
	Name        string    `json:"name" binding:"required,min=3,max=30"`        // name is the  name of the service. It is required and should be between 3 and 30 characters.
	Description string    `json:"description" binding:"required,min=3,max=35"` // description is the description of the service. It is required and should be between 3 and 35 characters.
	CreatedAt   time.Time `json:"createdAt"`                                   // CreatedAt is the timestamp indicating when the service entry was created.
} //@name UsersTable

// @Description	ServicesList represents a simplified version of the service for listing purposes.
type ServicesList struct {
	ID   uuid.UUID `json:"id"`   // ID is the unique identifier for the service.
	Name string    `json:"name"` // Name is the full name of the service.
} //@name ServicesList

// @Description	ServicesCount represents the count of services.
type ServicesCount struct {
	Count uint `json:"count"` // Count is the number of services.
} //@name ServicesCount

// @Description	ServicesDetails represents detailed information about a specific service.
type ServicesDetails struct {
	ID          uuid.UUID `json:"id"`                                          // ID is the unique identifier for the user.
	Name        string    `json:"name" binding:"required,min=3,max=30"`        // name is the  name of the service. It is required and should be between 3 and 30 characters.
	Description string    `json:"description" binding:"required,min=3,max=35"` // description is the description of the service. It is required and should be between 3 and 35 characters.
	CreatedAt   time.Time `json:"createdAt"`                                   // CreatedAt is the timestamp indicating when the service entry was created.
} //@name ServicesDetails
