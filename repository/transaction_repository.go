package repository

import (
	"context"
	"fmt"

	"go-api/database"
	"go-api/models"
)

type TransactionRepository struct{}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	tx, err := database.Conn.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	totalAmount := 0
	details := make([]models.TransactionDetail, 0)

	for _, item := range items {
		var productPrice float64
		var stock int
		var productName string

		err := tx.QueryRow(context.Background(),
			"SELECT name, price, stock FROM products WHERE id = $1", item.ProductID).
			Scan(&productName, &productPrice, &stock)
		if err != nil {
			return nil, fmt.Errorf("product id %d not found: %v", item.ProductID, err)
		}

		if stock < item.Quantity {
			return nil, fmt.Errorf("insufficient stock for product %s (available: %d, requested: %d)",
				productName, stock, item.Quantity)
		}

		subtotal := int(productPrice) * item.Quantity
		totalAmount += subtotal

		_, err = tx.Exec(context.Background(),
			"UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		details = append(details, models.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID int
	err = tx.QueryRow(context.Background(),
		"INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id", totalAmount).
		Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	if len(details) > 0 {
		query := "INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES "
		args := []interface{}{}
		for i := range details {
			details[i].TransactionID = transactionID
			if i > 0 {
				query += ", "
			}
			paramIdx := i * 4
			query += fmt.Sprintf("($%d, $%d, $%d, $%d)", paramIdx+1, paramIdx+2, paramIdx+3, paramIdx+4)
			args = append(args, transactionID, details[i].ProductID, details[i].Quantity, details[i].Subtotal)
		}

		_, err = tx.Exec(context.Background(), query, args...)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(context.Background()); err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}
