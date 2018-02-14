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
	var err Error = *NewAppError(statusCode, "")

	appErr, ok := err.(AppError)
	msg := "type assersion should convert interface to concrete type AppError"
	assert.NotNil(t, &appErr, msg)
	assert.True(t, ok, msg)

	assert.Equal(t, statusCode, appErr.Status())
	assert.Equal(t, statusText, appErr.StatusText())
	assert.Equal(t, statusText, appErr.Error())

	expected := "missing code and status"
	appError := AppError{Err: errors.New(expected)}
	assert.Equal(t, expected, appError.Error())

	nilError := AppError{StatusCode: http.StatusForbidden}
	codeText := http.StatusText(http.StatusForbidden)
	assert.Equal(t, codeText, nilError.Error())
}
