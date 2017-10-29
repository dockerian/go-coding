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

	"github.com/dockerian/go-coding/utils"
	"github.com/stretchr/testify/assert"
)

// TestAppHandler tests func api.AppHandler

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
			"test2", "result2", "new2", AppError{
				errors.New(http.StatusText(http.StatusBadRequest)),
				http.StatusBadRequest,
			},
			"Bad Request", http.StatusBadRequest,
		},
		{
			"test3", "result3", "new3",
			errors.New(http.StatusText(http.StatusBadRequest)),
			"Internal Server Error", http.StatusBadRequest,
		},
	}

	for idx, test := range tests {
		env := utils.Env{
			test.envKey: test.envOldText,
		}
		appHandler := &AppHandler{
			env,
			func(e utils.Env, w http.ResponseWriter, r *http.Request) error {
				e.Set(test.envKey, test.envNewText)
				w.WriteHeader(test.httpStatus)
				w.Header().Set("Content-Type", "application/json")
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
		msg3 := fmt.Sprintf("Test %2d.3 - http response - expected: %v, actual: %v",
			idx, test.httpResponse, body)
		t.Log(msg3)
		assert.Equal(t, test.httpResponse, body, msg3)
	}

	testHandler := AppHandler{}
	t.Logf("Test %2d: testHandler = %+v", len(tests), testHandler)
	testRecorder := httptest.NewRecorder()
	testHandler.ServeHTTP(testRecorder, req)
	assert.Equal(t, http.StatusBadRequest, testRecorder.Code)
}
