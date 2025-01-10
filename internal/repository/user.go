package repository

import (
	"TestTask/internal/models"
	"database/sql"
	"errors"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users
		(username, password, role, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
	`
	_, err := r.db.Exec(query, user.Username, user.Password, user.Role)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	query := "SELECT id, username, password, role, created_at, updated_at FROM users WHERE username = $1"
	row := r.db.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil // No user found
	} else if err != nil {
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	query := "SELECT id, username, password, role, created_at, updated_at FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil // No user found
	} else if err != nil {
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}

	return &user, nil
}
