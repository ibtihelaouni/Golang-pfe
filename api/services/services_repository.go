package services

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

// NewServiceRepository performs automatic migration of service-related structures in the database.
func NewServiceRepository(db *gorm.DB) {
	if err := db.AutoMigrate(&domains.Services{}); err != nil {
		logrus.Fatal("An error occurred during automatic migration of the Service structure. Error: ", err)
	}
}

// ReadAllPagination retrieves a paginated list of services based on company ID, limit, and offset.
func ReadAllPagination(db *gorm.DB, model []domains.Services, modelID uuid.UUID, limit, offset int) ([]domains.Services, error) {
	err := db.Where("company_id = ? ", modelID).Limit(limit).Offset(offset).Find(&model).Error
	return model, err
}

// ReadAllList retrieves a list of services based on company ID.
func ReadAllList(db *gorm.DB, model []domains.Services, modelID uuid.UUID) ([]domains.Services, error) {
	err := db.Where("company_id = ? ", modelID).Find(&model).Error
	return model, err
}

// ReadByID retrieves a service by its unique identifier.
func ReadByID(db *gorm.DB, model domains.Services, id uuid.UUID) (domains.Services, error) {
	err := db.First(&model, id).Error
	return model, err
}
