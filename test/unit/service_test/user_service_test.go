package service_test

import (
	"TestTask/internal/models"
	"TestTask/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// MockUserRepository реализует интерфейс UserRepositoryInterface для тестов
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByID(id int) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByUsername(username string) (*models.User, error) {
	args := m.Called(username)
	return args.Get(0).(*models.User), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := service.NewUserService(mockRepo)

	user := &models.User{
		ID:       1,
		Username: "testuser",
		Password: "password123",
	}

	// Мокаем успешное создание пользователя
	mockRepo.On("CreateUser", user).Return(nil)

	// Тест: успешное создание
	err := userService.CreateUser(user)
	assert.NoError(t, err)

	// Проверка, что мок был вызван
	mockRepo.AssertExpectations(t)
}

func TestGetUserByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := service.NewUserService(mockRepo)

	user := &models.User{
		ID:       1,
		Username: "testuser",
		Password: "password123",
	}

	// Мокаем успешное получение пользователя
	mockRepo.On("GetUserByID", 1).Return(user, nil)

	// Тест: успешное получение пользователя
	result, err := userService.GetUserByID(1)
	assert.NoError(t, err)
	assert.Equal(t, user, result)

	// Мокаем случай, когда пользователь не найден
	mockRepo.On("GetUserByID", 2).Return((*models.User)(nil), nil)

	// Тест: пользователь не найден
	result, err = userService.GetUserByID(2)
	assert.NoError(t, err)
	assert.Nil(t, result)

	// Проверка, что мок был вызван
	mockRepo.AssertExpectations(t)
}

func TestGetUserByUsername(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := service.NewUserService(mockRepo)

	user := &models.User{
		ID:       1,
		Username: "testuser",
		Password: "password123",
	}

	// Мокаем успешное получение пользователя по имени
	mockRepo.On("GetUserByUsername", "testuser").Return(user, nil)

	// Тест: успешное получение пользователя по имени
	result, err := userService.GetUserByUsername("testuser")
	assert.NoError(t, err)
	assert.Equal(t, user, result)

	// Мокаем случай, когда пользователь не найден
	mockRepo.On("GetUserByUsername", "unknown").Return((*models.User)(nil), nil)

	// Тест: пользователь не найден
	result, err = userService.GetUserByUsername("unknown")
	assert.NoError(t, err)
	assert.Nil(t, result)

	// Проверка, что мок был вызван
	mockRepo.AssertExpectations(t)
}
