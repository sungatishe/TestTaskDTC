package service

import (
	"TestTask/internal/models"
	"fmt"
	"time"
)

type OrderService struct {
	repo  OrderRepositoryInterface
	cache CacheInterface
}

func NewOrderService(repo OrderRepositoryInterface, cache CacheInterface) *OrderService {
	return &OrderService{repo: repo, cache: cache}
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	if order.CustomerName == "" || order.TotalPrice <= 0 || order.ProductID <= 0 {
		return fmt.Errorf("invalid order data")
	}

	err := s.repo.CreateOrder(order)
	if err != nil {
		return err
	}

	s.cache.SetOrder(order.ID, order)

	return s.repo.CreateOrder(order)
}

func (s *OrderService) UpdateOrder(order *models.Order) error {
	if order.CustomerName == "" || order.TotalPrice <= 0 {
		return fmt.Errorf("invalid order data")
	}

	order.UpdatedAt = time.Now()

	err := s.repo.UpdateOrder(order)
	if err != nil {
		return err
	}

	s.cache.SetOrder(order.ID, order)

	return s.repo.UpdateOrder(order)
}

func (s *OrderService) DeleteOrder(orderID int) error {
	err := s.repo.DeleteOrder(orderID)
	if err != nil {
		return err
	}

	s.cache.DeleteOrder(orderID)
	return nil
}

func (s *OrderService) GetOrderByID(orderID int) (*models.Order, error) {
	cachedOrder, found := s.cache.GetOrder(orderID)
	if found {
		return cachedOrder, nil
	}

	order, err := s.repo.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	s.cache.SetOrder(orderID, order)
	return order, nil
}

func (s *OrderService) GetOrdersByFilters(status string, minPrice, maxPrice float64) ([]models.Order, error) {
	cacheKey := fmt.Sprintf("%s_%f_%f", status, minPrice, maxPrice)

	cachedOrders, found := s.cache.GetOrders(cacheKey)
	if found {
		return cachedOrders, nil
	}

	orders, err := s.repo.GetOrdersByFilters(status, minPrice, maxPrice)
	if err != nil {
		return nil, err
	}

	s.cache.SetOrders(cacheKey, orders)
	return orders, nil
}
