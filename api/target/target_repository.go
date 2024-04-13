package target

import (
	"labs/domains"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Database represents the database instance for the services package.
type Database struct {
	DB *gorm.DB
}

// NewTargetRepository performs automatic migration of target-related structures in the database.
func NewTargetRepository(db *gorm.DB) {
	if err := db.AutoMigrate(&domains.Target{}); err != nil {
		logrus.Fatal("An error occurred during automatic migration of the Target structure. Error: ", err)
	}
}

// ReadAllPagination retrieves a paginated list of services based on company ID, limit, and offset.
func ReadAllPagination(db *gorm.DB, model []domains.Target, modelID uuid.UUID, limit, offset int) ([]domains.Target, error) {
	err := db.Where("company_id = ? ", modelID).Limit(limit).Offset(offset).Find(&model).Error
	return model, err
}

// ReadAllList retrieves a list of services based on company ID.
func ReadAllList(db *gorm.DB, model []domains.Target, modelID uuid.UUID) ([]domains.Target, error) {
	err := db.Where("company_id = ? ", modelID).Find(&model).Error
	return model, err
}

// ReadByID retrieves a service by its unique identifier.
func ReadByID(db *gorm.DB, model domains.Target, id uuid.UUID) (domains.Target, error) {
	err := db.First(&model, id).Error
	return model, err
}
