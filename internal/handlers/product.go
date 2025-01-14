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

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product by providing product data
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product data"
// @Success 201 {string} string "Product successfully created"
// @Failure 400 {object} ErrorResponse "Invalid product data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security ApiKeyAuth
// @Roles Admin
// @Router /products [post]
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

// UpdateProduct godoc
// @Summary Update an existing product
// @Description Update an existing product by providing product data
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.Product true "Updated product data"
// @Success 200 {string} string "Product successfully updated"
// @Failure 400 {object} ErrorResponse "Invalid product ID or data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security ApiKeyAuth
// @Roles Admin
// @Router /products/{id} [put]
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

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product by providing product ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {string} string "Product successfully deleted"
// @Failure 400 {object} ErrorResponse "Invalid product ID"
// @Failure 404 {object} ErrorResponse "Product not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security ApiKeyAuth
// @Roles Admin
// @Router /products/{id} [delete]
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

// GetProductByID godoc
// @Summary Get a product by ID
// @Description Get a specific product by providing the product ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product "Product details"
// @Failure 400 {object} ErrorResponse "Invalid product ID"
// @Failure 404 {object} ErrorResponse "Product not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security ApiKeyAuth
// @Roles User, Admin
// @Router /products/{id} [get]
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

// GetAllProducts godoc
// @Summary Get all products
// @Description Retrieve a list of all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} models.Product "List of all products"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Security ApiKeyAuth
// @Roles User, Admin
// @Router /products [get]
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
