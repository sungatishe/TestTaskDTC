package repository

import (
	"TestTask/internal/models"
	"TestTask/internal/repository"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("could not create mock database: %v", err)
	}
	defer db.Close()

	productRepo := repository.NewProductRepository(db)

	product := &models.Product{
		Name:     "Product",
		Price:    99.99,
		Quantity: 50,
	}

	mock.ExpectExec(`INSERT INTO products`).
		WithArgs(product.Name, product.Price, product.Quantity).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = productRepo.CreateProduct(product)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %v", err)
	}
}

func TestUpdateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("could not create mock database: %v", err)
	}
	defer db.Close()

	productRepo := repository.NewProductRepository(db)

	product := &models.Product{
		ID:       1,
		Name:     "Updated Product",
		Price:    15.99,
		Quantity: 5,
	}

	mock.ExpectExec(`UPDATE products`).
		WithArgs(product.Name, product.Price, product.Quantity, product.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = productRepo.UpdateProduct(product)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %v", err)
	}
}

func TestGetProductByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("could not create mock database: %v", err)
	}
	defer db.Close()

	productRepo := repository.NewProductRepository(db)

	productID := 1
	expectedProduct := &models.Product{
		ID:       1,
		Name:     "Test Product",
		Price:    12.99,
		Quantity: 10,
	}

	mock.ExpectQuery(`SELECT id, name, price, quantity`).
		WithArgs(productID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price", "quantity"}).
			AddRow(expectedProduct.ID, expectedProduct.Name, expectedProduct.Price, expectedProduct.Quantity))

	result, err := productRepo.GetProductByID(productID)
	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, result)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %v", err)
	}
}

func TestGetAllProducts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("could not create mock database: %v", err)
	}
	defer db.Close()

	productRepo := repository.NewProductRepository(db)

	expectedProducts := []models.Product{
		{ID: 1, Name: "Product 1", Price: 10.99, Quantity: 5},
		{ID: 2, Name: "Product 2", Price: 20.99, Quantity: 2},
	}

	mock.ExpectQuery(`SELECT id, name, price, quantity`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price", "quantity"}).
			AddRow(expectedProducts[0].ID, expectedProducts[0].Name, expectedProducts[0].Price, expectedProducts[0].Quantity).
			AddRow(expectedProducts[1].ID, expectedProducts[1].Name, expectedProducts[1].Price, expectedProducts[1].Quantity))

	result, err := productRepo.GetAllProducts()
	assert.NoError(t, err)
	assert.Equal(t, expectedProducts, result)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %v", err)
	}
}

func TestDeleteProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("could not create mock database: %v", err)
	}
	defer db.Close()

	productRepo := repository.NewProductRepository(db)

	productID := 1

	mock.ExpectExec(`DELETE FROM`).
		WithArgs(productID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = productRepo.DeleteProductByID(productID)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %v", err)
	}
}
