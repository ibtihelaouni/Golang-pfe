package domains

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Notifications represents information about notifications in the system.
type Keyword struct {
	ID        uuid.UUID `gorm:"column:id; primaryKey; type:uuid; not null;"` // Unique identifier for the keyword
	Name      string    `gorm:"column:name; not null"`                       // name of the keyword
	ServiceID uuid.UUID `gorm:"column:service_id;type:uuid; not null;"`      // service ID associated with the keyword
	gorm.Model
}

// CheckServiceBelonging checks if the service belongs to the specified company.
func CheckServiceBelonging(db *gorm.DB, pathCompanyID, sessionServiceID, sessionCompanyID uuid.UUID) error {
	if pathCompanyID != sessionCompanyID {
		return errors.New("error occurred when attempting to verify service belonging")
	}
	return db.Select("id, company_id").Where("id = ? AND company_id = ?", sessionServiceID, pathCompanyID).First(&Services{}).Error
}
