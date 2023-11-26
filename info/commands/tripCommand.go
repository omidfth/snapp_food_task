package commands

type ChangeStatus struct {
	TripID uint `json:"trip_id"`
	Status uint `json:"status"`
}
