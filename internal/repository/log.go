package repository

import (
	"TestTask/internal/models"
	"database/sql"
	"fmt"
)

type LogRepository struct {
	db *sql.DB
}

func NewLogRepository(db *sql.DB) *LogRepository {
	return &LogRepository{db: db}
}

func (r *LogRepository) CreateLog(log *models.Log) error {
	query := `INSERT INTO logs (action, user_id, details) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, log.Action, log.UserID, log.Details)
	if err != nil {
		return fmt.Errorf("failed to create log: %w", err)
	}
	return nil
}
