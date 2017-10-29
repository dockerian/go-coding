// Package api :: route.go
package api

import (
	"net/http"

	"github.com/dockerian/go-coding/utils"
	"github.com/gorilla/mux"
)

// RouteConfig struct
type RouteConfig struct {
	Pattern     string
	Method      string
	Name        string
	HandlerFunc func(e utils.Env, w http.ResponseWriter, r *http.Request) error
	Proxy       ProxyRoute
}

// RouteConfigs is a list of RouteConfig struct
type RouteConfigs []RouteConfig

// NewRouter returns *mux.Router
func NewRouter(env utils.Env, routeConfigs RouteConfigs) *mux.Router {
	var methods []string
	var handler http.Handler
	// mux.Router implements http.Handler
	router := mux.NewRouter().StrictSlash(true) // mux.Router
	for _, config := range routeConfigs {
		handler = &AppHandler{Env: env, Handle: config.HandlerFunc}
		methods = []string{config.Method}

		if config.Proxy.Prefix != "" && config.Proxy.RedirectURL != "" {
			handler = ProxyHandler(config.Proxy.Prefix, config.Proxy.RedirectURL)
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
