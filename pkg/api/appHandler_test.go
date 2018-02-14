// +build all common pkg api handler

// Package api :: appHandler_test.go
package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dockerian/go-coding/pkg/cfg"
	"github.com/stretchr/testify/assert"
)

// TestAppHandler tests func api.AppHandler
func TestAppHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/app/handler", nil)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		envKey       string
		envOldText   string
		envNewText   string
		handlerError error
		httpResponse string
		httpStatus   int
	}{
		{
			"test1", "result1", "new1", nil, `{"res1": "response1"}`, http.StatusOK,
		},
		{
			"test2", "result2", "new2", NewAppError(http.StatusBadRequest, ""),
			"", http.StatusBadRequest,
		},
		{
			"test3", "result3", "new3",
			errors.New("new server error"),
			`{
				  "message": "new server error",
				  "code": 500
			}`,
			http.StatusInternalServerError,
		},
	}

	for idx, test := range tests {
		env := cfg.Env{
			test.envKey: test.envOldText,
		}
		ctx := cfg.Context{
			Env: &env,
		}
		appHandler := &AppHandler{
			ctx,
			func(ctx cfg.Context, w http.ResponseWriter, r *http.Request) error {
				env := ctx.Env
				env.Set(test.envKey, test.envNewText)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(test.httpStatus)
				if test.handlerError == nil {
					io.WriteString(w, test.httpResponse)
				}
				return test.handlerError
			},
		}
		rr := httptest.NewRecorder()
		appHandler.ServeHTTP(rr, req)

		// check env setting
		val1 := env.Get(test.envKey)
		msg1 := fmt.Sprintf("Test %2d.1 - env[%s] - expected: %v, actual: %v",
			idx, test.envKey, test.envNewText, val1)
		t.Log(msg1)
		assert.Equal(t, test.envNewText, val1)

		// check http status code in httptest recorder
		msg2 := fmt.Sprintf("Test %2d.2 - http status - expected: %v, actual: %v",
			idx, test.httpStatus, rr.Code)
		t.Log(msg2)
		assert.Equal(t, test.httpStatus, rr.Code, msg2)

		// check response body
		body := strings.Trim(rr.Body.String(), "\n")
		expt := strings.Replace(test.httpResponse, "\t", "", -1)
		msg3 := fmt.Sprintf("Test %2d.3 - http response - expected: %v, actual: %v",
			idx, expt, body)
		t.Log(msg3)
		assert.Equal(t, expt, body, msg3)
	}

	testHandler := AppHandler{}
	t.Logf("Test %2d: testHandler = %+v", len(tests), testHandler)
	testRecorder := httptest.NewRecorder()
	testHandler.ServeHTTP(testRecorder, req)
	assert.Equal(t, http.StatusBadRequest, testRecorder.Code)
}
