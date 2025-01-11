package handlers

import (
	"TestTask/internal/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	service ProductServiceInterface
}

func NewProductHandler(service ProductServiceInterface) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(rw http.ResponseWriter, r *http.Request) {
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Invalid JSON body: %v", err), http.StatusBadRequest)
		return
	}

	err = h.service.CreateProduct(&product)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Product creation failed: %v", err), http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte("Product successfully created"))
}

func (h *ProductHandler) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	productIDStr := chi.URLParam(r, "id")

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		http.Error(rw, "Invalid order ID", http.StatusBadRequest)
		return
	}

	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Invalid JSON body: %v", err), http.StatusBadRequest)
		return
	}

	product.ID = productID

	err = h.service.UpdateProduct(&product)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Product update failed: %v", err), http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Product successfully updated"))
}

func (h *ProductHandler) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	productIDStr := chi.URLParam(r, "id")

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		http.Error(rw, "Invalid order ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteProduct(productID)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Product deletion failed: %v", err), http.StatusNotFound)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Product successfully deleted"))
}

func (h *ProductHandler) GetProductByID(rw http.ResponseWriter, r *http.Request) {
	productIDStr := chi.URLParam(r, "id")

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		http.Error(rw, "Invalid order ID", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetProductByID(productID)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Product not found: %v", err), http.StatusNotFound)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(product)
}

func (h *ProductHandler) GetAllProducts(rw http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		http.Error(rw, fmt.Sprintf("Failed to retrieve products: %v", err), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(products)
}
