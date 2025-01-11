package service

import "TestTask/internal/models"

type UserRepositoryInterface interface {
	CreateUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
}

type OrderRepositoryInterface interface {
	CreateOrder(order *models.Order) error
	UpdateOrder(order *models.Order) error
	DeleteOrder(orderID int) error
	GetOrderByID(orderID int) (*models.Order, error)
	GetOrdersByFilters(status string, minPrice, maxPrice float64) ([]models.Order, error)
}

type ProductRepositoryInterface interface {
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProductByID(productID int) error
	GetProductByID(productID int) (*models.Product, error)
	GetAllProducts() ([]models.Product, error)
}
