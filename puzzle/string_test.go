// +build all puzzle string test

package puzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StringTestCase struct {
	Expected string
	Data     string
	Start    int
	End      int
	Len      int
}

// TestGetLongestSubstringLength tests GetLongestSubstringLength
func TestGetLongestSubstringLength(t *testing.T) {
	// NOTE: GetLongestSubstringLength does not work with UTF-8 string
	tests := []StringTestCase{
		{"abc", "--##aaabcabcbb", 6, 9, 3},
		{"mi cas", "casa mi casa es tu casa", 5, 11, 6},
		{"o", "ooooooooo", 0, 1, 1},
	}

	for index, test := range tests {
		var len = len(test.Expected)
		var val, slice = GetLongestSubstringLength(test.Data)
		var msg = fmt.Sprintf("expecting longest '%v' [%v] in '%v'", slice, len, test.Data)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.Equal(t, len, val, msg)
	}
}

// TestGetLongestSubstringUTF8 tests GetLongestSubstringUTF8
func TestGetLongestSubstringUTF8(t *testing.T) {
	tests := []StringTestCase{
		{"abc", "--##aaabcabcbb", 6, 9, 3},
		{"mi cas", "casa mi casa es tu casa", 5, 11, 6},
		{"一定不吐葡萄皮", "吃葡萄不吐葡萄皮不吃葡萄不一定不吐葡萄皮吐葡萄籽儿", 39, 60, 7},
		{"o", "ooooooooo", 0, 1, 1},
	}

	for index, test := range tests {
		var val, str, indexL, indexR = GetLongestSubstringUTF8(test.Data)
		var msg = fmt.Sprintf("expecting longest '%v' [%v] in '%v'",
			test.Expected, test.Len, test.Data)
		t.Logf("Test %v: %v [%v : %v]\n", index+1, msg, indexL, indexR)
		assert.Equal(t, test.Len, val, msg)
		assert.Equal(t, test.Expected, str, msg)
		assert.Equal(t, test.Start, indexL, msg)
		assert.Equal(t, test.End, indexR, msg)
		assert.True(t, test.Len <= indexR-indexL, msg)
	}
}

// TestGetLongestUniqueSubstring tests GetLongestUniqueSubstring
func TestGetLongestUniqueSubstring(t *testing.T) {
	tests := []StringTestCase{
		{"abc", "--##aaabcabcbb", 6, 9, 3},
		{"mi cas", "casa mi casa es tu casa", 5, 11, 6},
		{"一定不吐葡萄皮", "吃葡萄不吐葡萄皮不吃葡萄不一定不吐葡萄皮吐葡萄籽儿", 39, 60, 7},
		{"o", "ooooooooo", 0, 1, 1},
	}

	for index, test := range tests {
		var val = GetLongestUniqueSubstring(test.Data)
		var msg = fmt.Sprintf("expecting '%v' from '%v'", test.Expected, test.Data)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.Equal(t, test.Expected, val, msg)
	}
}
