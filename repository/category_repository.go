package repository

import (
	"context"

	"go-api/database"
	"go-api/models"
)

type CategoryRepository struct{}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

func (r *CategoryRepository) FindAll() ([]models.Category, error) {
	rows, err := database.Conn.Query(context.Background(),
		"SELECT id, name, description FROM categories ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		if err := rows.Scan(&cat.ID, &cat.Name, &cat.Description); err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}
	return categories, nil
}

func (r *CategoryRepository) FindByID(id int) (*models.Category, error) {
	var cat models.Category
	err := database.Conn.QueryRow(context.Background(),
		"SELECT id, name, description FROM categories WHERE id = $1", id).
		Scan(&cat.ID, &cat.Name, &cat.Description)
	if err != nil {
		return nil, err
	}
	return &cat, nil
}

func (r *CategoryRepository) Create(cat *models.Category) error {
	return database.Conn.QueryRow(context.Background(),
		"INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id",
		cat.Name, cat.Description).Scan(&cat.ID)
}

func (r *CategoryRepository) Update(cat *models.Category) error {
	_, err := database.Conn.Exec(context.Background(),
		"UPDATE categories SET name = $1, description = $2 WHERE id = $3",
		cat.Name, cat.Description, cat.ID)
	return err
}

func (r *CategoryRepository) Delete(id int) error {
	_, err := database.Conn.Exec(context.Background(),
		"DELETE FROM categories WHERE id = $1", id)
	return err
}
