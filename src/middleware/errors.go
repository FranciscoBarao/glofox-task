package middleware

import "strconv"

type CustomError struct {
	status  int
	message string
}

func NewCustomError(status int, message string) *CustomError {
	return &CustomError{
		status:  status,
		message: message,
	}
}

func (ce *CustomError) Error() string {
	return ce.message
}

func (ce *CustomError) GetStatus() int {
	return ce.status
}

func (ce *CustomError) GetMessage() string {
	return strconv.Itoa(ce.status) + " - " + ce.message
}
