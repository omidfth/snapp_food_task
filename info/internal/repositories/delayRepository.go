package repositories

import (
	"gorm.io/gorm"
	"snapp_food_task/info/constants/times"
	"snapp_food_task/info/internal/repositories/models"
	"snapp_food_task/info/internal/types/tripTypes"
	"time"
)

type DelayRepository interface {
	Migrate() error
	Report(order *models.Order, trip *models.Trip) (*models.DelayReport, *models.EstimateDeliveredTime)
	FetchReports() []models.ReportResult
}

type delayRepository struct {
	db *gorm.DB
}

func NewDelayRepository(db *gorm.DB) DelayRepository {
	return &delayRepository{db: db}
}

func (r delayRepository) Migrate() error {
	return r.db.AutoMigrate(models.DelayReport{})
}

func (r delayRepository) Report(order *models.Order, trip *models.Trip) (*models.DelayReport, *models.EstimateDeliveredTime) {
	remainingTime := order.EndTime.Sub(time.Now()).Seconds()
	if remainingTime > 0 {
		return nil, nil
	}
	if trip == nil || trip.Status.IsEqual(tripTypes.DELIVERED) {
		delayReport := models.DelayReport{
			TripID:   order.TripID,
			OrderID:  order.ID,
			VendorID: order.VendorID,
		}

		r.db.Create(&delayReport)

		return &delayReport, &models.EstimateDeliveredTime{
			EstimatedTime: time.Now().Add(times.ESTIMATE_NEW_DELIVERY_TIME),
		}
	}
	delayReport := models.DelayReport{
		TripID:   order.TripID,
		OrderID:  order.ID,
		VendorID: order.VendorID,
	}

	r.db.Create(&delayReport)
	return &delayReport, nil
}

func (r delayRepository) FetchReports() []models.ReportResult {
	var reports []models.ReportResult
	r.db.Debug().Table("vendors").
		Order("sum(orders.delay_time) desc").
		Select("vendors.name as vendor_name, sum(orders.delay_time) as total").
		Joins("JOIN orders ON vendors.id = orders.vendor_id").
		Where("orders.delivered_time > ? AND orders.delivered_time < ?", time.Now().AddDate(0, 0, -7), time.Now()).
		Group("vendors.id").
		Scan(&reports)
	return reports
}
