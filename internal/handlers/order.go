package handlers

import (
	"TestTask/internal/models"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	service OrderServiceInterface
}

func NewOrderHandler(service OrderServiceInterface) *OrderHandler {
	return &OrderHandler{service: service}
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
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(order)
}

func (h *OrderHandler) UpdateOrder(rw http.ResponseWriter, r *http.Request) {
	orderIDStr := r.URL.Query().Get("id")
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
