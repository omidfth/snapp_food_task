package dto

import "snapp_food_task/info/internal/repositories/models"

type Vendor struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewVendor(vendor *models.Vendor) *Vendor {
	return &Vendor{
		ID:   vendor.ID,
		Name: vendor.Name,
	}
}
