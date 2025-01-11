package service

import (
	"TestTask/internal/models"
	"TestTask/pkg/utils"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userService UserServiceInterface
}

func NewAuthService(userService UserServiceInterface) *AuthService {
	return &AuthService{userService: userService}
}

func (s *AuthService) RegisterUser(user *models.User) error {
	if user.Username == "" || user.Password == "" {
		return fmt.Errorf("invalid user data")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = string(hashedPassword)

	return s.userService.CreateUser(user)
}

func (s *AuthService) LoginUser(username, password string) (string, error) {
	user, err := s.userService.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
