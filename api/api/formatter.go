// Package api :: formatter.go - api formatters
package api

import (
	"encoding/json"
	"fmt"
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

// WriteError writes status code and ApiError
func WriteError(w http.ResponseWriter, code int, apiInfo, message, trace string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	log.Printf("[status] %d: %s - %s\n", code, apiInfo, message)
	apiError := AppError{
		Err:        fmt.Errorf("%d | %s | %s | %s", code, apiInfo, message, trace),
		StatusCode: code,
	}
	encoder := GetJSONEncoder(w, "  ")
	encoder.Encode(apiError)
}

// WriteJSON writes status code and response data
func WriteJSON(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	encoder := GetJSONEncoder(w, "  ")
	encoder.Encode(data)
}
