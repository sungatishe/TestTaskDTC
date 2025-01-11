package service

import (
	"TestTask/internal/models"
	"fmt"
	"time"
)

type OrderService struct {
	repo OrderRepositoryInterface
}

func NewOrderService(repo OrderRepositoryInterface) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	if order.CustomerName == "" || order.TotalPrice <= 0 || order.ProductID <= 0 {
		return fmt.Errorf("invalid order data")
	}

	return s.repo.CreateOrder(order)
}

func (s *OrderService) UpdateOrder(order *models.Order) error {
	if order.CustomerName == "" || order.TotalPrice <= 0 {
		return fmt.Errorf("invalid order data")
	}

	order.UpdatedAt = time.Now()
	return s.repo.UpdateOrder(order)
}

func (s *OrderService) DeleteOrder(orderID int) error {
	return s.repo.DeleteOrder(orderID)
}

func (s *OrderService) GetOrderByID(orderID int) (*models.Order, error) {
	return s.repo.GetOrderByID(orderID)
}

func (s *OrderService) GetOrdersByFilters(status string, minPrice, maxPrice float64) ([]models.Order, error) {
	return s.repo.GetOrdersByFilters(status, minPrice, maxPrice)
}
