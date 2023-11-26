package services

import (
	"log"
	"snapp_food_task/info/constants/times"
	"snapp_food_task/info/internal/repositories"
	"snapp_food_task/info/internal/repositories/models"
	"snapp_food_task/info/internal/services/dto"
	"time"
)

type OrderService interface {
	GetByID(id string) (*models.Order, error)
	AppendNewOrder(vendor *models.Vendor) *dto.Order
}

type orderService struct {
	orderRepository repositories.OrderRepository
	tripRepository  repositories.TripRepository
}

func NewOrderService(orderRepository repositories.OrderRepository, tripRepository repositories.TripRepository) OrderService {
	return &orderService{
		orderRepository: orderRepository,
		tripRepository:  tripRepository,
	}
}

func (s orderService) AppendNewOrder(vendor *models.Vendor) *dto.Order {
	order := s.orderRepository.AppendNewOrder(vendor)

	//assign trip after ? minutes
	time.AfterFunc(times.ASSIGN_TRIP_TIME, func() {
		trip := s.tripRepository.Append(order, func(trip *models.Trip) {
			err := s.orderRepository.SetDeliveredTime(trip)
			if err != nil {
				log.Println("set delivered time error", err.Error())
			}
		})
		s.orderRepository.AssignTrip(trip, order)
	})
	orderDto := dto.NewOrder(order)
	return orderDto
}

func (s orderService) GetByID(id string) (*models.Order, error) {
	return s.orderRepository.GetByID(id)
}
