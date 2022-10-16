package middleware

import "encoding/json"

type CustomError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewCustomError(status int, message string) *CustomError {
	return &CustomError{
		Status:  status,
		Message: message,
	}
}

func (ce *CustomError) Error() string {
	return ce.Message
}

func (ce *CustomError) GetStatus() int {
	return ce.Status
}

func (ce *CustomError) GetMessage() string {
	b, _ := json.Marshal(ce)
	return string(b)
}
