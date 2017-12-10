// Package api :: auth.go - auth handlers
package api

import (
	"log"
	"net/http"
	"strings"
)

// Auth creates a http.Handler wrapper to check api token in request
func Auth(next http.Handler, token string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log.Printf("[auth] checking %s\n", r.URL.Path)
		byPass := r.URL.Path == "/info"
		if !byPass {
			auth := r.Header["Authorization"]
			if len(auth) > 0 {
				// log.Printf("[auth] header: %+v\n", auth)
				authToken := strings.TrimPrefix(auth[0], "Token token=")
				if authToken != token {
					// log.Printf("[auth] token '%s' != '%s'\n", authToken, token)
					code := http.StatusUnauthorized
					http.Error(w, http.StatusText(code), code)
					return
				}
			}
		}
		log.Printf("[next] handler= %#v\n", next)
		next.ServeHTTP(w, r)
	})
}

// BasicAuth creates a http.Handler wrapper to check auth in request
func BasicAuth(next http.Handler, user, password string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log.Printf("[auth] checking %s\n", r.URL.Path)
		byPass := r.URL.Path == "/info"
		if !byPass {
			reqUser, reqPass, hasAuth := r.BasicAuth()
			// log.Printf("[auth] header: user= %v, pass= %v, hasAuth= %v\n", reqUser, reqPass, hasAuth)
			if !hasAuth || reqUser != user || reqPass != password {
				// log.Printf("[auth] check: user= %v, password= %v\n", user, password)
				w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
				code := http.StatusUnauthorized
				http.Error(w, http.StatusText(code), code)
				return
			}
		}
		log.Printf("[next] handler= %#v\n", next)
		next.ServeHTTP(w, r)
	})
}
