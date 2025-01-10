package models

import "time"

// Order represents a customer's order.
type Order struct {
	ID           string    `json:"id"`
	CustomerName string    `json:"customer_name"`
	Status       string    `json:"status"`
	TotalPrice   string    `json:"total_price"`
	Products     []Product `json:"products"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	IsDeleted    bool      `json:"is_deleted"`
}
