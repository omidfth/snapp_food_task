package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"snapp_food_task/info/internal/repositories/models"
	"snapp_food_task/info/internal/services/dto"
	"testing"
)

type mockDelayRepository struct {
	mock.Mock
}

func (m *mockDelayRepository) Migrate() error {
	args := m.Called()
	return args.Error(0)
}

func (m *mockDelayRepository) Report(order *models.Order, trip *models.Trip) (*models.DelayReport, *models.EstimateDeliveredTime) {
	args := m.Called(order, trip)
	return args.Get(0).(*models.DelayReport), args.Get(1).(*models.EstimateDeliveredTime)
}

func (m *mockDelayRepository) FetchReports() []models.ReportResult {
	args := m.Called()
	return args.Get(0).([]models.ReportResult)
}

func TestDelayService_Report(t *testing.T) {
	mockRepo := new(mockDelayRepository)

	service := NewDelayService(mockRepo)

	order := &models.Order{}
	trip := &models.Trip{}

	mockDelayReport := &models.DelayReport{}
	mockEstimateTime := &models.EstimateDeliveredTime{}
	mockRepo.On("Report", order, trip).Return(mockDelayReport, mockEstimateTime)

	delayReport, estimateTime := service.Report(order, trip)

	assert.Equal(t, dto.NewDelayReport(mockDelayReport), delayReport)
	assert.Equal(t, dto.NewEstimateDeliveredTime(mockEstimateTime), estimateTime)

	mockRepo.AssertExpectations(t)
}

func TestDelayService_FetchReports(t *testing.T) {
	mockRepo := new(mockDelayRepository)

	service := NewDelayService(mockRepo)

	mockDelayReports := []models.ReportResult{{VendorName: "Vendor1", Total: 10}, {VendorName: "Vendor2", Total: 15}}
	mockRepo.On("FetchReports").Return(mockDelayReports)

	reports := service.FetchReports()

	var expectedReports []dto.ReportResult
	for _, report := range mockDelayReports {
		expectedReports = append(expectedReports, *dto.NewReportResult(&report))
	}
	assert.Equal(t, expectedReports, reports)

	mockRepo.AssertExpectations(t)
}
