// +build all puzzle palindrome test

package puzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPalindrome tests palindrome
func TestPalindrome(t *testing.T) {
	negativeInt := -1111111111
	testPalindromeNumber(t, 12345, false)
	testPalindromeNumber(t, 9245224529, false)
	testPalindromeNumber(t, 543212345, true)
	testPalindromeNumber(t, 33333333333333, true)
	testPalindromeNumber(t, uint64(negativeInt), false)
	testPalindromeNumber(t, 0, true)
	testPalindromeString(t, "input", false)
	testPalindromeString(t, "abba", true)
	testPalindromeString(t, "ZZZZZZZZZ", true)
	testPalindromeString(t, "A", true)
}

// testPalindrome tests palindrome for any input
func testPalindrome(t *testing.T, input Palindrome, expected bool) {
	result := input.IsPalindrome()
	t.Logf("%v:\tpalindrome ? %5v (expected %5v)\n", input.GetData(), result, expected)
	assert.Equal(t, expected, result, fmt.Sprintf("%v : palindrome ? %v\n", input, expected))
}

func testPalindromeNumber(t *testing.T, data uint64, expected bool) {
	testPalindrome(t, &PalindromeNumber{input: data}, expected)
}

func testPalindromeString(t *testing.T, data string, expected bool) {
	testPalindrome(t, &PalindromeString{input: data}, expected)
}
