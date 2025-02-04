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

type CacheInterface interface {
	SetOrder(orderID int, order *models.Order)
	GetOrder(orderID int) (*models.Order, bool)
	DeleteOrder(orderID int)
	SetOrders(key string, orders []models.Order)
	GetOrders(key string) ([]models.Order, bool)
}

type LogRepository interface {
	CreateLog(log *models.Log) error
}

type ProducerInterface interface {
	Publish(key, value []byte) error
	Close() error
}

type EventServiceInterface interface {
	PublishOrderStatusChanged(orderID int, oldStatus, newStatus string)
}
