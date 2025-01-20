package service_test

import (
	"TestTask/internal/cache"
	"TestTask/internal/models"
	"TestTask/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) CreateOrder(order *models.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderRepository) UpdateOrder(order *models.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderRepository) DeleteOrder(orderID int) error {
	args := m.Called(orderID)
	return args.Error(0)
}

func (m *MockOrderRepository) GetOrderByID(orderID int) (*models.Order, error) {
	args := m.Called(orderID)
	if result := args.Get(0); result != nil {
		return result.(*models.Order), args.Error(1)
	}
	return nil, args.Error(1)
}

type MockEventService struct {
	mock.Mock
}

func (m *MockEventService) PublishOrderStatusChanged(orderID int, oldStatus, newStatus string) {
	m.Called(orderID, oldStatus, newStatus)
}

func (m *MockOrderRepository) GetOrdersByFilters(status string, minPrice, maxPrice float64) ([]models.Order, error) {
	args := m.Called(status, minPrice, maxPrice)
	return args.Get(0).([]models.Order), args.Error(1)
}

func TestCreateOrder(t *testing.T) {
	mockRepo := new(MockOrderRepository)
	mockCache := cache.NewCacheService() // Добавляем инстанс CacheService
	mockEventService := new(MockEventService)
	orderService := service.NewOrderService(mockRepo, mockCache, mockEventService) // Передаем cache сюда

	order := &models.Order{
		CustomerName: "John Doe",
		TotalPrice:   99.99,
		ProductID:    1,
	}

	// Мокаем успешное выполнение создания заказа
	mockRepo.On("CreateOrder", order).Return(nil)

	// Тест: успешное создание
	err := orderService.CreateOrder(order)
	assert.NoError(t, err)

	// Мокаем ошибку для invalid данных
	invalidOrder := &models.Order{
		CustomerName: "",
		TotalPrice:   99.99,
		ProductID:    1,
	}
	err = orderService.CreateOrder(invalidOrder)
	assert.Error(t, err)
	assert.Equal(t, "invalid order data", err.Error())

	// Проверка, что мок был вызван
	mockRepo.AssertExpectations(t)
}

func TestUpdateOrder(t *testing.T) {
	mockRepo := new(MockOrderRepository)
	mockCache := cache.NewCacheService()
	mockEventService := new(MockEventService) // Используем MockEventService

	orderService := service.NewOrderService(mockRepo, mockCache, mockEventService)

	existingOrder := &models.Order{
		ID:           1,
		CustomerName: "John Doe",
		TotalPrice:   99.99,
		Status:       "pending",
		ProductID:    1,
	}

	updatedOrder := &models.Order{
		ID:           1,
		CustomerName: "John Doe Updated",
		TotalPrice:   199.99,
		Status:       "completed",
		ProductID:    1,
	}

	mockRepo.On("GetOrderByID", existingOrder.ID).Return(existingOrder, nil)
	mockRepo.On("UpdateOrder", updatedOrder).Return(nil)
	mockEventService.On("PublishOrderStatusChanged", updatedOrder.ID, "pending", "completed").Return()

	err := orderService.UpdateOrder(updatedOrder)
	assert.NoError(t, err)
	mockEventService.AssertCalled(t, "PublishOrderStatusChanged", updatedOrder.ID, "pending", "completed")
	mockRepo.AssertExpectations(t)
}

func TestDeleteOrder(t *testing.T) {
	mockRepo := new(MockOrderRepository)
	mockCache := cache.NewCacheService() // Добавляем инстанс CacheService
	mockEventService := new(MockEventService)
	orderService := service.NewOrderService(mockRepo, mockCache, mockEventService) // Передаем cache сюда

	// Мокаем успешное выполнение удаления
	mockRepo.On("DeleteOrder", 1).Return(nil)

	// Тест: успешное удаление
	err := orderService.DeleteOrder(1)
	assert.NoError(t, err)

	// Проверка, что мок был вызван
	mockRepo.AssertExpectations(t)
}

func TestGetOrderByID(t *testing.T) {
	mockRepo := new(MockOrderRepository)
	mockCache := cache.NewCacheService() // Добавляем инстанс CacheService
	mockEventService := new(MockEventService)
	orderService := service.NewOrderService(mockRepo, mockCache, mockEventService) // Передаем cache сюда

	order := &models.Order{
		ID:           1,
		CustomerName: "John Doe",
		TotalPrice:   99.99,
		ProductID:    1,
	}

	// Мокаем успешное получение заказа
	mockRepo.On("GetOrderByID", 1).Return(order, nil)

	// Тест: успешное получение
	result, err := orderService.GetOrderByID(1)
	assert.NoError(t, err)
	assert.Equal(t, order, result)

	// Тест: ошибка для несуществующего заказа
	mockRepo.On("GetOrderByID", 2).Return(nil, nil)
	result, err = orderService.GetOrderByID(2)
	assert.NoError(t, err)
	assert.Nil(t, result)

	// Проверка, что мок был вызван
	mockRepo.AssertExpectations(t)
}

func TestGetOrdersByFilters(t *testing.T) {
	mockRepo := new(MockOrderRepository)
	mockCache := cache.NewCacheService() // Добавляем инстанс CacheService
	mockEventService := new(MockEventService)
	orderService := service.NewOrderService(mockRepo, mockCache, mockEventService) // Передаем cache сюда

	orders := []models.Order{
		{ID: 1, CustomerName: "John Doe", TotalPrice: 99.99, ProductID: 1},
		{ID: 2, CustomerName: "Jane Doe", TotalPrice: 149.99, ProductID: 2},
	}

	// Мокаем успешное выполнение фильтрации заказов
	mockRepo.On("GetOrdersByFilters", "pending", float64(0), float64(200)).Return(orders, nil)

	// Тест: успешное получение заказов
	result, err := orderService.GetOrdersByFilters("pending", 0, 200)
	assert.NoError(t, err)
	assert.Equal(t, orders, result)

	// Проверка, что мок был вызван
	mockRepo.AssertExpectations(t)
}
