package models

// Product represents a single product within an order
type Product struct {
	ID       int     `swaggerignore:"true" ,json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
