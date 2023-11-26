package models

import "gorm.io/gorm"

type Agent struct {
	gorm.Model
	Name     string
	OrderID  uint
	HasOrder bool `gorm:"default:false"`
}
