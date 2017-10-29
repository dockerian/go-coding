// +build all common pkg api route

// Package api :: route_test.go
package api

import (
	"net/http"
	"testing"

	"github.com/dockerian/go-coding/utils"
	"github.com/stretchr/testify/assert"
)

// MockHandler implements AppHandler
func MockHandler(env utils.Env, w http.ResponseWriter, r *http.Request) error {
	env["test"] = "response"
	return nil
}

// TestNewRouter tests func api.NewRouter
func TestNewRouter(t *testing.T) {
	env := utils.Env{
		"test": "result",
	}
	routeConfigs := RouteConfigs{
		{
			"/mock", "*", "Mock", MockHandler,
			ProxyRoute{"/prefix", "http://redirect"},
		},
	}
	router := NewRouter(env, routeConfigs)
	// t.Logf("router: %+v\n", router)
	assert.NotNil(t, router)
	assert.NotNil(t, router.Get("Mock"))
	assert.NotNil(t, router.GetRoute("Mock"))
	assert.Nil(t, router.NotFoundHandler)

}
