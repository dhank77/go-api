package service

import (
	"time"

	"go-api/models"
	"go-api/repository"
)

type ReportService struct {
	repo *repository.ReportRepository
}

func NewReportService(repo *repository.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetTodayReport() (*models.SalesReport, error) {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	return s.repo.GetSalesReport(startOfDay, endOfDay)
}

func (s *ReportService) GetReportByDateRange(startDate, endDate time.Time) (*models.SalesReport, error) {
	// Add 1 day to endDate to include the entire end date
	endDate = endDate.Add(24 * time.Hour)
	return s.repo.GetSalesReport(startDate, endDate)
}
