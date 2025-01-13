package handlers

import (
	"TestTask/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthHandler struct {
	service AuthServiceInterface
}

func NewAuthHandlers(service AuthServiceInterface) *AuthHandler {
	return &AuthHandler{service: service}
}

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

func (h *AuthHandler) LoginUser(rw http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

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
