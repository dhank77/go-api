package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"go-api/service"
)

type ReportHandler struct {
	service *service.ReportService
}

func NewReportHandler(s *service.ReportService) *ReportHandler {
	return &ReportHandler{service: s}
}

func (h *ReportHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/api/report")
	path = strings.TrimPrefix(path, "/")

	switch path {
	case "hari-ini":
		h.GetTodayReport(w, r)
	case "":
		h.GetReportByDateRange(w, r)
	default:
		http.Error(w, `{"error":"Not found"}`, http.StatusNotFound)
	}
}

func (h *ReportHandler) GetTodayReport(w http.ResponseWriter, r *http.Request) {
	report, err := h.service.GetTodayReport()
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(report)
}

func (h *ReportHandler) GetReportByDateRange(w http.ResponseWriter, r *http.Request) {
	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")

	if startDateStr == "" || endDateStr == "" {
		http.Error(w, `{"error":"start_date and end_date are required"}`, http.StatusBadRequest)
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		http.Error(w, `{"error":"Invalid start_date format, use YYYY-MM-DD"}`, http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		http.Error(w, `{"error":"Invalid end_date format, use YYYY-MM-DD"}`, http.StatusBadRequest)
		return
	}

	report, err := h.service.GetReportByDateRange(startDate, endDate)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(report)
}
