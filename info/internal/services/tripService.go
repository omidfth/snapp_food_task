package services

import (
	"snapp_food_task/info/internal/repositories"
	"snapp_food_task/info/internal/repositories/models"
	"snapp_food_task/info/internal/services/dto"
	"snapp_food_task/info/internal/types/tripTypes"
)

type TripService interface {
	GetByID(id uint) (*models.Trip, error)
	GetByOrderID(id uint) (*models.Trip, error)
	ChangeStatus(trip *models.Trip, status models.TripStatus) (*dto.Trip, error)
}

type tripService struct {
	tripRepository  repositories.TripRepository
	orderRepository repositories.OrderRepository
}

func NewTripService(
	tripRepository repositories.TripRepository,
	orderRepository repositories.OrderRepository,
) TripService {
	return &tripService{
		tripRepository:  tripRepository,
		orderRepository: orderRepository,
	}
}

func (s tripService) ChangeStatus(trip *models.Trip, status models.TripStatus) (*dto.Trip, error) {
	t := s.tripRepository.ChangeStatus(trip, status)
	if t.Status.IsEqual(tripTypes.DELIVERED) {
		if err := s.orderRepository.SetDeliveredTime(t); err != nil {
			return nil, err
		}
	}
	data := dto.NewTrip(t)
	return data, nil
}

func (s tripService) GetByID(id uint) (*models.Trip, error) {
	return s.tripRepository.GetByID(id)
}

func (s tripService) GetByOrderID(id uint) (*models.Trip, error) {
	return s.tripRepository.GetByOrderID(id)
}
