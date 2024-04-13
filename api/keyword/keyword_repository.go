package keyword

import (
	"labs/domains"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Database represents the database instance for the notifications package.
type Database struct {
	DB *gorm.DB
}

// NewKeywordRepository performs automatic migration of keyword-related structures in the database.
func NewKeywordRepository(db *gorm.DB) {
	if err := db.AutoMigrate(&domains.Keyword{}); err != nil {
		logrus.Fatal("An error occurred during automatic migration of the keyword structure. Error: ", err)
	}
}

// ReadAllList retrieves a list of services based on company ID.
func ReadAllList(db *gorm.DB, model []domains.Keyword, modelID uuid.UUID) ([]domains.Keyword, error) {
	err := db.Where("service_id = ? ", modelID).Find(&model).Error
	return model, err
}

// ReadByID retrieves a keyword by its unique identifier.
func ReadByID(db *gorm.DB, model domains.Keyword, id uuid.UUID) (domains.Keyword, error) {
	err := db.First(&model, id).Error
	return model, err
}

// ReadAllPagination retrieves a paginated list of services based on company ID, limit, and offset.
func ReadAllPagination(db *gorm.DB, model []domains.Keyword, modelID uuid.UUID, limit, offset int) ([]domains.Keyword, error) {
	err := db.Where("service_id = ? ", modelID).Limit(limit).Offset(offset).Find(&model).Error
	return model, err
}
