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

var (
	bodyReadErr = func(r io.Reader) ([]byte, error) {
		return []byte{}, errors.New("body read error")
	}
	bodyReadOkay = func(r io.Reader) ([]byte, error) {
		return []byte{}, nil
	}

	clientResponse200 = &MockProxyClient{
		StatusCode:   http.StatusOK,
		ResponseText: "{data: {}}",
		Err:          nil,
	}
	clientResponse404 = &MockProxyClient{
		StatusCode:   http.StatusNotFound,
		ResponseText: fmt.Sprintf("{code: %d}", http.StatusNotFound),
		Err:          nil,
	}
	clientResponseErr = &MockProxyClient{
		StatusCode:   http.StatusForbidden,
		ResponseText: fmt.Sprintf("{err: \"%s\"}", http.StatusText(http.StatusForbidden)),
		Err:          errors.New(http.StatusText(http.StatusForbidden)),
	}

	ioCopyErrorFunc = func(dst io.Writer, src io.Reader) (int64, error) {
		return 0, errors.New("mocked io.Copy error")
	}
	ioCopyOkayFunc = func(dst io.Writer, src io.Reader) (int64, error) {
		return 1024, nil
	}
)

// MockProxyClient struct
type MockProxyClient struct {
	StatusCode   int
	Request      *http.Request
	ResponseText string
	Err          error
}

// Do implements ProxyClient interface
func (mpc *MockProxyClient) Do(r *http.Request) (*http.Response, error) {
	mpc.Request = r
	response := &http.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte(mpc.ResponseText))),
		Header: map[string][]string{
			"Content-Type": {
				"application/json",
			},
		},
	}
	response.Status = http.StatusText(mpc.StatusCode)
	response.StatusCode = mpc.StatusCode
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
		errFunc     func(io.Writer, io.Reader) (int64, error)
		errCode     int
	}{
		{
			"/no-prefix", "http://redir", "", "",
			"", "", nil, http.StatusBadRequest,
		},
		{
			"/prefix1", "http://redir1", "/prefix1/test", "http://redir1/test",
			"good", "", nil,
			http.StatusOK,
		},
		{
			"/pre2", "http://re2", "/pre2/test/value", "http://re2/test/value",
			"BAD GATEWAY", "bad gateway err test", nil,
			http.StatusBadGateway,
		},
		{
			"/pre3", "http://re3", "/pre3/test/3", "http://re3/test/3",
			"io.Copy err test", "", ioCopyErrorFunc,
			http.StatusInternalServerError,
		},
	}

	savedIoCopy := ioCopy
	savedClient := proxyClient
	defer func() {
		// restoring ioCopy and proxyClient (see definition in proxy.go)
		ioCopy = savedIoCopy
		proxyClient = savedClient
	}()

	for idx, test := range tests {
		var err error
		if test.errText != "" {
			err = errors.New(test.errText)
		}
		var clientCode = http.StatusOK
		if test.errFunc != nil {
			clientCode = http.StatusNotAcceptable
			ioCopy = test.errFunc
		} else {
			ioCopy = savedIoCopy
		}
		client := &MockProxyClient{
			StatusCode:   clientCode,
			ResponseText: test.response,
			Err:          err,
		}

		req, _ := http.NewRequest("GET", test.reqURL, nil)
		req.Header.Set("Foobar", "Foobar header test")
		msg := fmt.Sprintf("%s => %s", test.reqURL, test.restURL)
		t.Logf("Test %2d: %s\n", idx+1, msg)
		proxyClient = client
		rwr := httptest.NewRecorder()
		// here start to test the function
		Proxy(test.prefix, test.redirectURL, rwr, req)

		if test.reqURL == "" {
			assert.Equal(t, http.StatusBadRequest, rwr.Code)
			continue
		}

		data, err := ioutil.ReadAll(rwr.Body)
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
			assert.Contains(t, string(data), "mocked io.Copy error")
			assert.Equal(t, http.StatusNotAcceptable, rwr.Code)
		} else {
			if test.errText != "" {
				assert.Contains(t, string(data), test.errText)
				assert.Equal(t, http.StatusBadGateway, rwr.Code)
			} else if err != nil {
				assert.Contains(t, string(data), test.response)
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

	savedClient := proxyClient
	savedIoCopy := ioCopy
	defer func() {
		// restoring ioCopy and proxyClient (see definition in proxy.go)
		ioCopy = savedIoCopy
		proxyClient = savedClient
	}()

	tests := []struct {
		requestURL   string
		expectedCode int
		ioCopyMock   func(dst io.Writer, src io.Reader) (int64, error)
		clientMock   *MockProxyClient
	}{
		{"", http.StatusBadRequest, nil, nil},
		{"http://test1/foo", http.StatusBadGateway, ioCopyOkayFunc, clientResponseErr},
		{"http://test2/foo", http.StatusOK, ioCopyErrorFunc, clientResponse200},
		{"http://test3/foo", http.StatusNotFound, ioCopyOkayFunc, clientResponse404},
		{"http://test4/foo", http.StatusOK, ioCopyOkayFunc, clientResponse200},
	}

	for idx, test := range tests {
		proxyClient = test.clientMock
		ioCopy = test.ioCopyMock
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
