package domains

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Target represents information about target in the system.
type Target struct {
	ID        uuid.UUID `gorm:"column:id; primaryKey; type:uuid; not null;"` // Unique identifier for the notification
	Name      string    `gorm:"column:name; not null"`                       // name of the services
	CompanyID uuid.UUID `gorm:"column:company_id; type:uuid; not null;"`     // company ID associated with the services
	gorm.Model
}
