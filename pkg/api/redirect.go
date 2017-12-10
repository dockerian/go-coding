// Package api :: redirect.go - http redirect handlers
package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/dockerian/go-coding/pkg/str"
)

var (
	// httpRedirect is an http redirect function
	httpRedirect = http.Redirect
)

// RedirectHandler constructs an http.Handler by prefix path and redirect URL
func RedirectHandler(prefix, redirectURL string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Redirect(prefix, redirectURL, w, r)
	})
}

// Redirect forwards call by prefix path to redirected url
func Redirect(prefix, redirectURL string, w http.ResponseWriter, r *http.Request) error {
	requestURL := r.URL.String()
	forwardURL := str.ReplaceProxyURL(requestURL, prefix, redirectURL)
	if forwardURL == "" {
		msg := fmt.Sprintf("cannot match prefix '%s' in '%s'", prefix, r.URL)
		log.Printf("[redirect] err: %s", msg)
		http.Error(w, msg, http.StatusBadRequest)
		return errors.New(msg)
	}
	log.Printf("[redirect] %s (%s) => %+v\n", prefix, r.URL, forwardURL)
	httpRedirect(w, r, forwardURL, http.StatusTemporaryRedirect)
	return nil
}

// RedirectToHTTPS is a middleware handler returns http.HandlerFunc
func RedirectToHTTPS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		proto := req.Header.Get("x-forwarded-proto")
		if proto == "http" || proto == "HTTP" {
			http.Redirect(res, req,
				fmt.Sprintf("https://%s%s", req.Host, req.URL),
				http.StatusPermanentRedirect)
			return
		}

		next.ServeHTTP(res, req)
	})
}
