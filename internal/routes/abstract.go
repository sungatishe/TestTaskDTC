package routes

import (
	"net/http"
)

// OrderHandlerInterface определяет методы для управления заказами.
type OrderHandlerInterface interface {
	GetOrdersByFilters(w http.ResponseWriter, r *http.Request)
	GetOrderByID(w http.ResponseWriter, r *http.Request)
	CreateOrder(w http.ResponseWriter, r *http.Request)
	UpdateOrder(w http.ResponseWriter, r *http.Request)
	DeleteOrder(w http.ResponseWriter, r *http.Request)
}

// ProductHandlerInterface определяет методы для управления продуктами.
type ProductHandlerInterface interface {
	GetAllProducts(w http.ResponseWriter, r *http.Request)
	GetProductByID(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
}

// AuthHandlerInterface определяет методы для управления аутентификацией.
type AuthHandlerInterface interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
}
