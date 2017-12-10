// Package api :: formatter.go - api formatters
package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

// GetJSONEncoder returns JSON by specified indent
func GetJSONEncoder(w io.Writer, indent string) *json.Encoder {
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", indent)
	return encoder
}

// WriteError writes status code and returns Error
func WriteError(w http.ResponseWriter, code int, message string) Error {
	appError := AppError{
		Err:        errors.New(message),
		StatusCode: code,
	}
	return WriteAppError(w, &appError)
}

// WriteAppError writes status code and returns Error
func WriteAppError(w http.ResponseWriter, apiError Error) Error {
	code := apiError.Status()
	errorMessage := apiError.Error()
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	log.Printf("[status] %d: %s\n", code, errorMessage)
	encoder := GetJSONEncoder(w, "  ")
	encoder.Encode(apiError)
	return apiError
}

// WriteJSON writes status code and response data
func WriteJSON(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	encoder := GetJSONEncoder(w, "  ")
	encoder.Encode(data)
}
