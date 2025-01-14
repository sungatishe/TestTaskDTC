package models

// ErrorResponse структура для ошибки
type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
