package api

import (
	"log"
	"net/http"
	"time"
)

var (
	// EncodeError is for json.NewEncoder().Encode() failure
	EncodeError = Error{
		"Encoder error",
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
	}
)

// Error struct
type Error struct {
	Err       string `json:"error,omitempty"`
	ErrCode   int    `json:"error_code,omitempty"`
	ErrStatus string `json:"status,omitempty"`
}

// Logger func
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
