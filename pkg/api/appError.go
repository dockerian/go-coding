// Package api :: appError.go
//
// api.Error interface wraps error with http status.
// api.AppError composes error and http status code for http handler without
// accessing to header in http.ResponseWriter.
package api

import (
	"errors"
	"net/http"
)

// Error represents a handler error to provide
// Status() and embed the built-in error interface.
type Error interface {
	error
	Status() int
}

// AppError represents an error with an associated HTTP status code.
type AppError struct {
	// Err inherits standard error interface
	Err error `json:"-"`
	// ErrorMessage represents Err.Error()
	ErrorMessage string `json:"message,omitempty"`
	// StatusCode is http status code
	StatusCode int `json:"code,omitempty"`
}

// Error for AppError to implement the error interface.
func (ae AppError) Error() string {
	if ae.ErrorMessage != "" {
		return ae.ErrorMessage
	} else if ae.Err == nil {
		return http.StatusText(ae.StatusCode)
	}
	return ae.Err.Error()
}

// Status returns HTTP status code.
func (ae AppError) Status() int {
	return ae.StatusCode
}

// StatusText returns HTTP status text.
func (ae AppError) StatusText() string {
	return http.StatusText(ae.StatusCode)
}

// NewAppError constructs an AppError
func NewAppError(code int, message string) *AppError {
	if message == "" {
		message = http.StatusText(code)
	}
	return &AppError{
		Err:          errors.New(message),
		ErrorMessage: message,
		StatusCode:   code,
	}
}
