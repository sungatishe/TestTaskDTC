package service

import (
	"TestTask/internal/models"
	"fmt"
)

type LogService struct {
	repo LogRepository
}

func NewLogService(repo LogRepository) *LogService {
	return &LogService{repo: repo}
}

func (s *LogService) CreateLog(action, details string, userID int) error {
	log := &models.Log{
		Action:  action,
		UserID:  userID,
		Details: details,
	}

	err := s.repo.CreateLog(log)
	if err != nil {
		return fmt.Errorf("failed to create log: %w", err)
	}
	return nil
}
