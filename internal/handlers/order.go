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

func NewOrderHandler(service OrderServiceInterface, logService LogServiceInterface) *OrderHandler {
	return &OrderHandler{service: service, logService: logService}
}

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

func (h *OrderHandler) GetOrdersByFilters(rw http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	minPriceStr := r.URL.Query().Get("min_price")
	maxPriceStr := r.URL.Query().Get("max_price")

	if status == "" || minPriceStr == "" || maxPriceStr == "" {
		http.Error(rw, "Missing query parameters", http.StatusBadRequest)
		return
	}

	minPrice, err := strconv.ParseFloat(minPriceStr, 64)
	if err != nil {
		http.Error(rw, "Invalid min_price parameter", http.StatusBadRequest)
		return
	}

	maxPrice, err := strconv.ParseFloat(maxPriceStr, 64)
	if err != nil {
		http.Error(rw, "Invalid max_price parameter", http.StatusBadRequest)
		return
	}

	// Запрашиваем заказы с фильтрацией через сервис
	orders, err := h.service.GetOrdersByFilters(status, minPrice, maxPrice)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем результат
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(orders)
}
