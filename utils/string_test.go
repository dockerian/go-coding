// +build all utils str

// Package utils: utils/str_test.go

package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// defines the spec of how to search prefix and replace with proxyURL
	_replaceProxyTestCases = []ReplaceProxyTestCase{
		{"", "", "local:9001", ""},
		{"", "/redirect/", "local:9001", ""},
		{"", "/redirect/path", "local:9001", ""},
		{"host:80", "", "local:9001", "host:80"},
		{"host:80/test", "", "local:9001", "host:80/test"},
		{"host:80/test", "/redir/", "local:9001", "host:80/test"},
		{"host:80/redirpath", "/redir", "local:9001", "host:80/redirpath"},
		{"host:80/redirpath", "/redir/", "local:9001", "host:80/redirpath"},
		{"host:80/test/path", "/redir/", "local:9001", "host:80/test/path"},
		{"host:80/test/redir", "/redir", "local:9001", "local:9001"},
		{"host:80/test/redir", "/redir/", "local:9001", "local:9001"},
		{"host:80/redir", "/redir", "local:9001", "local:9001"},
		{"host:80/redir", "/redir/", "local:9001", "local:9001"},
		{"host:80/redir/", "/redir", "local:9001", "local:9001/"},
		{"host:80/redir/", "/redir/", "local:9001", "local:9001/"},
		{"host:80/redir/path", "/redir/", "local:9001", "local:9001/path"},
		{"host:80/redir/path/", "/redir/", "local:9001", "local:9001/path/"},
		{"host:80/redir/path/", "/redir", "local:9001", "local:9001/path/"},
		{},
	}
	// defines the spec of how to search string in a list
	_stringInTestCases = []StringInTestCase{
		{"", []string{""}, false, false},
		{"", []string{"a", "b"}, false, false},
		{"xyz", []string{"", "", ""}, false, false},
		{"xyz", []string{"abc", "def"}, false, false},
		{"xyz", []string{"abc", "XYZ"}, true, false},
		{"xyz", []string{"abc", "abcdefghijklmnopqrstuvwxyz"}, false, false},
		{"123", []string{"123", "abcdefghijklmnopqrstuvwxyz"}, true, true},
		{"123456", []string{"123", "123abc"}, false, false},
	}
)

// ReplaceProxyTestCaae struct
type ReplaceProxyTestCase struct {
	originalURL string
	prefix      string
	proxyURL    string
	expected    string
}

type StringInTestCase struct {
	search               string
	stringList           []string
	expectedOnIgnoreCase bool
	expected             bool
}

func TestReplaceProxyURL(t *testing.T) {
	for index, test := range _replaceProxyTestCases {
		result := ReplaceProxyURL(test.originalURL, test.prefix, test.proxyURL)
		msg := fmt.Sprintf("search: '%s' in %v --> result: %v (expected: %v) - proxyURL: %v",
			test.prefix, test.originalURL, result, test.expected, test.proxyURL)
		t.Logf("Test %2d: %v\n", index+1, msg)
		assert.Equal(t, test.expected, result, msg)
	}
}

// TestStringIn tests common.StringIn function
func TestStringIn(t *testing.T) {
	for index, test := range _stringInTestCases {
		result := StringIn(test.search, test.stringList)
		resultIgnoreCase := StringIn(test.search, test.stringList, true)
		msg := fmt.Sprintf("search: '%s' in %v -->\nresult: %t (expected: %t) | ignore-case: %t (%t)",
			test.search, test.stringList, result, test.expected,
			resultIgnoreCase, test.expectedOnIgnoreCase)
		t.Logf("Test %2d: \n%v\n", index+1, msg)
		assert.Equal(t, test.expectedOnIgnoreCase, resultIgnoreCase, msg)
		assert.Equal(t, test.expected, result, msg)
	}
}
