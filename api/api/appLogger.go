// Package api :: appLogger.go
package api

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/urfave/negroni"
)

// AppLoggerResponseWriter implements http.ResponseWriter
type AppLoggerResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

// WriteHeader implements WriteHeader for http.ResponseWriter interface
func (alw *AppLoggerResponseWriter) WriteHeader(code int) {
	alw.ResponseWriter.WriteHeader(code)
	alw.StatusCode = code
}

// AppLogger is a middleware handler returns http.HandlerFunc
func AppLogger(handler http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		elapsed := time.Since(start) // time.Duration

		log.Printf(
			"[%s] %s %s | %s",
			name,
			r.Method,
			r.RequestURI,
			elapsed,
		)
	})
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

// NewLogger returns a new negroni.Logger instance
func NewLogger() *negroni.Logger {
	logger := &negroni.Logger{ALogger: log.New(os.Stdout, "", 0)}
	logger.SetDateFormat(negroni.LoggerDefaultDateFormat)
	logger.SetFormat("{{.StartTime}} {{.Request.Proto}} {{.Hostname}} | {{.Method}} {{.Path}} | {{.Status}} | {{.Duration}} \n")
	return logger
}
