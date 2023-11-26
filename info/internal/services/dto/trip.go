package dto

import "snapp_food_task/info/internal/repositories/models"

type Trip struct {
	ID         uint   `json:"id"`
	OrderID    uint   `json:"order_id"`
	StatusType uint   `json:"statusType"`
	Status     string `json:"status"`
}

func NewTrip(trip *models.Trip) *Trip {
	return &Trip{
		ID:         trip.ID,
		OrderID:    trip.ID,
		StatusType: uint(trip.Status),
		Status:     trip.Status.ToString(),
	}
}
