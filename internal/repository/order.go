package repository

import (
	"TestTask/internal/models"
	"database/sql"
	"fmt"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(order *models.Order) error {
	query := `
		INSERT INTO orders (customer_name, status, total_price, product_id)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(query, order.CustomerName, order.Status, order.TotalPrice, order.ProductID)
	if err != nil {
		return fmt.Errorf("could not create order: %v", err)
	}
	return nil
}

func (r *OrderRepository) UpdateOrder(order *models.Order) error {
	query := `
	UPDATE orders
        SET customer_name = $1, status = $2, total_price = $3, updated_at = $4
        WHERE id = $5 AND is_deleted = false
	`

	result, err := r.db.Exec(query, order.CustomerName, order.Status, order.TotalPrice, order.UpdatedAt, order.ID)
	if err != nil {
		return fmt.Errorf("could not update order: %v", err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get affected rows: %v", err)
	}

	if affectedRows == 0 {
		return fmt.Errorf("no order found with the given id")
	}

	return nil
}

func (r *OrderRepository) DeleteOrder(orderID int) error {
	query := `
		UPDATE orders
		SET is_deleted = true
		WHERE id = $1
	`
	result, err := r.db.Exec(query, orderID)
	if err != nil {
		return fmt.Errorf("could not delete order: %v", err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get affected rows: %v", err)
	}

	if affectedRows == 0 {
		return fmt.Errorf("no order found with the given id")
	}

	return nil
}

func (r *OrderRepository) GetOrderByID(orderID int) (*models.Order, error) {
	query := `
		SELECT id, customer_name, status, total_price, product_id, created_at, updated_at, is_deleted
        FROM orders
        WHERE id = $1 AND is_deleted = false
	`
	var order models.Order
	err := r.db.QueryRow(query, orderID).Scan(&order.ID, &order.CustomerName, &order.Status, &order.TotalPrice, &order.ProductID, &order.CreatedAt, &order.UpdatedAt, &order.IsDeleted)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no order found with id: %d", orderID)
		}
		return nil, fmt.Errorf("could not get order by id: %v", err)
	}
	return &order, nil
}

func (r *OrderRepository) GetOrdersByFilters(status string, minPrice, maxPrice float64) ([]models.Order, error) {
	query := `
		SELECT id, customer_name, status, total_price, created_at, updated_at, is_deleted
        FROM orders
        WHERE is_deleted = false
	`

	if status != "" {
		query += " AND status = $1"
	}
	if minPrice > 0 {
		query += " AND total_price >= $2"
	}
	if maxPrice > 0 {
		query += " AND total_price <= $3"
	}

	rows, err := r.db.Query(query, status, minPrice, maxPrice)
	if err != nil {
		return nil, fmt.Errorf("could not get orders: %v", err)
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.CustomerName, &order.Status, &order.TotalPrice, &order.CreatedAt, &order.UpdatedAt, &order.IsDeleted); err != nil {
			return nil, fmt.Errorf("could not scan order: %v", err)
		}
		orders = append(orders, order)
	}

	return orders, nil
}
