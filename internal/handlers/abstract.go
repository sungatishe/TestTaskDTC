package handlers

import "TestTask/internal/models"

type OrderServiceInterface interface {
	CreateOrder(order *models.Order) error
	UpdateOrder(order *models.Order) error
	DeleteOrder(orderID int) error
	GetOrderByID(orderID int) (*models.Order, error)
	GetOrdersByFilters(status string, minPrice, maxPrice float64) ([]models.Order, error)
}

type AuthServiceInterface interface {
	RegisterUser(user *models.User) error
	LoginUser(username, password string) (string, error)
}

type ProductServiceInterface interface {
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(productID int) error
	GetProductByID(productID int) (*models.Product, error)
	GetAllProducts() ([]models.Product, error)
}

type LogServiceInterface interface {
	CreateLog(action, details string, userID int) error
}
