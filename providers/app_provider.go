package providers

import (
	"go-api/handler"
	"go-api/repository"
	"go-api/service"
)

type Handlers struct {
	CategoryHandler *handler.CategoryHandler
	ProductHandler  *handler.ProductHandler
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

	return &Handlers{
		CategoryHandler: categoryHandler,
		ProductHandler:  productHandler,
	}
}
