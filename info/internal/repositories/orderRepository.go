package repositories

import (
	"errors"
	"gorm.io/gorm"
	"snapp_food_task/info/constants/errorKeys"
	"snapp_food_task/info/constants/times"
	"snapp_food_task/info/internal/repositories/models"
	"time"
)

type OrderRepository interface {
	Migrate() error
	GetByID(id string) (*models.Order, error)
	AppendNewOrder(vendor *models.Vendor) *models.Order
	AssignTrip(trip *models.Trip, order *models.Order) *models.Order
	SetDeliveredTime(trip *models.Trip) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r orderRepository) Migrate() error {
	return r.db.AutoMigrate(models.Order{})
}

func (r orderRepository) AppendNewOrder(vendor *models.Vendor) *models.Order {
	startTime := time.Now()
	endTime := startTime.Add(times.DELIVERY_TIME)
	deliveryTime := endTime.Sub(startTime).Seconds()

	order := models.Order{
		DeliveryTime: deliveryTime,
		StartTime:    startTime,
		EndTime:      endTime,
		VendorID:     vendor.ID,
	}

	r.db.Create(&order)
	return &order
}

func (r orderRepository) AssignTrip(trip *models.Trip, order *models.Order) *models.Order {
	order.TripID = trip.ID
	r.db.Save(&order)
	return order
}

func (r orderRepository) GetByID(id string) (*models.Order, error) {
	var order models.Order
	if r.db.Where("id=?", id).First(&order).RowsAffected < 1 {
		return nil, errors.New(errorKeys.ORDER_NOT_FOUND)
	}
	return &order, nil
}

func (r orderRepository) SetDeliveredTime(trip *models.Trip) error {
	var order models.Order
	if r.db.Where("id=?", trip.OrderID).First(&order).RowsAffected < 1 {
		return errors.New(errorKeys.ORDER_NOT_FOUND)
	}
	order.DeliveredTime = time.Now()
	delayTime := order.DeliveredTime.Sub(order.EndTime).Seconds()
	if delayTime > 0 {
		order.DelayTime = delayTime
	}
	r.db.Save(&order)
	return nil
}
