package repository

import (
	"TestTask/internal/models"
	"TestTask/internal/repository"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("could not create mock database: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)

	user := &models.User{
		Username: "john_doe",
		Password: "password123",
		Role:     "admin",
	}

	mock.ExpectExec(`INSERT INTO users`).
		WithArgs(user.Username, user.Password, user.Role).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = userRepo.CreateUser(user)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %v", err)
	}
}

func TestGetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("could not create mock database: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)

	userID := 1
	user := &models.User{
		ID:        userID,
		Username:  "john_doe",
		Password:  "password123",
		Role:      "admin",
		CreatedAt: "",
		UpdatedAt: "",
	}

	mock.ExpectQuery(`SELECT id, username, password, role, created_at, updated_at`).
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "role", "created_at", "updated_at"}).
			AddRow(user.ID, user.Username, user.Password, user.Role, user.CreatedAt, user.UpdatedAt))

	result, err := userRepo.GetUserByID(userID)
	assert.NoError(t, err)

	user.CreatedAt = result.CreatedAt
	user.UpdatedAt = result.UpdatedAt

	assert.Equal(t, user, result)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %v", err)
	}
}
