package handlers

import (
	"TestTask/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
)

// AuthData структура для авторизационных данных
type AuthData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterData структура для регистрационных данных
type RegisterData struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type AuthHandler struct {
	service AuthServiceInterface
}

func NewAuthHandlers(service AuthServiceInterface) *AuthHandler {
	return &AuthHandler{service: service}
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Registers a new user by providing user data
// @Tags auth
// @Accept json
// @Produce json
// @Param user body RegisterData true "User registration data"
// @Success 201 {string} string "User successfully registered"
// @Failure 400 {object} ErrorResponse "Invalid user data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /register [post]
func (h *AuthHandler) RegisterUser(rw http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Invalid JSON body: %v", err), http.StatusBadRequest)
		return
	}

	err = h.service.RegisterUser(&user)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Registration failed: %v", err), http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte("User successfully registered"))
}

// LoginUser godoc
// @Summary Log in a user
// @Description Logs in a user and returns a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body AuthData true "User login credentials"
// @Success 200 {object} map[string]string "Token: {token}"
// @Failure 400 {object} ErrorResponse "Invalid login data"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /login [post]
func (h *AuthHandler) LoginUser(rw http.ResponseWriter, r *http.Request) {
	var credentials AuthData

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Invalid JSON body: %v", err), http.StatusBadRequest)
		return
	}

	token, err := h.service.LoginUser(credentials.Username, credentials.Password)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Login failed: %v", err), http.StatusUnauthorized)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]string{
		"token": token,
	})
}
