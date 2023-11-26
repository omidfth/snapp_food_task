package dto

import (
	"snapp_food_task/info/internal/repositories/models"
	"time"
)

type DelayReport struct {
	ID       uint `json:"id"`
	TripID   uint `json:"trip_id"`
	OrderID  uint `json:"order_id"`
	VendorID uint `json:"vendor_id"`
}

func NewDelayReport(report *models.DelayReport) *DelayReport {
	return &DelayReport{
		ID:       report.ID,
		TripID:   report.TripID,
		OrderID:  report.OrderID,
		VendorID: report.VendorID,
	}
}

type EstimateDeliveredTime struct {
	EstimatedTime time.Time `json:"estimated_time"`
}

func NewEstimateDeliveredTime(deliveredTime *models.EstimateDeliveredTime) *EstimateDeliveredTime {
	return &EstimateDeliveredTime{EstimatedTime: deliveredTime.EstimatedTime}
}
