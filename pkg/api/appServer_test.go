// +build all common pkg api server

// Package api :: appServer_test.go
package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNegroni tests functions to get negroni middleware and handler
func TestNegroni(t *testing.T) {
	mw := NewMiddleware()
	handlers := mw.Handlers()
	t.Logf("negroni.New - handlers count= %d\n", len(handlers))
	assert.True(t, len(handlers) == 0)
	mw.Use(NewRecovery())
	assert.True(t, len(mw.Handlers()) == 1)
}
