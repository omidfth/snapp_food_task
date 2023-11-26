package repositories

import (
	"errors"
	"gorm.io/gorm"
	"snapp_food_task/info/constants/errorKeys"
	"snapp_food_task/info/internal/repositories/models"
)

type VendorRepository interface {
	Migrate() error
	GetByID(id string) (*models.Vendor, error)
}

type vendorRepository struct {
	db *gorm.DB
}

func NewVendorRepository(db *gorm.DB) VendorRepository {
	return &vendorRepository{db: db}
}

func (r vendorRepository) Migrate() error {
	return r.db.AutoMigrate(models.Vendor{})
}

func (r vendorRepository) GetByID(id string) (*models.Vendor, error) {
	var vendor models.Vendor
	if r.db.Where("id=?", id).First(&vendor).RowsAffected < 1 {
		return nil, errors.New(errorKeys.VENDOR_NOT_FOUND)
	}
	return &vendor, nil
}
