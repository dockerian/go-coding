// Package apimain :: appServer.go
package apimain

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dockerian/go-coding/pkg/api"
	"github.com/dockerian/go-coding/pkg/cfg"
)

// ListenAndServe starts a server
func ListenAndServe(server api.AppServerInterface, ctx *cfg.Context) error {
	env := ctx.Env
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
func NewAppServer() *api.AppServer {
	ctx := NewAppContext()
	env := ctx.Env
	doc := api.NewMarkdown("/doc", env.Get("docLocation"), env.Get("docIndex"))
	app := api.AppServer{
		Ctx:     ctx,
		Doc:     doc,
		Handler: api.NewMiddleware(), // can initiate multiple negroni.Handler
		Logger:  api.NewLogger(""),   // *negroni.Logger is a negroni.Handler
		Router:  api.NewRouter(*ctx, routeConfigs),
	}

	// For other middleware, see https://github.com/urfave/negroni#third-party-middleware
	app.Handler.Use(app.Doc)
	app.Handler.Use(api.NewRecovery())
	app.Handler.UseHandler(app.Router) // map to http.Handler or mux.Router
	app.Handler.Use(app.Logger)        // map to negroni.Handler

	app.Server = &http.Server{
		Addr:           app.Ctx.Env.Get("appAddress"),
		Handler:        app.Handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &app
}
