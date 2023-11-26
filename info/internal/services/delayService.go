package services

import (
	"snapp_food_task/info/internal/repositories"
	"snapp_food_task/info/internal/repositories/models"
	"snapp_food_task/info/internal/services/dto"
)

type DelayService interface {
	Report(order *models.Order, trip *models.Trip) (*dto.DelayReport, *dto.EstimateDeliveredTime)
	FetchReports() []dto.ReportResult
}

type delayService struct {
	delayRepository repositories.DelayRepository
}

func NewDelayService(delayRepository repositories.DelayRepository) DelayService {
	return &delayService{delayRepository: delayRepository}
}

func (s delayService) Report(order *models.Order, trip *models.Trip) (*dto.DelayReport, *dto.EstimateDeliveredTime) {
	report, estimateTime := s.delayRepository.Report(order, trip)
	if estimateTime == nil && report == nil {
		return nil, nil
	}
	delayReport := dto.NewDelayReport(report)
	if estimateTime == nil {
		return delayReport, nil
	}
	return delayReport, dto.NewEstimateDeliveredTime(estimateTime)
}

func (s delayService) FetchReports() []dto.ReportResult {
	reports := s.delayRepository.FetchReports()
	var ret []dto.ReportResult
	for _, report := range reports {
		ret = append(ret, *dto.NewReportResult(&report))
	}
	return ret
}
