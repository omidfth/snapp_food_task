package dto

import "snapp_food_task/info/internal/repositories/models"

type ReportResult struct {
	VendorName string  `json:"vendor_name"`
	Total      float64 `json:"total"`
}

func NewReportResult(result *models.ReportResult) *ReportResult {
	return &ReportResult{
		VendorName: result.VendorName,
		Total:      result.Total,
	}
}
