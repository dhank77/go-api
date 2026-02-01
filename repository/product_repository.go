package repository

import (
	"go-api/database"
	"go-api/models"
)

type ProductRepository struct{}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r *ProductRepository) FindAll() ([]models.Product, error) {
	rows, err := database.DB.Query("SELECT id, name, price, stock, category_id FROM products ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepository) FindByID(id int) (*models.Product, error) {
	var p models.Product
	err := database.DB.QueryRow("SELECT id, name, price, stock, category_id FROM products WHERE id = $1", id).
		Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) FindByIDWithCategory(id int) (*models.ProductDetail, error) {
	var p models.ProductDetail
	err := database.DB.QueryRow(
		`SELECT p.id, p.name, p.price, p.stock, p.category_id, c.name as category_name 
		 FROM products p 
		 JOIN categories c ON p.category_id = c.id 
		 WHERE p.id = $1`, id).
		Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID, &p.CategoryName)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) Create(p *models.Product) error {
	return database.DB.QueryRow(
		"INSERT INTO products (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id",
		p.Name, p.Price, p.Stock, p.CategoryID).Scan(&p.ID)
}

func (r *ProductRepository) Update(p *models.Product) error {
	_, err := database.DB.Exec(
		"UPDATE products SET name = $1, price = $2, stock = $3, category_id = $4 WHERE id = $5",
		p.Name, p.Price, p.Stock, p.CategoryID, p.ID)
	return err
}

func (r *ProductRepository) Delete(id int) error {
	_, err := database.DB.Exec("DELETE FROM products WHERE id = $1", id)
	return err
}
