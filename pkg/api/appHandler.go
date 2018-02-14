// Package api :: appHandler.go
//
// api.AppHandler declares an extended http.Handler with configuration data
// and error (see api.AppError).
package api

import (
	"log"
	"net/http"

	"github.com/dockerian/go-coding/pkg/cfg"
)

// HandlerFunc represents a handler function with cfg.Context
type HandlerFunc func(ctx cfg.Context, w http.ResponseWriter, r *http.Request) error

// AppHandler struct wraps Env and Handle function implementing http.Handler
type AppHandler struct {
	cfg.Context
	Handle func(e cfg.Context, w http.ResponseWriter, r *http.Request) error
}

// ServeHTTP implements http.Handler
func (ah AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if ah.Handle == nil {
		message := "missing handler func in AppHandler"
		log.Printf("[AppHandler] %s\n", message)
		WriteError(w, http.StatusBadRequest, message)
		return
	}
	err := ah.Handle(ah.Context, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// retrieve the status and print out
			log.Printf("[AppHandler] HTTP %d - %v", e.Status(), e.Error())
			// assuming AppError has been written to w
			// WriteError(w, e.Status(), e.Error())
		default:
			// Any other error types, default to serving a HTTP 500
			code := http.StatusInternalServerError
			log.Printf("[AppHandler] default: %+v\n", err)
			// http.Error(w, http.StatusText(code), code)
			WriteError(w, code, err.Error())
		}
	}
}
