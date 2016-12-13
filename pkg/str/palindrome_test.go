// +build all str palindrome test

package str

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	phaseFile = "palindrome_test_phase.json"
	vocabFile = "palindrome_test_voc.json"
)

// PalindromePhaseTestCase struct
type PalindromePhaseTestCase struct {
	Data     string `json:"data,omitempty"`
	HasPhase bool   `json:"test,omitempty"`
	Comment  string `json:"note,omitempty"`
}

// PalindromeTestCase struct
type PalindromeTestCase struct {
	Data     string
	Expected string
}

// TestPalindrome tests palindrome
func TestPalindrome(t *testing.T) {
	// Testing palindrome numbers
	negativeInt := -1111111111
	testPalindromeNumber(t, 12345, false)
	testPalindromeNumber(t, 9245224529, false)
	testPalindromeNumber(t, 543212345, true)
	testPalindromeNumber(t, 33333333333333, true)
	testPalindromeNumber(t, uint64(negativeInt), false)
	testPalindromeNumber(t, 0, true)

	// Testing palindrome strings, read from test file
	file, err1 := os.Open(vocabFile)
	if err1 != nil {
		t.Errorf("Cannot read test file (%v): %v\n", vocabFile, err1)
		t.Fail()
	}
	defer file.Close()

	keys := make(map[string]bool)
	decoder := json.NewDecoder(file)
	if err2 := decoder.Decode(&keys); err2 != nil {
		t.Errorf("Cannot decode test file (%v): %v\n", vocabFile, err2)
		t.Fail()
	}

	tests := make([]PalindromePhaseTestCase, len(keys))

	i := 0
	for key, val := range keys {
		test := PalindromePhaseTestCase{Data: key, HasPhase: val}
		// setting by index instead of using `append` to initial zero-size slice
		tests[i] = test
		i++
	}

	for _, test := range tests {
		testPalindromeString(t, strings.ToLower(test.Data), test.HasPhase)
	}
}

// TestPalindromePhase tests if a string is a palindrome
// of digits and letters (ignoring spaces and symbols)
func TestPalindromePhase(t *testing.T) {
	tests := []PalindromePhaseTestCase{}

	data, err := ioutil.ReadFile(phaseFile)
	if err != nil {
		t.Errorf("Cannot read test file (%v): %v\n", phaseFile, err)
		t.Fail()
	}

	err = json.Unmarshal(data, &tests)
	if err != nil {
		t.Errorf("Cannot parse test file (%v): %v\n", phaseFile, err)
		t.Fail()
	}

	for index, test := range tests {
		var pal = PalindromeString{input: test.Data}
		var val = pal.IsPalindromePhase()
		var msg = fmt.Sprintf("expecting palindrome phase '%v' ? %v", test.Data, test.HasPhase)
		t.Logf("Test %03d: %v\n", index+1, msg)
		assert.Equal(t, test.HasPhase, val, msg)
	}
}

// TestPalindromeSubstring tests GetPalindromicSubstring function
// See: https://leetcode.com/problems/longest-palindromic-substring/
func TestPalindromeSubstring(t *testing.T) {
	tests := []PalindromeTestCase{
		{"afdjfjdfdjfj", "jfjdfdjfj"},
		{"abc112123xx1xx3211cba", "123xx1xx321"},
		{"abc", "a"},
		{"bbq", "bb"},
		{"vvvv", "vvvv"},
		{"bbb", "bbb"},
		{"", ""},
	}

	for index, test := range tests {
		var val = GetPalindromicSubstring(test.Data)
		var msg = fmt.Sprintf("expecting palindrome '%v' in '%v'", test.Expected, test.Data)
		var pal = PalindromeString{input: test.Data}
		var sub = PalindromeString{input: val}
		var foo = test.Data == val
		var ms1 = fmt.Sprintf("expecting '%v' is palindrome ? %v", pal, foo)
		var ms2 = fmt.Sprintf("expecting '%v' is palindrome", sub)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.Equal(t, test.Expected, val, msg)
		assert.Equal(t, pal.IsPalindrome(), foo, ms1)
		assert.True(t, sub.IsPalindrome(), ms2)
	}
}

// testPalindrome tests palindrome for any input
func testPalindrome(t *testing.T, input Palindrome, expected bool) {
	result := input.IsPalindrome()
	t.Logf("%v:\tpalindrome ? %5v (expected %5v)\n", input.GetData(), result, expected)
	msg := fmt.Sprintf("%v : palindrome ? %v\n", input, expected)
	assert.Equal(t, expected, result, msg)
}

func testPalindromeNumber(t *testing.T, data uint64, expected bool) {
	testPalindrome(t, &PalindromeNumber{input: data}, expected)
}

func testPalindromeString(t *testing.T, data string, expected bool) {
	testPalindrome(t, &PalindromeString{input: data}, expected)
}
