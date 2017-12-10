// +build all common pkg api proxy

// Package api :: proxy_test.go
package api

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockProxyClient struct
type MockProxyClient struct {
	Request      *http.Request
	ResponseText string
	Err          error
}

// Do implements ProxyClient interface
func (mpc *MockProxyClient) Do(r *http.Request) (*http.Response, error) {
	mpc.Request = r
	response := &http.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte(mpc.ResponseText))),
	}
	return response, mpc.Err
}

// TestProxyFunc tests func Proxy
func TestProxyFunc(t *testing.T) {
	tests := []struct {
		prefix      string
		redirectURL string
		reqURL      string
		restURL     string
		response    string
		errText     string
		errFunc     func(io.Reader) ([]byte, error)
	}{
		{
			"/prefix1", "http://redir1", "/prefix1/test", "http://redir1/test",
			"good", "", nil,
		},
		{
			"/pre2", "http://re2", "/pre2/test/value", "http://re2/test/value",
			"BAD GATEWAY", "bad gateway err test", nil,
		},
		{
			"/pre3", "http://re3", "/pre3/test/3", "http://re3/test/3",
			"read err test", "", func(_ io.Reader) ([]byte, error) {
				return []byte{}, errors.New("read err test")
			},
		},
	}

	savedClient := proxyClient
	savedReadAll := bodyReadAll
	defer func() {
		// restoring proxyClient (see definition in proxy.go)
		proxyClient = savedClient
	}()

	for idx, test := range tests {
		var err error
		if test.errText != "" {
			err = errors.New(test.errText)
		}
		if test.errFunc == nil {
			bodyReadAll = savedReadAll
		} else {
			bodyReadAll = test.errFunc
		}
		client := &MockProxyClient{
			ResponseText: test.response, Err: err,
		}
		req, _ := http.NewRequest("GET", test.reqURL, nil)
		req.Header.Set("Foobar", "Foobar header test")
		msg := fmt.Sprintf("%s => %s", test.reqURL, test.restURL)
		t.Logf("Test %2d: %s\n", idx+1, msg)
		proxyClient = client
		rwr := httptest.NewRecorder()
		// here start to test the function
		Proxy(test.prefix, test.redirectURL, rwr, req)

		data, err := savedReadAll(rwr.Body)
		proxyHeader := client.Request.Header
		proxyReq := client.Request

		// t.Logf("request: %#v\n", req)
		// t.Logf("proxy: %#v (%s)\n", proxyReq, proxyReq.URL)
		// t.Logf("record: %#v\n", rwr)
		assert.Equal(t, req.URL.Path, test.reqURL)
		assert.Equal(t, req.Method, proxyReq.Method)
		assert.Equal(t, req.Host, proxyHeader.Get("Host"))
		assert.Equal(t, "Foobar header test", proxyHeader.Get("Foobar"))
		assert.Equal(t, "gzip;q=0,deflate;q=0", proxyHeader.Get("Accept-Encoding"))
		assert.Equal(t, req.RemoteAddr, proxyHeader.Get("X-Forwarded-For"))
		assert.Equal(t, test.restURL, proxyReq.URL.String())

		if test.errFunc != nil {
			assert.Equal(t, test.response+"\n", string(data))
			assert.Equal(t, http.StatusInternalServerError, rwr.Code)
		} else {
			if test.errText != "" {
				assert.Equal(t, test.errText+"\n", string(data))
				assert.Equal(t, http.StatusBadGateway, rwr.Code)
			} else {
				assert.Equal(t, test.response, string(data), err)
				assert.Equal(t, http.StatusOK, rwr.Code)
			}
		}
	}
}

// TestProxyHandler tests Proxy constructor
func TestProxyHandler(t *testing.T) {
	proxyRoute := &ProxyRoute{
		Prefix: "/prefix", RedirectURL: "http://redirect",
	}
	handler := ProxyHandler(proxyRoute.Prefix, proxyRoute.RedirectURL)
	switch v := handler.(type) {
	case http.HandlerFunc:
		log.Println("ProxyHandler is an http.HandlerFunc")
	default:
		msg := fmt.Sprintf("%v is not http.HandlerFunc", v)
		assert.Fail(t, msg)
	}

	tests := []struct {
		requestURL   string
		expectedCode int
	}{
		{"", http.StatusBadRequest},
		{"http://test/foo", http.StatusBadGateway},
	}

	for idx, test := range tests {
		req, _ := http.NewRequest("GET", test.requestURL, nil)
		rwr := httptest.NewRecorder()

		log.Printf("Test %2d: %s\n", idx, test.requestURL)
		proxyRoute.ServeHTTP(rwr, req)
		handler.ServeHTTP(rwr, req)

		// log.Printf("Test %2d - request: %s - %#v\n", idx, req.URL, req)
		// log.Printf("Test %2d - response: %#v\n", idx, rwr)
		assert.Equal(t, test.expectedCode, rwr.Code)
	}
}
