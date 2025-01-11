package service_test

import (
	"TestTask/internal/models"
	"TestTask/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// MockProductRepository для мока ProductRepositoryInterface
type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) CreateProduct(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) UpdateProduct(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) DeleteProductByID(productID int) error {
	args := m.Called(productID)
	return args.Error(0)
}

func (m *MockProductRepository) GetProductByID(productID int) (*models.Product, error) {
	args := m.Called(productID)
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductRepository) GetAllProducts() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func TestCreateProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	productService := service.NewProductService(mockRepo)

	product := &models.Product{
		Name:  "Product A",
		Price: 99.99,
	}

	// Мокаем успешное выполнение создания продукта
	mockRepo.On("CreateProduct", product).Return(nil)

	// Тест: успешное создание
	err := productService.CreateProduct(product)
	assert.NoError(t, err)

	// Тест: ошибка для невалидных данных
	invalidProduct := &models.Product{
		Name:  "",
		Price: -10.00,
	}
	err = productService.CreateProduct(invalidProduct)
	assert.Error(t, err)
	assert.Equal(t, "invalid product data", err.Error())

	// Проверяем вызов мока
	mockRepo.AssertExpectations(t)
}

func TestUpdateProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	productService := service.NewProductService(mockRepo)

	product := &models.Product{
		ID:    1,
		Name:  "Product A",
		Price: 99.99,
	}

	// Мокаем успешное выполнение обновления
	mockRepo.On("UpdateProduct", product).Return(nil)

	// Тест: успешное обновление
	err := productService.UpdateProduct(product)
	assert.NoError(t, err)

	// Тест: ошибка для невалидных данных
	invalidProduct := &models.Product{
		ID:    1,
		Name:  "",
		Price: -20.00,
	}
	err = productService.UpdateProduct(invalidProduct)
	assert.Error(t, err)
	assert.Equal(t, "invalid product data", err.Error())

	// Проверяем вызов мока
	mockRepo.AssertExpectations(t)
}

func TestDeleteProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	productService := service.NewProductService(mockRepo)

	// Мокаем успешное удаление
	mockRepo.On("DeleteProductByID", 1).Return(nil)

	// Тест: успешное удаление
	err := productService.DeleteProduct(1)
	assert.NoError(t, err)

	// Проверяем вызов мока
	mockRepo.AssertExpectations(t)
}

func TestGetProductByID(t *testing.T) {
	mockRepo := new(MockProductRepository)
	productService := service.NewProductService(mockRepo)

	product := &models.Product{
		ID:    1,
		Name:  "Product A",
		Price: 99.99,
	}

	// Мокаем успешное получение продукта
	mockRepo.On("GetProductByID", 1).Return(product, nil)

	// Тест: успешное получение
	result, err := productService.GetProductByID(1)
	assert.NoError(t, err)
	assert.Equal(t, product, result)

	// Мокаем случай, когда продукт не найден (возвращаем nil как *models.Product)
	mockRepo.On("GetProductByID", 2).Return((*models.Product)(nil), nil)

	// Тест: продукт не найден
	result, err = productService.GetProductByID(2)
	assert.NoError(t, err)
	assert.Nil(t, result)

	// Проверяем вызов мока
	mockRepo.AssertExpectations(t)
}

func TestGetAllProducts(t *testing.T) {
	mockRepo := new(MockProductRepository)
	productService := service.NewProductService(mockRepo)

	products := []models.Product{
		{ID: 1, Name: "Product A", Price: 99.99},
		{ID: 2, Name: "Product B", Price: 149.99},
	}

	// Мокаем успешное выполнение получения всех продуктов
	mockRepo.On("GetAllProducts").Return(products, nil)

	// Тест: успешное получение всех продуктов
	result, err := productService.GetAllProducts()
	assert.NoError(t, err)
	assert.Equal(t, products, result)

	// Проверяем вызов мока
	mockRepo.AssertExpectations(t)
}
