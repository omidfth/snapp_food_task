package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"snapp_food_task/info/internal/repositories/models"
	"snapp_food_task/info/internal/services/dto"
	"testing"
)

type mockOrderRepository struct {
	mock.Mock
}

func (m *mockOrderRepository) Migrate() error {
	args := m.Called()
	return args.Error(0)
}

func (m *mockOrderRepository) GetByID(id string) (*models.Order, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Order), args.Error(1)
}

func (m *mockOrderRepository) AppendNewOrder(vendor *models.Vendor) *models.Order {
	args := m.Called(vendor)
	return args.Get(0).(*models.Order)
}

func (m *mockOrderRepository) AssignTrip(trip *models.Trip, order *models.Order) *models.Order {
	args := m.Called(trip, order)
	return args.Get(0).(*models.Order)
}

func (m *mockOrderRepository) SetDeliveredTime(trip *models.Trip) error {
	args := m.Called(trip)
	return args.Error(0)
}

func createTestVendor() *models.Vendor {
	return &models.Vendor{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "TestVendor",
	}
}

func TestOrderService_AppendNewOrder(t *testing.T) {
	mockOrderRepo := new(mockOrderRepository)
	mockTripRepo := new(mockTripRepository)

	service := NewOrderService(mockOrderRepo, mockTripRepo)

	vendor := createTestVendor()

	mockOrder := &models.Order{
		Model: gorm.Model{
			ID: 1,
		},
	}
	mockOrderRepo.On("AppendNewOrder", vendor).Return(mockOrder)

	mockTrip := &models.Trip{
		Model: gorm.Model{
			ID: 1,
		},
		OrderID: 1,
		Status:  0,
	}
	mockTripRepo.On("Append", mockOrder, mock.Anything).Return(mockTrip)

	mockOrderRepo.On("AssignTrip", mockTrip, mockOrder).Return(mockOrder)

	result := service.AppendNewOrder(vendor)

	assert.Equal(t, dto.NewOrder(mockOrder), result)
}

func TestOrderService_GetByID(t *testing.T) {
	mockRepo := new(mockOrderRepository)

	service := NewOrderService(mockRepo, nil)

	id := "1"

	mockOrder := &models.Order{}
	mockRepo.On("GetByID", id).Return(mockOrder, nil)

	result, err := service.GetByID(id)

	assert.NoError(t, err)
	assert.Equal(t, mockOrder, result)

	mockRepo.AssertExpectations(t)
}
