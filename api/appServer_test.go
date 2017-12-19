// +build all app api server

// Package api :: appServer_test.go
package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockAppServer struct
type MockAppServer struct {
	certFile          string
	keyFile           string
	funcBeenCalled    bool
	funcTLSBeenCalled bool
	result            string
}

// ListenAndServe implements AppServerInterface
func (mas *MockAppServer) ListenAndServe() error {
	mas.funcBeenCalled = true
	mas.result = "ListenAndServe has been called"
	return nil
}

// ListenAndServeTLS implements AppServerInterface
func (mas *MockAppServer) ListenAndServeTLS(cert, key string) error {
	mas.keyFile = key
	mas.certFile = cert
	mas.funcTLSBeenCalled = true
	mas.result = "ListenAndServeTLS has been called"
	return nil
}

// TestAppServer tests func NewAppServer in app/sng/appEnv.go
func TestAppServer(t *testing.T) {
	app := NewAppServer()
	assert.NotNil(t, app.Ctx)
	assert.NotNil(t, app.Doc)
	assert.NotNil(t, app.Handler)
	assert.NotNil(t, app.Router)
	assert.NotNil(t, app.Server)
}

// TestListenAndServe tests ListenAndServe
func TestListenAndServe(t *testing.T) {
	ctx := NewAppContext()
	mockApp := &MockAppServer{
		funcBeenCalled:    false,
		funcTLSBeenCalled: false,
		result:            "",
	}

	ListenAndServe(mockApp, ctx)
	assert.Equal(t, "ListenAndServe has been called", mockApp.result)
	assert.True(t, mockApp.funcBeenCalled)
}
