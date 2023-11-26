package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"snapp_food_task/info/internal/repositories/models"
	"snapp_food_task/info/internal/types/tripTypes"
	"testing"
)

type mockTripRepository struct {
	mock.Mock
}

func (m *mockTripRepository) Migrate() error {
	args := m.Called()
	return args.Error(0)
}

func (m *mockTripRepository) Append(order *models.Order, onEnd func(trip *models.Trip)) *models.Trip {
	args := m.Called(order, onEnd)
	return args.Get(0).(*models.Trip)
}

func (m *mockTripRepository) ChangeStatus(trip *models.Trip, status models.TripStatus) *models.Trip {
	args := m.Called(trip, status)
	return args.Get(0).(*models.Trip)
}

func (m *mockTripRepository) GetByID(id uint) (*models.Trip, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Trip), args.Error(1)
}

func (m *mockTripRepository) GetByOrderID(orderID uint) (*models.Trip, error) {
	args := m.Called(orderID)
	return args.Get(0).(*models.Trip), args.Error(1)
}

func TestTripService_GetByID(t *testing.T) {
	mockRepo := new(mockTripRepository)

	service := NewTripService(mockRepo, nil)

	id := uint(1)

	mockTrip := &models.Trip{Model: gorm.Model{
		ID: 1,
	}, OrderID: 1, Status: tripTypes.ASSIGNED}
	mockRepo.On("GetByID", id).Return(mockTrip, nil)

	result, err := service.GetByID(id)

	assert.NoError(t, err)
	assert.Equal(t, mockTrip, result)

	mockRepo.AssertExpectations(t)
}

func TestTripService_GetByOrderID(t *testing.T) {
	mockRepo := new(mockTripRepository)

	service := NewTripService(mockRepo, nil)

	orderID := uint(1)

	mockTrip := &models.Trip{Model: gorm.Model{
		ID: 1,
	}, OrderID: 1, Status: tripTypes.ASSIGNED}
	mockRepo.On("GetByOrderID", orderID).Return(mockTrip, nil)

	result, err := service.GetByOrderID(orderID)

	assert.NoError(t, err)
	assert.Equal(t, mockTrip, result)

	mockRepo.AssertExpectations(t)
}
