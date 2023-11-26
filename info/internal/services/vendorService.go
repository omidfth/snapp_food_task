package services

import (
	"snapp_food_task/info/internal/repositories"
	"snapp_food_task/info/internal/repositories/models"
)

type VendorService interface {
	GetByID(id string) (*models.Vendor, error)
}

type vendorService struct {
	repository repositories.VendorRepository
}

func NewVendorService(repository repositories.VendorRepository) VendorService {
	return &vendorService{repository: repository}
}

func (s vendorService) GetByID(id string) (*models.Vendor, error) {
	return s.repository.GetByID(id)
}
