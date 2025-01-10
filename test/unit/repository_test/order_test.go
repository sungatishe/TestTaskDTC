package repository_test

import (
	"TestTask/internal/models"
	"TestTask/internal/repository"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("could not create mock database: %v", err)
	}
	defer db.Close()

	orderRepo := repository.NewOrderRepository(db)

	order := &models.Order{
		CustomerName: "John Doe",
		Status:       "pending",
		TotalPrice:   99.99,
		ProductID:    1,
	}

	mock.ExpectExec(`INSERT INTO orders`).
		WithArgs(order.CustomerName, order.Status, order.TotalPrice, order.ProductID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = orderRepo.CreateOrder(order)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %v", err)
	}

}

func TestUpdateOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("could not create mock database: %v", err)
	}
	defer db.Close()

	orderRepo := repository.NewOrderRepository(db)

	order := &models.Order{
		ID:           1,
		CustomerName: "Jane Doe",
		Status:       "confirmed",
		TotalPrice:   119.99,
		UpdatedAt:    time.Now(),
	}

	mock.ExpectExec(`UPDATE orders`).
		WithArgs(order.CustomerName, order.Status, order.TotalPrice, order.UpdatedAt, order.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = orderRepo.UpdateOrder(order)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %v", err)
	}
}

func TestDeleteOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("could not create mock database: %v", err)
	}
	defer db.Close()

	orderRepo := repository.NewOrderRepository(db)

	orderID := 1

	mock.ExpectExec(`UPDATE orders`).
		WithArgs(orderID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = orderRepo.DeleteOrder(orderID)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %v", err)
	}
}

func TestGetOrderByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("could not create mock database: %v", err)
	}
	defer db.Close()

	orderRepo := repository.NewOrderRepository(db)

	orderID := 1
	order := &models.Order{
		ID:           1,
		CustomerName: "John Doe",
		Status:       "pending",
		TotalPrice:   99.99,
		ProductID:    1,
		IsDeleted:    false,
	}

	mock.ExpectQuery(`SELECT`).
		WithArgs(orderID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "customer_name", "status", "total_price", "product_id", "created_at", "updated_at", "is_deleted"}).
			AddRow(order.ID, order.CustomerName, order.Status, order.TotalPrice, order.ProductID, time.Now(), time.Now(), order.IsDeleted))

	result, err := orderRepo.GetOrderByID(orderID)
	assert.NoError(t, err)
	assert.Equal(t, order.ID, result.ID)
	assert.Equal(t, order.CustomerName, result.CustomerName)
	assert.Equal(t, order.Status, result.Status)
	assert.Equal(t, order.TotalPrice, result.TotalPrice)
	assert.Equal(t, order.ProductID, result.ProductID)
	assert.Equal(t, order.IsDeleted, result.IsDeleted)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %v", err)
	}
}
