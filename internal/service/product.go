package service

import (
	"TestTask/internal/models"
	"fmt"
)

type ProductService struct {
	repo ProductRepositoryInterface
}

func NewProductService(repo ProductRepositoryInterface) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	if product.Name == "" || product.Price <= 0 {
		return fmt.Errorf("invalid product data")
	}

	return s.repo.CreateProduct(product)
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
	if product.Name == "" || product.Price <= 0 {
		return fmt.Errorf("invalid product data")
	}

	return s.repo.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(productID int) error {
	return s.repo.DeleteProductByID(productID)
}

func (s *ProductService) GetProductByID(productID int) (*models.Product, error) {
	return s.repo.GetProductByID(productID)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAllProducts()
}
