package models

import "time"

type Log struct {
	ID        int       `json:"id"`
	Action    string    `json:"action"`
	CreatedAt time.Time `json:"created_at"`
	UserID    int       `json:"user_id"`
	Details   string    `json:"details"`
}
