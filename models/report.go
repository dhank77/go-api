package models

type TopProduct struct {
	Name    string `json:"name"`
	QtySold int    `json:"qty_sold"`
}

type SalesReport struct {
	TotalRevenue      int         `json:"total_revenue"`
	TotalTransactions int         `json:"total_transactions"`
	TopProduct        *TopProduct `json:"top_product"`
}
