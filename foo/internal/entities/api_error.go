package entities

import "time"

type APIError struct {
	Message   string `json:"message"`
	TimeStamp string `json:"timeStamp"`
	Details   any    `json:"details"`
}

func (e *APIError) Error() string { return e.Message }

func NewAPIError(message string) *APIError {
	return &APIError{Message: message, TimeStamp: time.Now().String()}
}
