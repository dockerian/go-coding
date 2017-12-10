// Package api :: appLogger.go - logging handlers
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

// WriteHeader implements http.ResponseWriter interface
func (alw *AppLoggerResponseWriter) WriteHeader(code int) {
	alw.ResponseWriter.WriteHeader(code)
	alw.StatusCode = code
}

// AppLogger is a middleware handler returns http.HandlerFunc
func AppLogger(handler http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		appLoggerWriter := &AppLoggerResponseWriter{
			ResponseWriter: w,
		}
		handler.ServeHTTP(appLoggerWriter, r)
		statusCode := appLoggerWriter.StatusCode
		statusText := http.StatusText(statusCode)
		elapsed := time.Since(start) // time.Duration

		log.Printf(
			"[%s] %s %s | %v - %v %s",
			name,
			r.Method,
			r.RequestURI,
			elapsed,
			statusCode,
			statusText,
		)
	})
}

// NewLogger returns a new negroni.Logger instance
func NewLogger(prefix string) *negroni.Logger {
	logger := &negroni.Logger{ALogger: log.New(os.Stdout, prefix, 0)}
	logger.SetDateFormat(negroni.LoggerDefaultDateFormat)
	logger.SetFormat("{{.StartTime}} {{.Request.Proto}} {{.Hostname}} | {{.Method}} {{.Path}} | {{.Status}} | {{.Duration}} \n")
	return logger
}
