// Package api :: proxy.go - proxy handler
//
// For each Proxy redirect/forward call, api.ProxyRoute defines a RedirectURL
// per prefix path. The api.ProxyRoute implements http.Handler interface
// so that the struct pointer itself can be wrapped in a routing configuration;
// optionally, a Proxy() can construct a http.Handler with prefix and
// predefined redirect URL.
package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dockerian/go-coding/pkg/str"
)

var (
	// ioCopy is a copy function for response body
	ioCopy = io.Copy
	// bodyReadAll is a reader function for response body
	bodyReadAll = ioutil.ReadAll
	// proxyClient is an http client
	proxyClient ProxyClient = &http.Client{
		Timeout: 5 * time.Second,
	}
)

type bodyReadFunc func(r io.Reader) ([]byte, error)

// ProxyClient interface
type ProxyClient interface {
	Do(*http.Request) (*http.Response, error)
}

// ProxyRoute struct defines a redirecting URL based on pattern
// by converting matched Prefix path to RedirectURL
type ProxyRoute struct {
	// Prefix defines the matching prefix path to be replaced
	Prefix string
	// RedirectOnly specifies to use httpRedirect rather than a proxy client
	RedirectOnly bool
	// RedirectURL defines the replacing URL path
	RedirectURL string
}

// Proxy generates new request with prefix path to redirected url
func Proxy(prefix, redirectURL string, w http.ResponseWriter, r *http.Request) error {
	// log.Printf("[proxy] parsing '%s' in '%s' => '%s'", prefix, r.URL, redirectURL)
	restURL := str.ReplaceProxyURL(r.URL.String(), prefix, redirectURL)
	if restURL == "" {
		log.Printf("[proxy] empty url: '%s' by prefix '%s'\n", r.URL, prefix)
		msg := fmt.Sprintf("cannot match prefix '%s' in '%s'", prefix, r.URL)
		return WriteError(w, http.StatusBadRequest, msg)
	}

	log.Printf("[proxy] %s (%s) => %+v\n", prefix, r.URL, restURL)
	proxyReq, err := http.NewRequest(r.Method, restURL, r.Body)
	if err != nil {
		log.Println("[proxy] err:", err)
		return WriteError(w, http.StatusBadRequest, err.Error())
	}
	proxyReq.Header.Add("Connection", "close")

	// clone copying header, not just a shallow copy proxyReq.Header = req.Header
	proxyReq.Header = make(http.Header)
	for key, val := range r.Header {
		proxyReq.Header[key] = val
	}
	proxyReq.Header.Set("Host", r.Host)
	// set "accept-encoding" to prevent from g-zipped content by browser client
	proxyReq.Header.Set("Accept-Encoding", "gzip;q=0,deflate;q=0")
	proxyReq.Header.Set("X-Forwarded-For", r.RemoteAddr)

	resp, err := proxyClient.Do(proxyReq)
	if err != nil {
		log.Println("[proxy] err:", err)
		return WriteError(w, http.StatusBadGateway, err.Error())
	}
	defer resp.Body.Close()

	// setting headers
	for name, values := range resp.Header {
		w.Header()[name] = values
	}
	// writting status after setting headers and before copying body
	w.WriteHeader(resp.StatusCode)

	// copying body
	_, errCopy := ioCopy(w, resp.Body)
	if errCopy != nil {
		log.Println("[proxy] io.Copy err:", errCopy)
		return WriteError(w, http.StatusInternalServerError, errCopy.Error())
	}

	if code := resp.StatusCode / 100; code == 2 {
		log.Println("[proxy] success:", resp.Status)
	}

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
