// +build all interview justify test

package interview

//----------------------------
// file: justify_test.go
//----------------------------

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// JustifyTestCase struct
type JustifyTestCase struct {
	Length   int
	Line     string
	Expected string
}

// getRuler returns a ruler string per length
func getRuler(length int) string {
	ruler := make([]byte, length)
	rules := make([]byte, length)
	for i := 1; i <= length; i++ {
		mark := i % 10
		if mark == 0 {
			ruler[i-1] = '0' + byte(i/10%10)
		} else {
			ruler[i-1] = '-'
		}
		rules[i-1] = '0' + byte(i%10)
	}
	return fmt.Sprintf("%s\n%s", string(ruler), rules)
}

// TestJustify tests func justfiy
func TestJustify(t *testing.T) {
	tests := []JustifyTestCase{
		{20, "abc", "abc                 "},
		{30, "this is a test.", "this      is      a      test."},
		{37, "Here is another test.", "Here       is      another      test."},
		{52, "The quick brown fox jumps over the lazy dog.", "The  quick  brown  fox  jumps  over  the  lazy  dog."},
		{20, "This line exceeds the length.", "This line exceeds the length."},
		{30, "The line length is exactly 30.", "The line length is exactly 30."},
		{120,
			"This one line will be justified in a 120 width column.",
			"This        one        line        will        be        justified        in       a       120       width       column."},
		{80, "", ""},
	}

	for index, test := range tests {
		var str = strings.Trim(test.Line, " ")
		var out = justify(str, test.Length)
		var msg = fmt.Sprintf("\n%v ==> [%v] \n%v\n%v",
			test.Line, test.Length, test.Expected, getRuler(test.Length))
		t.Logf("Test %d: %v\n", index+1, msg)
		var siz = len(str)
		if siz > 0 && siz <= test.Length {
			assert.Equal(t, test.Length, len(out), msg)
			assert.Equal(t, out, test.Expected, msg)
		} else {
			assert.Equal(t, out, test.Line, msg)
		}
	}
}
