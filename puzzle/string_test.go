// +build all puzzle string test

package puzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type RuneTestCase struct {
	Data     string
	Expected rune
	Count    int
}

type StringPairTestCase struct {
	Data     string
	Begin    string
	Close    string
	IsPaired bool
}

type StringTestCase struct {
	Expected string
	Data     string
	Start    int
	End      int
	Len      int
}

// TestCheckMatchedPair tests CheckMatchedPair
func TestCheckMatchedPair(t *testing.T) {
	tests := []StringPairTestCase{
		{"", "", "", true},
		{"<h1>Hello</h1>", "<h1>", "</h1>", true},
		{"<p>你好<p>", "<p>", "</p>", false},
		{"for(var i=0;i<10;i++){\n\tconsole.log(i);\n}", "{", "}", true},
		{"So sad :( and so happy :)", "(", ")", true},
		{"Miss you in the heaven :)", "(", ")", false},
		{"Not okay *)(*", "(", ")", false},
	}

	for index, test := range tests {
		var val, err = CheckMatchedPair(test.Data, test.Begin, test.Close)
		var msg = fmt.Sprintf("expecting '%v' and '%v' matched in '%v' ? %v",
			test.Begin, test.Close, test.Data, test.IsPaired)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.Equal(t, test.IsPaired, val, msg)
		if !test.IsPaired {
			assert.NotNil(t, err, msg)
		}
	}
}

// TestGetLongestSubstringLength tests GetLongestSubstringLength
// See: https://leetcode.com/problems/longest-substring-without-repeating-characters/
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

// TestGetMostFrequentRune tests GetGetMostFrequentRune
func TestGetMostFrequentRune(t *testing.T) {
	tests := []RuneTestCase{
		{"--##aaabcabcbb", 'a', 4},
		{"吃葡萄不吐葡萄皮不吃葡萄不一定不吐葡萄皮吐葡萄籽儿", '葡', 5},
		{"寻寻觅觅，冷冷清清，凄凄惨惨戚戚。乍暖还寒时候，最难将息。", '，', 3},
		{"", '\x00', 0},
	}

	for index, test := range tests {
		var val, count = GetMostFrequentRune(test.Data)
		var msg = fmt.Sprintf("expecting ('%c', %v) from '%v'", test.Expected, test.Count, test.Data)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.Equal(t, test.Expected, val, msg)
		assert.Equal(t, test.Count, count, msg)
	}
}
