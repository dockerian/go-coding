// Package api :: appHandler.go
package api

import (
	"log"
	"net/http"

	"github.com/dockerian/go-coding/utils"
)

// AppHandler struct wraps Env and Handle function implementing http.Handler
type AppHandler struct {
	utils.Env
	Handle func(e utils.Env, w http.ResponseWriter, r *http.Request) error
}

// ServeHTTP implements http.Handler
func (ah *AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if ah.Handle == nil {
		log.Println("[error] Missing handler func in AppHandler")
		code := http.StatusBadRequest
		http.Error(w, http.StatusText(code), code)
		return
	}
	err := ah.Handle(ah.Env, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// retrieve the status and print out
			log.Printf("[AppError] HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			// Any other error types, default to serving a HTTP 500
			code := http.StatusInternalServerError
			http.Error(w, http.StatusText(code), code)
		}
	}
}
