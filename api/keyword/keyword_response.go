package keyword

import (
	"time"

	"github.com/google/uuid"
)

// @Description	KeywordIn represents the input structure for creating a new role.
type KeywordIn struct {
	Name string `json:"name" binding:"required,min=2,max=40"` // Name is the name of the role. It is required and should be between 2 and 40 characters.
} //@name KeywordIn

// @Description	KeywordsCount represents the count of keywords.
type KeywordsCount struct {
	Count uint `json:"count"` // Count is the number of keywords.
} //@name KeywordsCount

// @Description	KeywordsDetails represents detailed information about a specific keyword.
type KeywordsDetails struct {
	ID        uuid.UUID `json:"id"`        // ID is the unique identifier for the keyword.
	Name      string    `json:"name"`      // name is the name of the keyword.
	CreatedAt time.Time `json:"createdAt"` // CreatedAt is the timestamp indicating when the keyword was created.
} //@name KeywordsDetails

// @Description	KeywordsPagination represents the paginated list of keywords.
type KeywordsPagination struct {
	Items      []KeywordsTable `json:"items"`      // Items is a slice containing individual service details.
	Page       uint            `json:"page"`       // Page is the current page number in the pagination.
	Limit      uint            `json:"limit"`      // Limit is the maximum number of items per page in the pagination.
	TotalCount uint            `json:"totalCount"` // TotalCount is the total number of services in the entire list.
} //@name KeywordsPagination

// @Description	KeywordsTable represents a single service entry in a table.
type KeywordsTable struct {
	ID        uuid.UUID `json:"id"`                                   // ID is the unique identifier for the user.
	Name      string    `json:"name" binding:"required,min=3,max=30"` // name is the  name of the service. It is required and should be between 3 and 30 characters.
	CreatedAt time.Time `json:"createdAt"`                            // CreatedAt is the timestamp indicating when the service entry was created.
} //@name ServicesTable

// @Description	KeywordsList represents a simplified version of the keyword for listing purposes.
type KeywordsList struct {
	ID   uuid.UUID `json:"id"`   // ID is the unique identifier for the service.
	Name string    `json:"name"` // Name is the full name of the service.
} //@name KeywordsList
