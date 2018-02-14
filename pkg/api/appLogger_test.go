// +build all common pkg api logger

// Package api :: appLogger_test.go
package api

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAppLogger tests func api.AppLogger
func TestAppLogger(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/logger", nil)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name         string
		httpResponse string
		httpStatus   int
	}{
		{"test1", "Bad", http.StatusBadRequest},
		{"test2", "Good", http.StatusOK},
	}

	for idx, test := range tests {
		appHandler := AppLogger(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(test.httpStatus)
				io.WriteString(w, test.httpResponse)
			}), test.name)

		rr := httptest.NewRecorder()
		appHandler.ServeHTTP(rr, req)

		// check http status code in httptest recorder
		msg2 := fmt.Sprintf("Test %2d.2 - http status - expected: %v, actual: %v",
			idx, test.httpStatus, rr.Code)
		t.Log(msg2)
		assert.Equal(t, test.httpStatus, rr.Code, msg2)

		// check response body
		body := strings.Trim(rr.Body.String(), "\n")
		msg3 := fmt.Sprintf("Test %2d.3 - http response - expected: %v, actual: %v",
			idx, test.httpResponse, body)
		t.Log(msg3)
		assert.Equal(t, test.httpResponse, body, msg3)
	}
}

// TestResponseWriter struct
type TestResponseWriter struct {
	ResponseHeader http.Header
	StatusCode     int
}

// Header implements http.ResponseWriter interface
func (trw *TestResponseWriter) Header() http.Header {
	return trw.ResponseHeader
}

// Write implements http.ResponseWriter interface
func (trw *TestResponseWriter) Write([]byte) (int, error) {
	return 200, nil
}

// WriteHeader implements http.ResponseWriter interface
func (trw *TestResponseWriter) WriteHeader(code int) {
	trw.ResponseHeader = http.Header{"Status": []string{http.StatusText(code)}}
	trw.StatusCode = code
}

// TestAppLoggerResponseWriter tests func api.AppAppLoggerResponseWriter
func TestAppLoggerResponseWriter(t *testing.T) {
	var resw http.ResponseWriter = &TestResponseWriter{}
	test := AppLoggerResponseWriter{resw, 200}
	test.WriteHeader(http.StatusNotFound)
	header := test.ResponseWriter.Header()
	expected := http.StatusText(http.StatusNotFound)
	assert.Equal(t, expected, header["Status"][0])
	assert.Equal(t, http.StatusNotFound, test.StatusCode)
	assert.Equal(t, resw.Header(), header)
}

// TestNewLogger tests func api.NewLogger
func TestNewLogger(t *testing.T) {
	test := NewLogger("prefix")
	assert.NotNil(t, test.ALogger)
}
