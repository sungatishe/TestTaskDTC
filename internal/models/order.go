package models

import "time"

// Order модель для заказа
// @Description Order struct
// @example {"customer_name": "John Doe", "status": "pending", "total_price": 100.5, "product_id": 1}
type Order struct {
	ID           int       `swaggerignore:"true" ,json:"id"`
	CustomerName string    `json:"customer_name" example:"John Doe"`
	Status       string    `json:"status" example:"pending"`
	TotalPrice   float64   `json:"total_price" example:"100.5"`
	ProductID    int       `json:"product_id" example:"1"`
	CreatedAt    time.Time `swaggerignore:"true" ,json:"created_at"`
	UpdatedAt    time.Time `swaggerignore:"true" ,json:"updated_at"`
	IsDeleted    bool      `swaggerignore:"true" ,json:"is_deleted"`
}
