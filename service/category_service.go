package service

import (
	"go-api/models"
	"go-api/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.repo.FindAll()
}

func (s *CategoryService) GetByID(id int) (*models.Category, error) {
	return s.repo.FindByID(id)
}

func (s *CategoryService) Create(cat *models.Category) error {
	return s.repo.Create(cat)
}

func (s *CategoryService) Update(cat *models.Category) error {
	return s.repo.Update(cat)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
