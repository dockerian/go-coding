// +build all common pkg cfg context

// Package cfg :: context_test.go
package cfg

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestContextValue tests func Value in pkg/cfg/context.go
func TestContextValue(t *testing.T) {
	env := Env{
		"test": "result",
	}
	ctx := Context{
		Env:     &env,
		Context: context.WithValue(context.Background(), "Foo", "Bar"),
	}

	bar := ctx.Value("Foo")
	assert.Equal(t, "Bar", bar)

	resultNil := ctx.Value("DOES NOT EXIST")
	assert.Equal(t, nil, resultNil)

	resultValue := ctx.Value("Env")
	assert.NotNil(t, resultValue)

	resultEnv, okay := resultValue.(*Env)
	assert.True(t, okay)
	assert.NotNil(t, resultEnv)
	result := *resultEnv
	assert.Equal(t, env["test"], result["test"])

	nilContext := Context{
		Env: &env,
	}
	nilValue := nilContext.Value("NO_SUCH_THING")
	assert.Nil(t, nilValue)
}
