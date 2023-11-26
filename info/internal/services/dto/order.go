package dto

import (
	"snapp_food_task/info/internal/repositories/models"
	"time"
)

type Order struct {
	ID            uint           `json:"id"`
	DeliveryTime  float64        `json:"delivery_time"`
	StartTime     time.Time      `json:"start_time"`
	EndTime       time.Time      `json:"end_time"`
	DeliveredTime time.Time      `json:"delivered_time"`
	DelayTime     float64        `json:"delay_time"`
	VendorID      uint           `json:"vendor_id"`
	TripID        uint           `json:"trip_id"`
	DelayReports  []*DelayReport `json:"delay_reports"`
}

func NewOrder(order *models.Order) *Order {
	return &Order{
		ID:            order.ID,
		DeliveryTime:  order.DeliveryTime,
		StartTime:     order.StartTime,
		EndTime:       order.EndTime,
		DeliveredTime: order.DeliveredTime,
		DelayTime:     order.DelayTime,
		VendorID:      order.VendorID,
		TripID:        order.TripID,
	}
}
