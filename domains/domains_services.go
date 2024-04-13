package domains

/*

	Package domains provides the data structures representing entities in the project.

	Structures:
	- Service: Represents the services of the companies.
		- ID (uuid.UUID): Unique identifier for the service.
		- Name (string): name of the services
		- Description (string): description of the services
		- CompanyID (uuid.UUID): company ID associated with the services.
		- gorm.Model: Standard GORM model fields (ID, CreatedAt, UpdatedAt, DeletedAt).

	Usage:
	- Import this package to utilize the provided data structures for handling services in the project.

	Note:
	- The services structure represents information about services in the system.

*/

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Services represents information about notifications in the system.
type Services struct {
	ID          uuid.UUID `gorm:"column:id; primaryKey; type:uuid; not null;"` // Unique identifier for the notification
	Name        string    `gorm:"column:name; not null"`                       // name of the services
	Description string    `gorm:"column:description; not null;"`               // description of the services
	CompanyID   uuid.UUID `gorm:"column:company_id; type:uuid; not null;"`     // company ID associated with the services
	gorm.Model
}
