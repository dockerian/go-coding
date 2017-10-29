// +build all common pkg app env

// Package apimain :: appEnv_test.go
package apimain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAppEnv tests func api.TestAppEnv
func TestAppEnv(t *testing.T) {
	env := NewAppEnv()

	assert.NotNil(t, env)
	assert.True(t, len(env) > 0)
}
