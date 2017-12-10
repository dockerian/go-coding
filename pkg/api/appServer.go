// Package api :: appServer.go
package api

import (
	"github.com/dockerian/go-coding/pkg/cfg"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// AppServer represents partial http.Server
type AppServer struct {
	Ctx     *cfg.Context       // app context
	Doc     negroni.Handler    // negroni static handler
	Handler *negroni.Negroni   // negroni handler
	Logger  *negroni.Logger    // negroni logger
	Router  *mux.Router        // mux router
	Server  AppServerInterface // http.Server implements ListenAndServe
}

// AppServerInterface represents partial http.Server interface
type AppServerInterface interface {
	ListenAndServe() error
	ListenAndServeTLS(string, string) error
}

// NewMiddleware returns a negroni middleware
func NewMiddleware() *negroni.Negroni {
	return negroni.New()
}

// NewRecovery returns a negroni recovery handler
func NewRecovery() *negroni.Recovery {
	return negroni.NewRecovery()
}
