package models

import "gorm.io/gorm"

type Vendor struct {
	gorm.Model
	Name   string   `json:"name"`
	Orders []*Order `json:"orders"`
}
