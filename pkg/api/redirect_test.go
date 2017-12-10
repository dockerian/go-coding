// +build all common pkg api redirect

// Package api :: redirect_test.go
package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// httpRedirectMock replaces httpRedirect in tests
func httpRedirectMock(w http.ResponseWriter, r *http.Request, url string, code int) {
	log.Println("[redirect] mock handling url:", url)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(url))
}

// TestRedirectHandler tests Proxy constructor
func TestRedirectHandler(t *testing.T) {
	orURL := "http://host:80/foo/test/?v=1"
	tests := []struct {
		prefix       string
		requestURL   string
		redirectURL  string
		response     string
		expectedCode int
	}{
		{"", "", "http://redir1", "cannot match prefix '' in ''\n", http.StatusBadRequest},
		{"/bar", orURL, "http://redir2", orURL, http.StatusOK},
		{"/foo", orURL, "http://redir3", "http://redir3/test/?v=1", http.StatusOK},
		{"/foo/test", orURL, "http://redir4", "http://redir4/?v=1", http.StatusOK},
	}

	originalFunc := httpRedirect
	httpRedirect = httpRedirectMock
	defer func() {
		httpRedirect = originalFunc
	}()

	for idx, test := range tests {
		handler := RedirectHandler(test.prefix, test.redirectURL)
		switch v := handler.(type) {
		case http.HandlerFunc:
			log.Println("[redirect] PASS: RedirectHandler is an http.HandlerFunc")
		default:
			msg := fmt.Sprintf("%v is not http.HandlerFunc", v)
			assert.Fail(t, msg)
		}

		req, _ := http.NewRequest("GET", test.requestURL, nil)
		rwr := httptest.NewRecorder()

		msg := fmt.Sprintf("Test %2d: prefix (%s) => %s\n", idx, test.prefix, test.response)
		t.Log(msg)
		handler.ServeHTTP(rwr, req)

		// log.Printf("Test %2d - request: %s - %#v\n", idx, req.URL, req)
		// log.Printf("Test %2d - response: %#v\n", idx, rwr)
		data, _ := ioutil.ReadAll(rwr.Body)
		assert.Equal(t, test.response, string(data), msg)
		assert.Equal(t, test.expectedCode, rwr.Code, msg)
	}
}
