package target

import (
	"time"

	"github.com/google/uuid"
)

// @Description	TargetIn represents the input structure for creating a new target.
type TargetIn struct {
	Name      string    `json:"name" binding:"required,min=3,max=30"` // name is the  name of the service. It is required and should be between 3 and 30 characters.
	CompanyID uuid.UUID `json:"companyID" binding:"required"`         // CompanyID is the unique identifier for the company associated with the service. It is required.

} //@name TargetIn

// @Description	ServicesPagination represents the paginated list of Services.
type TargetPagination struct {
	Items      []TargetTable `json:"items"`      // Items is a slice containing individual service details.
	Page       uint          `json:"page"`       // Page is the current page number in the pagination.
	Limit      uint          `json:"limit"`      // Limit is the maximum number of items per page in the pagination.
	TotalCount uint          `json:"totalCount"` // TotalCount is the total number of services in the entire list.
} //@name TargetPagination

// @Description	TargetTable represents a single target entry in a table.
type TargetTable struct {
	ID          uuid.UUID `json:"id"`                                          // ID is the unique identifier for the user.
	Name        string    `json:"name" binding:"required,min=3,max=30"`        // name is the  name of the service. It is required and should be between 3 and 30 characters.
	Description string    `json:"description" binding:"required,min=3,max=35"` // description is the description of the service. It is required and should be between 3 and 35 characters.
	CreatedAt   time.Time `json:"createdAt"`                                   // CreatedAt is the timestamp indicating when the service entry was created.
} //@name TargetTable

// @Description	TargetList represents a simplified version of the target for listing purposes.
type TargetList struct {
	ID   uuid.UUID `json:"id"`   // ID is the unique identifier for the service.
	Name string    `json:"name"` // Name is the full name of the service.
} //@name TargetList

// @Description	TargetCount represents the count of target.
type TargetCount struct {
	Count uint `json:"count"` // Count is the number of target.
} //@name TargetCount

// @Description	TargetDetails represents detailed information about a specific target.
type TargetDetails struct {
	ID          uuid.UUID `json:"id"`                                          // ID is the unique identifier for the user.
	Name        string    `json:"name" binding:"required,min=3,max=30"`        // name is the  name of the service. It is required and should be between 3 and 30 characters.
	Description string    `json:"description" binding:"required,min=3,max=35"` // description is the description of the service. It is required and should be between 3 and 35 characters.
	CreatedAt   time.Time `json:"createdAt"`                                   // CreatedAt is the timestamp indicating when the service entry was created.
} //@name TargetDetails
