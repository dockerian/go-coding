// Package api :: route.go
package api

import (
	"log"
	"net/http"

	"github.com/dockerian/go-coding/pkg/cfg"
	"github.com/gorilla/mux"
)

// RouteConfig struct
type RouteConfig struct {
	Pattern     string
	Method      string
	Name        string
	HandlerFunc func(e cfg.Context, w http.ResponseWriter, r *http.Request) error
	Proxy       ProxyRoute
}

// RouteConfigs is a list of RouteConfig struct
type RouteConfigs []RouteConfig

// NewRouter returns *mux.Router
func NewRouter(ctx cfg.Context, routeConfigs RouteConfigs) *mux.Router {
	var methods []string
	var handler http.Handler
	var noProxy bool
	// mux.Router implements http.Handler
	router := mux.NewRouter().StrictSlash(true) // mux.Router
	allowRedirect, _ := ctx.Env.GetValue("allowRedirect").(bool)

	log.Println("[router] allow redirect:", allowRedirect)

	for _, config := range routeConfigs {
		handler = &AppHandler{Context: ctx, Handle: config.HandlerFunc}
		methods = []string{config.Method}
		noProxy = allowRedirect &&
			config.Proxy.RedirectOnly &&
			config.Proxy.RedirectURL != "" &&
			config.Proxy.Prefix != ""

		if noProxy {
			handler = RedirectHandler(config.Proxy.Prefix, config.Proxy.RedirectURL)
		}
		if config.Method == "*" {
			methods = []string{"DELETE", "GET", "PUT", "POST"}
		}

		// mux.Router has Methods and Path functions both return *Route
		router.
			Path(config.Pattern).
			Methods(methods...).
			Handler(AppLogger(handler, config.Name)).
			Name(config.Name)
	}

	return router
}
