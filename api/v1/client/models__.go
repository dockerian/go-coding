// Package client :: models__.go - extended model functions
package client

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dockerian/go-coding/pkg/api"
)

//////// -- ApiError model functions

// NewApiError constructs an ApiError
//
// Note: ApiError implements both standard error and api.Error interfaces
// See github.com/dockerian/go-coding/pkg/api
func NewApiError(code int, apiInfo, message, log string) *ApiError {
	return &ApiError{
		ApiInfo:    apiInfo,
		Code:       int32(code),
		CodeStatus: http.StatusText(code),
		Message:    message,
		Log:        log,
	}
}

// Error implements error interface
func (apiError *ApiError) Error() string {
	return fmt.Sprintf("%s - %s", apiError.ApiInfo, apiError.Message)
}

// Status implements api.Error interface (see pkg/cli/appError.go)
func (apiError *ApiError) Status() int {
	return int(apiError.Code)
}

// WriteError writes status code and returns api.Error
func (apiError *ApiError) WriteError(w http.ResponseWriter) api.Error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(int(apiError.Code))
	log.Printf("[AppError] %d: %s - %s | %s\n",
		apiError.Code, apiError.ApiInfo, apiError.Message, apiError.Log)
	encoder := api.GetJSONEncoder(w, "  ")
	encoder.Encode(*apiError)
	return apiError
}

// WriteError writes status code and returns Error
func WriteError(w http.ResponseWriter, code int, apiInfo, message, trace string) api.Error {
	apiError := NewApiError(code, apiInfo, message, trace)
	return apiError.WriteError(w)
}
