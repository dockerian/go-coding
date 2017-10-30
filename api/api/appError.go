// Package api :: appError.go
// api.Error interface wraps error with http status.
// api.AppError composes error and http status code for http handler without
// accessing to header in http.ResponseWriter.
package api

import (
	"errors"
	"net/http"
)

var (
	// EncodeError is for json.NewEncoder().Encode() failure
	EncodeError = AppError{
		errors.New("encoder error"),
		http.StatusInternalServerError,
	}
)

// Error represents a handler error to provide
// Status() and embed the built-in error interface.
type Error interface {
	error
	Status() int
}

// AppError represents an error with an associated HTTP status code.
type AppError struct {
	Err        error
	StatusCode int
}

// Error for AppError to implement the error interface.
func (ae AppError) Error() string {
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
