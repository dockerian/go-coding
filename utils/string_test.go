// +build all utils str

// Package utils: utils/str_test.go

package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
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

type StringInTestCase struct {
	search               string
	stringList           []string
	expectedOnIgnoreCase bool
	expected             bool
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
