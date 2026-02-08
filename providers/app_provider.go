package providers

import (
	"go-api/handler"
	"go-api/repository"
	"go-api/service"
)

type Handlers struct {
	CategoryHandler    *handler.CategoryHandler
	ProductHandler     *handler.ProductHandler
	TransactionHandler *handler.TransactionHandler
	ReportHandler      *handler.ReportHandler
}

func RegisterServices() *Handlers {
	// Category layer
	categoryRepo := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// Product layer
	productRepo := repository.NewProductRepository()
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// Transaction layer
	transactionRepo := repository.NewTransactionRepository()
	transactionService := service.NewTransactionService(transactionRepo)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	// Report layer
	reportRepo := repository.NewReportRepository()
	reportService := service.NewReportService(reportRepo)
	reportHandler := handler.NewReportHandler(reportService)

	return &Handlers{
		CategoryHandler:    categoryHandler,
		ProductHandler:     productHandler,
		TransactionHandler: transactionHandler,
		ReportHandler:      reportHandler,
	}
}
