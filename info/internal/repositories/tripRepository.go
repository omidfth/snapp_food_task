package repositories

import (
	"errors"
	"gorm.io/gorm"
	"snapp_food_task/info/constants/errorKeys"
	"snapp_food_task/info/constants/times"
	"snapp_food_task/info/internal/repositories/models"
	"snapp_food_task/info/internal/types/tripTypes"
	"time"
)

type TripRepository interface {
	Migrate() error
	Append(order *models.Order, onEnd func(trip *models.Trip)) *models.Trip
	ChangeStatus(trip *models.Trip, status models.TripStatus) *models.Trip
	GetByID(id uint) (*models.Trip, error)
	GetByOrderID(orderID uint) (*models.Trip, error)
}

type tripRepository struct {
	db *gorm.DB
}

func NewTripRepository(db *gorm.DB) TripRepository {
	return &tripRepository{db: db}
}

func (r tripRepository) Migrate() error {
	return r.db.AutoMigrate(models.Trip{})
}

func (r tripRepository) Append(order *models.Order, onEnd func(trip *models.Trip)) *models.Trip {
	trip := models.Trip{
		OrderID: order.ID,
		Status:  tripTypes.ASSIGNED,
	}
	r.db.Create(&trip)
	time.AfterFunc(times.CHANGE_TRIP_STATE_TIME, func() { r.changeAutoStatus(&trip, onEnd) })
	return &trip
}

func (r tripRepository) changeAutoStatus(trip *models.Trip, onEnd func(trip *models.Trip)) {
	if trip.Status < tripTypes.DELIVERED {
		trip.Status++
		r.db.Save(&trip)
	}
	if trip.Status.IsNot(tripTypes.DELIVERED) {
		time.AfterFunc(times.CHANGE_TRIP_STATE_TIME, func() { r.changeAutoStatus(trip, onEnd) })
		return
	}
	onEnd(trip)
}

func (r tripRepository) ChangeStatus(trip *models.Trip, status models.TripStatus) *models.Trip {
	trip.Status = status
	r.db.Save(&trip)
	return trip
}

func (r tripRepository) GetByID(id uint) (*models.Trip, error) {
	var trip models.Trip
	if r.db.Where("id=?", id).First(&trip).RowsAffected < 1 {
		return nil, errors.New(errorKeys.TRIP_NOT_FOUND)
	}
	return &trip, nil
}

func (r tripRepository) GetByOrderID(orderID uint) (*models.Trip, error) {
	var trip models.Trip
	if r.db.Where("order_id=?", orderID).First(&trip).RowsAffected < 1 {
		return nil, errors.New(errorKeys.TRIP_NOT_FOUND)
	}
	return &trip, nil
}
