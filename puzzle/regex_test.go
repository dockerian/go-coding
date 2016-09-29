// +build all puzzle regex test

package puzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// BasicRegexTestCase struct
type BasicRegexTestCase struct {
	input string
	regex string
	match bool
}

// TestIsMatchBasicRegex tests IsMatchBasicRegex function
func TestIsMatchBasicRegex(t *testing.T) {
	var tests = []BasicRegexTestCase{
		{"", "", true},
		{"", "*", false},
		{"", ".*", true},
		{"a", ".*..a*", false},
		{"a", "a*b*", true},
		{"a", "a*b*.", true},
		{"a", "a*b*c", false},
		{"aa", "a", false},
		{"aa", ".", false},
		{"aa", "aa", true},
		{"aa", "a*", true},
		{"aaa", "a*a", true},
		{"aaa", "ab*a*c*a", true},
		{"aaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*c", false},
		{"aab", "c*a*b", true},
		{"aab", ".*a*b", true},
		{"ab", ".*..", true},
		{"ab", "ab", true},
		{"ab", ".*", true},
		{"abdefgh", ".*", true},
		{"abc", "...", true},
		{"abc", "a*b*c*", true},
	}

	for index, test := range tests {
		var reg = BasicRegex{test.input, test.regex}
		var val = reg.IsMatch()
		var msg = fmt.Sprintf("expecting '%v' matches '%v' ? %v", test.input, test.regex, test.match)
		t.Logf("Test %2d: %v\n", index+1, msg)
		assert.Equal(t, test.match, val, fmt.Sprintf("%v [%v]", msg, "IsMatch.*DP"))

		var chk = isMatchBasicRegexRecursive(&test.input, &test.regex)
		assert.Equal(t, test.match, chk, fmt.Sprintf("%v [%v]", msg, "isMatch.*Recursive"))

		// TODO: so far the following isMatch methods won't pass all tests
		// var ck1 = isMatchBasicRegexNR(&test.input, &test.regex)
		// var ck2 = isMatchBasicRegexSlice(&test.input, &test.regex)
		// assert.Equal(t, test.match, ck1, fmt.Sprintf("%v [%v]", msg, "isMatch.*Non-recursive"))
		// assert.Equal(t, test.match, ck2, fmt.Sprintf("%v [%v]", msg, "isMatch.*Slice"))
	}
}
