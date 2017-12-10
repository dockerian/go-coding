// +build all common pkg app env

// Package apimain :: appEnv_test.go
package apimain

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAppEnvGetConfig tests func GetConfig in app/sng/appEnv.go
func TestAppEnvGetConfig(t *testing.T) {
	key := "__NON_EXIST_VARIABLE__"
	val := "This is to test reading from environment variable"
	def := "Here is some default value for non-exist key"
	config := GetConfig()
	result := config.Get(key, def)
	assert.Equal(t, def, result, "should get value from default value")
	os.Setenv(key, val)
	keyval := config.Get(key, def)
	assert.Equal(t, val, keyval, "should get value from os environ")
}

// TestAppEnv tests func NewAppContext in app/sng/appEnv.go
func TestAppEnv(t *testing.T) {
	key := "__NON_EXIST_ENVOIRONMENT_KEY__"
	val := "This is to test reading from environment key"
	os.Setenv(key, val)

	pwd, _ := os.Getwd()
	context := NewAppContext()
	env := context.Env
	assert.NotNil(t, context.Env)
	assert.Equal(t, pwd, env.Get("appLocation"))
	assert.Equal(t, "README.md", env.Get("docIndex"))
	assert.Equal(t, val, env.Get(key))
}
