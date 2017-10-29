// Package apimain :: appServer.go
package apimain

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dockerian/go-coding/api/api"
	"github.com/dockerian/go-coding/utils"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// AppServer represents partial http.Server
type AppServer struct {
	Env     utils.Env          // app Env
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

// ListenAndServe starts a server
func ListenAndServe(server AppServerInterface, env utils.Env) error {
	appName := env.Get("appName")
	appAddress := env.Get("appAddress")
	keyFile := env.Get("sslKeyFile")
	certFile := env.Get("sslCertFile")
	if _, err := os.Stat(certFile); !os.IsNotExist(err) {
		if _, err := os.Stat(keyFile); !os.IsNotExist(err) {
			log.Printf("[sng] starting %s (https) %s", appName, appAddress)
			return server.ListenAndServeTLS(certFile, keyFile)
		}
	}
	log.Printf("[sng] starting %s (http) %s", appName, appAddress)
	return server.ListenAndServe()
}

// NewAppServer constructs AppServer with
// - a new *mux.Router
// - negroni middlewares
//   any middleware by negroni.Use() should implements negroni.Handler interface:
//      ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
//   otherwise, by negroni.UseHandler() should implements http.Handler
//   see https://github.com/urfave/negroni#handlers
func NewAppServer() *AppServer {
	env := NewAppEnv()
	doc := api.NewMarkdown("/doc", env.Get("docLocation"), env.Get("docIndex"))
	app := AppServer{
		Env:     env,
		Doc:     doc,
		Handler: negroni.New(),   // can initiate multiple negroni.Handler
		Logger:  api.NewLogger(), // *negroni.Logger is a negroni.Handler
		Router:  api.NewRouter(env, routeConfigs),
	}

	// For other middleware, see https://github.com/urfave/negroni#third-party-middleware
	app.Handler.Use(app.Doc)
	app.Handler.Use(negroni.NewRecovery())
	app.Handler.UseHandler(app.Router) // map to http.Handler or mux.Router
	app.Handler.Use(app.Logger)        // map to negroni.Handler

	app.Server = &http.Server{
		Addr:           app.Env.Get("appAddress"),
		Handler:        app.Handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &app
}
