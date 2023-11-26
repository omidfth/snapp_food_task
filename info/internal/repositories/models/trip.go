package models

import (
	"gorm.io/gorm"
	"snapp_food_task/info/internal/types/tripTypes"
)

type TripStatus uint8

func (t *TripStatus) ToString() string {
	amount := uint8(*t)
	switch amount {
	case tripTypes.ASSIGNED:
		return "Assigned"
	case tripTypes.AT_VENDOR:
		return "At vendor"
	case tripTypes.PICKED:
		return "Picked"
	case tripTypes.DELIVERED:
		return "Delivered"
	}
	return ""
}

func (t *TripStatus) IsEqual(tripType uint) bool {
	ts := TripStatus(tripType)
	if ts == *t {
		return true
	}

	return false
}

func (t *TripStatus) IsNot(tripType uint) bool {
	ts := TripStatus(tripType)
	if ts != *t {
		return true
	}

	return false
}

type Trip struct {
	gorm.Model
	OrderID uint       `json:"order_id" gorm:"unique"`
	Status  TripStatus `json:"status" gorm:"type:int8"`
}
