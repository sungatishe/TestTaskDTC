package models

import "time"

// Order represents a customer's order.
type Order struct {
	ID           int       `json:"id"`
	CustomerName string    `json:"customer_name"`
	Status       string    `json:"status"`
	TotalPrice   float64   `json:"total_price"`
	ProductID    int       `json:"product_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	IsDeleted    bool      `json:"is_deleted"`
}
