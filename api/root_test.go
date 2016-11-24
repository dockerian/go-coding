// +build all api http test

package api

// see
// - https://nathanleclaire.com/blog/2015/10/10/interfaces-and-composition-for-effective-unit-testing-in-golang/
// - https://medium.com/@PurdonKyle/building-a-unit-testable-api-in-golang-b42ff1fcbae7#.i006w37kh
// - https://blog.gopheracademy.com/advent-2014/testing-microservices-in-go/
// - https://code.thebur.net/2016/03/22/mocking-http-endpoints-in-golang/

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dockerian/go-coding/api/info"
	"github.com/stretchr/testify/assert"
)

var (
	mockServer *httptest.Server

	mockRoutes = Routes{
		{
			"/", "GET", mockHandler, "Index",
		},
		{
			"/info", "GET", info.GetInfo, "Info",
		},
	}

	mockTestData = []MockData{
		{"mock name", "mock info"},
		{"Go API", "Go RESTful API Example"},
	}

	reader io.Reader
)

// MockData struct
type MockData struct {
	Name string `json:"name,omitempty"`
	Info string `json:"desc,omitempty"`
}

func init() {
}

// mockHandler mocks api root handler
func mockHandler(res http.ResponseWriter, req *http.Request) {
	// io.WriteString(res, `{"name": "mock name", "desc": "mock info"}`)
	// or
	json.NewEncoder(res).Encode(mockTestData[0])
}

// TestRootRouter tests rootRouter
func TestRootRouter(t *testing.T) {
	route, _ := rootRouter(mockRoutes)
	mockServer = httptest.NewServer(route)
	defer mockServer.Close()

	tests := []struct {
		path string
		data MockData
		code int
	}{
		{"", mockTestData[0], http.StatusOK},
		{"info", mockTestData[1], http.StatusOK},
		{"none", MockData{}, http.StatusNotFound},
	}

	for i, test := range tests {
		t.Logf("Test %d: /%+v - %v - %+v\n", i, test.path, test.code, test.data)

		reader = strings.NewReader("")
		url := fmt.Sprintf("http://localhost/%v", test.path)
		req, err := http.NewRequest("GET", url, reader)
		rec := httptest.NewRecorder()
		route.ServeHTTP(rec, req)

		res := rec.Result()
		msg := fmt.Sprintf("Expected status: %v -- Received: %d (%s)\n", test.code, res.StatusCode, res.Status)
		assert.Equal(t, test.code, res.StatusCode, msg)

		var data MockData
		decoder := json.NewDecoder(res.Body)
		decoder.Decode(&data)

		msg = fmt.Sprintf("Expected %+v", test.data)
		assert.EqualValues(t, test.data, data, msg)

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(body, &data)
		res.Body.Close()

		assert.EqualValues(t, test.data, data, msg)
	}
}
