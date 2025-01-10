package repository

import (
	"TestTask/internal/models"
	"database/sql"
	"errors"
	"fmt"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(product *models.Product) error {
	query := `
		INSERT INTO products (name, price, quantity) VALUES ($1, $2, $3)
	`
	_, err := r.db.Exec(query, product.Name, product.Price, product.Quantity)
	if err != nil {
		return fmt.Errorf("could not create product: %v", err)
	}
	return nil
}

func (r *ProductRepository) GetProductByID(productID int) (*models.Product, error) {
	query := `
		SELECT id, name, price, quantity 
		FROM products
		WHERE id = $1
	`
	row := r.db.QueryRow(query, productID)

	var product models.Product
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get product by ID: %w", err)
	}

	return &product, nil
}

func (r *ProductRepository) UpdateProduct(product *models.Product) error {
	query := `
		UPDATE products
		SET name = $1, price = $2, quantity = $3 
		WHERE id = $4
	`
	result, err := r.db.Exec(query, product.Name, product.Price, product.Quantity, product.ID)
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no product found with id %d", product.ID)
	}

	return nil
}

func (r *ProductRepository) DeleteProductByID(productID int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := r.db.Exec(query, productID)
	if err != nil {
		return fmt.Errorf("failed tot delete product: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no product found with id %d", productID)
	}

	return nil
}

func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
	query := "SELECT id, name, price, quantity FROM products"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product row: %w", err)
		}
		products = append(products, product)
	}

	return products, nil
}
