package service

import (
	"TestTask/internal/models"
	"TestTask/pkg/utils"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type AuthService struct {
	userService UserRepositoryInterface
}

func NewAuthService(userService UserRepositoryInterface) *AuthService {
	return &AuthService{userService: userService}
}

func (s *AuthService) RegisterUser(user *models.User) error {
	if user.Username == "" {
		return fmt.Errorf("invalid user username")
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	isValidUsername := re.MatchString(user.Username)

	if !isValidUsername {
		return fmt.Errorf("invalid user username: username can only contain letters, digits, and underscores")
	}

	existingUser, _ := s.userService.GetUserByUsername(user.Username)
	if existingUser != nil {
		return fmt.Errorf("user with the same username already exists")
	}

	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = string(hashedPassword)

	return s.userService.CreateUser(user)
}

func (s *AuthService) LoginUser(username, password string) (string, error) {
	if username == "" || password == "" {
		return "", errors.New("username and password cannot be empty")
	}

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
