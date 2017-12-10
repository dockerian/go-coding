// Package api :: appHandler.go
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
		log.Printf("[error] %s\n", message)
		WriteError(w, http.StatusBadRequest, message)
		return
	}
	err := ah.Handle(ah.Context, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// retrieve the status and print out
			log.Printf("[AppError] HTTP %d - %s", e.Status(), e)
			// TODO: [jzhu] check if Error has been written
			// http.Error(w, e.Error(), e.Status())
		default:
			// Any other error types, default to serving a HTTP 500
			code := http.StatusInternalServerError
			http.Error(w, http.StatusText(code), code)
		}
	}
}
