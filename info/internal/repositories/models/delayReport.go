package models

import (
	"gorm.io/gorm"
	"time"
)

type DelayReport struct {
	gorm.Model
	VendorID uint `json:"vendor_id"`
	TripID   uint `json:"trip_id"`
	OrderID  uint `json:"order_id"`
}

type EstimateDeliveredTime struct {
	EstimatedTime time.Time `json:"estimated_time"`
}
