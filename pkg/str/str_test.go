// +build all common pkg str

// Package str :: str_test.go

package str

import (
	"fmt"
	"strings"
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
		{"http://host:80", "", "local:9001", "http://host:80"},
		{"http://host:80/test", "", "local:9001", "http://host:80/test"},
		{"http://host/prefix/test", "/foo", "local", "http://host/prefix/test"},
		{"http://host:80/prefix/test", "/prefix", "local:9001", "local:9001/test"},
		{"http://host:80/prefix", "/prefix", "local:9001", "local:9001"},
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

// StringInTestCase struct
type StringInTestCase struct {
	search               string
	stringList           []string
	expectedOnIgnoreCase bool
	expected             bool
}

// BenchmarkAppend
func BenchmarkAppend(b *testing.B) {
	var str string
	str1 := "slice1...."
	str2 := "slice2...."
	result := make([]byte, 0, 0)
	slice1 := []byte(str1)
	slice2 := []byte(str2)

	str = "Benchmark_Append_function"
	b.Run(str, func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result = Append(result, slice1)
			result = Append(result, slice2)
		}
	})

	str = "Benchmark_build-in-append"
	b.Run(str, func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result = append(result, slice1...)
			result = append(result, slice2...)
		}
	})

	str = "Benchmark_string_concat"
	b.Run(str, func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			str += str1
			str += str2
		}
	})
	str = "Benchmark pkg/str"
	fmt.Printf("DONE: %s\n", str)
}

// TestAppend
func TestAppend(t *testing.T) {
	slice1 := []byte(strings.Repeat("slice1....", 10))
	slice2 := []byte(strings.Repeat("slice2....", 10))
	cap1 := cap(slice1)
	cap2 := cap(slice2)

	result := Append(slice1, slice2)
	cap3 := cap(result)
	t.Logf("Test: appending '%s' [cap: %d] + '%s' [%d] => '%s' [cap: %d]\n",
		string(slice1), cap1, string(slice2), cap2, string(result), cap3)
	assert.True(t, (cap1+cap2) <= cap3)
}

// TestIndentJSON
func TestIndentJSON(t *testing.T) {
	ch1 := make(chan int)
	obj := struct {
		A string `json:"a"`
		B string `json:"b"`
		C int    `json:"count"`
	}{
		"aaa", "bb", 3,
	}
	out := `{
    "a": "aaa",
    "b": "bb",
    "count": 3
}
`
	tests := []struct {
		input    interface{}
		indent   string
		expected string
	}{
		{input: ch1, indent: "", expected: ""},
		{input: obj, indent: "    ", expected: out},
		{input: "test", indent: "", expected: "\"test\"\n"},
	}
	for idx, test := range tests {
		result := IndentJSON(test.input, test.indent)
		t.Logf("Test %2d: %+v => %s\n", idx+1, test.input, test.expected)
		assert.Equal(t, test.expected, result)
	}
}

// TestReplaceProxyURL
func TestReplaceProxyURL(t *testing.T) {
	for index, test := range _replaceProxyTestCases {
		result := ReplaceProxyURL(test.originalURL, test.prefix, test.proxyURL)
		msg := fmt.Sprintf("search: '%s' in '%v' --> result: %v (expected: %v) - proxyURL: %v",
			test.prefix, test.originalURL, result, test.expected, test.proxyURL)
		t.Logf("Test %2d: '%s' in '%v'\n", index+1, test.prefix, test.originalURL)
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
