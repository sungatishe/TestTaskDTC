package handlers

import (
	"TestTask/internal/models"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"strings"
)

type OrderHandler struct {
	service    OrderServiceInterface
	logService LogServiceInterface
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewOrderHandler(service OrderServiceInterface, logService LogServiceInterface) *OrderHandler {
	return &OrderHandler{service: service, logService: logService}
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order by providing order data
// @Tags orders
// @Accept json
// @Produce json
// @Param order body models.Order true "Order data"
// @Success 201 {object} models.Order
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Roles User, Admin
// @Router /orders [post]
func (h *OrderHandler) CreateOrder(rw http.ResponseWriter, r *http.Request) {
	var order models.Order

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&order)
	if err != nil {
		http.Error(rw, "Invalid input data", http.StatusBadRequest)
		return
	}

	err = h.service.CreateOrder(&order)
	if err != nil {
		if strings.Contains(err.Error(), "invalid order data") {
			http.Error(rw, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	cookie, err := r.Cookie("user_id")
	if err != nil || cookie == nil {
		http.Error(rw, "User ID cookie not found", http.StatusUnauthorized)
		return
	}

	userID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		http.Error(rw, "Invalid User ID", http.StatusUnauthorized)
		return
	}
	action := "create_order"
	details := "Order created successfully"

	err = h.logService.CreateLog(action, details, userID)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(order)
}

// UpdateOrder godoc
// @Summary Update an existing order
// @Description Update an existing order by providing order data
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param order body models.Order true "Updated order data"
// @Success 200 {object} models.Order "Order updated successfully"
// @Failure 400 {object} ErrorResponse "Invalid order ID or data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security ApiKeyAuth
// @Roles User, Admin
// @Router /orders/{id} [put]
func (h *OrderHandler) UpdateOrder(rw http.ResponseWriter, r *http.Request) {
	orderIDStr := chi.URLParam(r, "id")
	if orderIDStr == "" {
		http.Error(rw, "Missing order ID", http.StatusBadRequest)
		return
	}

	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		http.Error(rw, "Invalid order ID", http.StatusBadRequest)
		return
	}

	var order models.Order

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&order)
	if err != nil {
		http.Error(rw, "Invalid input data", http.StatusBadRequest)
		return
	}
	order.ID = orderID

	err = h.service.UpdateOrder(&order)
	if err != nil {
		if strings.Contains(err.Error(), "invalid order data") {
			http.Error(rw, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	cookie, err := r.Cookie("user_id")
	if err != nil || cookie == nil {
		http.Error(rw, "User ID cookie not found", http.StatusUnauthorized)
		return
	}

	userID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		http.Error(rw, "Invalid User ID", http.StatusUnauthorized)
		return
	}
	action := "update_order"
	details := "Order updated successfully"

	err = h.logService.CreateLog(action, details, userID)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(order)
}

// DeleteOrder godoc
// @Summary Delete an order
// @Description Delete an order by providing order ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 204 "Order deleted successfully"
// @Failure 400 {object} ErrorResponse "Invalid order ID"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security ApiKeyAuth
// @Roles Admin
// @Router /orders/{id} [delete]
func (h *OrderHandler) DeleteOrder(rw http.ResponseWriter, r *http.Request) {
	orderIDStr := chi.URLParam(r, "id")
	if orderIDStr == "" {
		http.Error(rw, "Missing order ID", http.StatusBadRequest)
		return
	}

	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		http.Error(rw, "Invalid order ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteOrder(orderID)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie, err := r.Cookie("user_id")
	if err != nil || cookie == nil {
		http.Error(rw, "User ID cookie not found", http.StatusUnauthorized)
		return
	}

	userID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		http.Error(rw, "Invalid User ID", http.StatusUnauthorized)
		return
	}
	action := "delete_order"
	details := "Order deleted successfully"

	err = h.logService.CreateLog(action, details, userID)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

// GetOrderByID godoc
// @Summary Get an order by ID
// @Description Get a specific order by providing the order ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order "Order details"
// @Failure 400 {object} ErrorResponse "Invalid order ID"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security ApiKeyAuth
// @Roles User, Admin
// @Router /orders/{id} [get]
func (h *OrderHandler) GetOrderByID(rw http.ResponseWriter, r *http.Request) {
	orderIDStr := chi.URLParam(r, "id")
	if orderIDStr == "" {
		http.Error(rw, "Missing order ID", http.StatusBadRequest)
		return
	}

	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		http.Error(rw, "Invalid order ID", http.StatusBadRequest)
		return
	}

	order, err := h.service.GetOrderByID(orderID)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(order)
}

// GetOrdersByFilters godoc
// @Summary Get orders by filters
// @Description Get all orders or filter them by status, min price, and max price
// @Tags orders
// @Accept json
// @Produce json
// @Param status query string false "Order status"
// @Param min_price query float64 false "Minimum order price"
// @Param max_price query float64 false "Maximum order price"
// @Success 200 {array} models.Order "List of orders"
// @Failure 400 {object} ErrorResponse "Invalid filter parameters"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security ApiKeyAuth
// @Router /orders [get]
func (h *OrderHandler) GetOrdersByFilters(rw http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	minPriceStr := r.URL.Query().Get("min_price")
	maxPriceStr := r.URL.Query().Get("max_price")

	var minPrice, maxPrice float64
	var err error

	if minPriceStr != "" {
		minPrice, err = strconv.ParseFloat(minPriceStr, 64)
		if err != nil {
			http.Error(rw, "Invalid min_price parameter", http.StatusBadRequest)
			return
		}
	}

	if maxPriceStr != "" {
		maxPrice, err = strconv.ParseFloat(maxPriceStr, 64)
		if err != nil {
			http.Error(rw, "Invalid max_price parameter", http.StatusBadRequest)
			return
		}
	}

	orders, err := h.service.GetOrdersByFilters(status, minPrice, maxPrice)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(orders)
}
