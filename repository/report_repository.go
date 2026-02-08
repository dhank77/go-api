package repository

import (
	"context"
	"time"

	"go-api/database"
	"go-api/models"
)

type ReportRepository struct{}

func NewReportRepository() *ReportRepository {
	return &ReportRepository{}
}

func (r *ReportRepository) GetSalesReport(startDate, endDate time.Time) (*models.SalesReport, error) {
	var totalRevenue, totalTransactions int
	err := database.Conn.QueryRow(context.Background(),
		`SELECT COALESCE(SUM(total_amount), 0), COUNT(*) 
		 FROM transactions 
		 WHERE created_at >= $1 AND created_at < $2`,
		startDate, endDate).Scan(&totalRevenue, &totalTransactions)
	if err != nil {
		return nil, err
	}

	var topProduct *models.TopProduct
	var productName string
	var qtySold int

	err = database.Conn.QueryRow(context.Background(),
		`SELECT p.name, COALESCE(SUM(td.quantity), 0) as total_qty
		 FROM transaction_details td
		 JOIN products p ON td.product_id = p.id
		 JOIN transactions t ON td.transaction_id = t.id
		 WHERE t.created_at >= $1 AND t.created_at < $2
		 GROUP BY p.id, p.name
		 ORDER BY total_qty DESC
		 LIMIT 1`,
		startDate, endDate).Scan(&productName, &qtySold)

	if err == nil && qtySold > 0 {
		topProduct = &models.TopProduct{
			Name:    productName,
			QtySold: qtySold,
		}
	}

	return &models.SalesReport{
		TotalRevenue:      totalRevenue,
		TotalTransactions: totalTransactions,
		TopProduct:        topProduct,
	}, nil
}
