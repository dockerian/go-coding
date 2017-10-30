// Package api :: proxy.go - proxy handler
// For each Proxy redirect/forward call, api.ProxyRoute defines a RedirectURL
// per prefix path. The api.ProxyRoute implements http.Handler interface
// so that the struct pointer itself can be wrapped in a routing configuration;
// optionally, a Proxy() can construct a http.Handler with prefix and
// predefined redirect URL.
package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dockerian/go-coding/utils"
)

var (
	// bodyReadAll is a reader function for response body
	bodyReadAll = ioutil.ReadAll
	// proxyClient is an http client
	proxyClient ProxyClient = &http.Client{}
)

// ProxyClient interface
type ProxyClient interface {
	Do(*http.Request) (*http.Response, error)
}

// ProxyRoute struct defines a redirecting URL based on pattern
type ProxyRoute struct {
	Prefix      string
	RedirectURL string
}

// Proxy forwards call with prefix to redirected url
func Proxy(prefix, redirectURL string, w http.ResponseWriter, r *http.Request) error {
	// log.Printf("[proxy] parsing '%s' in '%s'", prefix, r.URL)
	restURL := utils.ReplaceProxyURL(r.URL.String(), prefix, redirectURL)
	if restURL == "" {
		msg := fmt.Sprintf("cannot match prefix '%s' in '%s'", prefix, r.URL)
		log.Printf("[proxy] err: %s", msg)
		http.Error(w, msg, http.StatusBadRequest)
		return errors.New(msg)
	}
	log.Printf("[proxy] %s (%s) => %+v\n", prefix, r.URL, restURL)

	proxyReq, err := http.NewRequest(r.Method, restURL, r.Body)
	// clone copying header, not just a shallow copy proxyReq.Header = req.Header
	proxyReq.Header = make(http.Header)
	for key, val := range r.Header {
		proxyReq.Header[key] = val
	}
	proxyReq.Header.Set("Host", r.Host)
	// set "accept-encoding" to prevent from g-zipped content by browser client
	proxyReq.Header.Set("Accept-Encoding", "gzip;q=0,deflate;q=0")
	proxyReq.Header.Set("X-Forwarded-For", r.RemoteAddr)
	proxyReq.Body = r.Body

	resp, err := proxyClient.Do(proxyReq)
	if err != nil {
		log.Println("[proxy] err:", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return err
	}
	defer resp.Body.Close()

	data, err := bodyReadAll(resp.Body)
	if err != nil {
		log.Println("[response] err:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	log.Println("[response] data:", data)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

	return nil
}

// ProxyHandler constructs an http.Handler by prefix path and redirect URL
func ProxyHandler(prefix, redirectURL string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Proxy(prefix, redirectURL, w, r)
	})
}

// ServeHTTP implements http.Handler
func (p *ProxyRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Proxy(p.Prefix, p.RedirectURL, w, r)
}
