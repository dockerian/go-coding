// +build all common pkg api route

// Package api :: route_test.go
package api

import (
	"net/http"
	"testing"

	"github.com/dockerian/go-coding/pkg/cfg"
	"github.com/stretchr/testify/assert"
)

// MockHandler implements AppHandler
func MockHandler(ctx cfg.Context, w http.ResponseWriter, r *http.Request) error {
	env := *ctx.Env
	env["test"] = "response"
	return nil
}

// TestNewRouter tests func api.NewRouter
func TestNewRouter(t *testing.T) {
	env := cfg.Env{
		"test": "result",
	}
	ctx := cfg.Context{
		Env: &env,
	}
	routeConfigs := RouteConfigs{
		{
			"/mock", "*", "Mock", MockHandler,
			ProxyRoute{"/prefix", "http://redirect"},
		},
	}
	router := NewRouter(ctx, routeConfigs)
	// t.Logf("router: %+v\n", router)
	assert.NotNil(t, router)
	assert.NotNil(t, router.Get("Mock"))
	assert.NotNil(t, router.GetRoute("Mock"))
	assert.Nil(t, router.NotFoundHandler)

}
