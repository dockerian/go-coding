// +build all common pkg api error

// Package api :: appError_test.go
package api

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAppError tests func api.AppError
func TestAppError(t *testing.T) {
	statusCode := http.StatusBadRequest
	statusText := http.StatusText(statusCode)
	var err Error = AppError{
		Err:        errors.New(statusText),
		StatusCode: statusCode,
	}

	appErr, ok := err.(AppError)
	msg := "type assersion should convert interface to concrete type AppError"
	assert.NotNil(t, appErr, msg)
	assert.True(t, ok, msg)

	assert.Equal(t, statusCode, appErr.Status())
	assert.Equal(t, statusText, appErr.StatusText())
	assert.Equal(t, statusText, appErr.Error())
}
