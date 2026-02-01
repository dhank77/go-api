package service

import (
	"go-api/models"
	"go-api/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductService) GetByID(id int) (*models.Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) GetDetail(id int) (*models.ProductDetail, error) {
	return s.repo.FindByIDWithCategory(id)
}

func (s *ProductService) Create(p *models.Product) error {
	return s.repo.Create(p)
}

func (s *ProductService) Update(p *models.Product) error {
	return s.repo.Update(p)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
