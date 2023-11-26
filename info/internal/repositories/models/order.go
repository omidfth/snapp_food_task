package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	DeliveryTime  float64   `json:"delivery_time"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	DeliveredTime time.Time `json:"delivered_time"`
	DelayTime     float64   `json:"delay_time"`
	VendorID      uint      `json:"vendor_id"`
	TripID        uint      `json:"trip_id"`
	Vendor        *Vendor
	DelayReports  []*DelayReport `json:"delay_reports"`
}
