// +build all common pkg api auth

// Package api :: auth_test.go
package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AuthTestHandler struct
type AuthTestHandler struct{}

// ServeHTTP implements http.Handler
func (ath *AuthTestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// log.Printf("[AuthTestHandler] OK: %+v\n", w)
	// w.Write([]byte("Passed"))
}

// TestAuth tests func Auth
func TestAuth(t *testing.T) {
	tests := []struct {
		requestURL string
		statusCode int
		token      string
	}{
		{"http://test:8001/info", http.StatusOK, ""},
		{"http://test:8001/test", http.StatusOK, "token"},
		{"http://test:8001/test", http.StatusUnauthorized, ""},
	}

	for idx, test := range tests {
		req, _ := http.NewRequest("GET", test.requestURL, nil)
		req.Header.Set("Authorization", "Token token="+test.token)
		authHandler := Auth(&AuthTestHandler{}, "token")
		httRecorder := httptest.NewRecorder()
		t.Logf("Test %2d: expect %s\n", idx, http.StatusText(test.statusCode))
		authHandler.ServeHTTP(httRecorder, req)
		msg := fmt.Sprintf("Test %2d: httptest recorder= %#v\n", idx, httRecorder)
		assert.Equal(t, test.statusCode, httRecorder.Code, msg)
	}
}

// TestBasicAuth tests func BasicAuth
func TestBasicAuth(t *testing.T) {
	tests := []struct {
		requestURL string
		statusCode int
		user       string
		pass       string
	}{
		{"http://test:8001/info", http.StatusOK, "", ""},
		{"http://test:8001/test", http.StatusOK, "user", "pass"},
		{"http://test:8001/test", http.StatusUnauthorized, "", ""},
	}

	for idx, test := range tests {
		req, _ := http.NewRequest("GET", test.requestURL, nil)
		req.SetBasicAuth(test.user, test.pass)
		authHandler := BasicAuth(&AuthTestHandler{}, "user", "pass")
		httRecorder := httptest.NewRecorder()
		t.Logf("Test %2d: expect %s\n", idx, http.StatusText(test.statusCode))
		authHandler.ServeHTTP(httRecorder, req)
		msg := fmt.Sprintf("Test %2d: httptest recorder= %#v\n", idx, httRecorder)
		assert.Equal(t, test.statusCode, httRecorder.Code, msg)
		httRecorder.Flush()
	}
}
