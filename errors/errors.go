package errors

import (
	"fmt"
	"net/http"
)

type AppErr struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewUnexpectedError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewValidationError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

func NewNotFoundError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewBadRequestError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

func (e AppErr) Error() string {
	return fmt.Sprintf("error has occured with status %v: %v", e.Code, e.Message)
}

func (e AppErr) AsMessage() *AppErr {
	return &AppErr{
		Message: e.Message,
	}
}
