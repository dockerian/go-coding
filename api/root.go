package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dockerian/go-coding/api/info"
	"github.com/gorilla/mux"
)

var (
	// ListenPort is a configurable http port
	ListenPort = 8181

	// RootRoutes configures root routes
	// optionally read from config or move this to routes.go
	RootRoutes = Routes{
		{
			"/", "GET", rootHandler, "Index",
		},
		{
			"/info", "GET", info.GetInfo, "Info",
		},
	}
)

// Route struct encapsulates an http route
type Route struct {
	Pattern string
	Method  string
	Handler http.HandlerFunc
	Name    string
}

// Routes struct is an array of Route
type Routes []Route

// Index is api root entry pointer
func Index() {
	// handleRequests()
	// or
	router, port := rootRouter(RootRoutes)
	log.Printf("Listen on %v ...\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// handleRequests (deprecated) is using basic http
func handleRequests() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/info", info.GetInfo)
	port := fmt.Sprintf(":%d", ListenPort)
	log.Fatal(http.ListenAndServe(port, nil))
}

// rootHandler is api root handler
func rootHandler(res http.ResponseWriter, req *http.Request) {
	info.GetInfo(res, req)
}

// rootRouter returns a configured mux.Router
func rootRouter(routes []Route) (*mux.Router, string) {
	router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", rootHandler)
	// router.HandleFunc("/info", info.GetInfo)
	// or
	for _, route := range routes {
		var handler http.Handler
		handler = Logger(route.Handler, route.Name)
		router.
			Path(route.Pattern).
			Methods(route.Method).
			Name(route.Name).
			Handler(handler)
	}

	port := fmt.Sprintf(":%d", ListenPort)

	return router, port
}
