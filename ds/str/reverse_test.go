// +build all ds str test

package str

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	funcReverses = []ReverseFunc{
		{"Default", Reverse},
		{"Runes", reverseRunes},
		{"ByRunes", reverseByRunes},
		{"ByCopy", reverseByCopy},
		{"BySprintf", reverseBySprintf},
		{"Func", reverseFunc},
		{"Grapheme", reverseGrapheme},
		{"Norm", reverseNorm},
		{"Unicode", reverseGraphemeUnicode},
		{"UTF8", reverseUTF8},
	}
)

type ReverseFunc struct {
	Name string
	Func func(string) string
}

// ReverseStringTestCase struct
type ReverseStringTestCase struct {
	Data     string
	Expected string
}

// StringTestCase struct
type StringTestCase struct {
	Data     interface{}
	Expected interface{}
}

// generateTestString returns a specified string s*n
func generateTestString(s string, n int) string {
	test := ""
	for i := 0; i < n; i++ {
		test += s
	}
	return test
}

// TestReverseString tests reversing string functions
func TestReverseString(t *testing.T) {
	var tests = []ReverseStringTestCase{
		{"", ""},
		{"AbCdEfgHiJ", "JiHgfEdCbA"},
		{"帘卷晚晴天", "天晴晚卷帘"},
		{"solos:…花枝弄影照窗纱，影照窗纱映日斜……", "……斜日映纱窗照影，纱窗照影弄枝花…:solos"},
		{"deified civic", "civic deified"},
		{"radar tenet", "tenet radar"},
	}

	for index, test := range tests {
		for idx, exec := range funcReverses {
			var val = exec.Func(test.Data)
			var tst = exec.Func(val)
			var msg = fmt.Sprintf("'%v'<=>'%v'", tst, test.Expected)
			t.Logf("Test %2d.%02v [%9s]: %v\n", index+1, idx+1, exec.Name, msg)
			assert.Equal(t, test.Expected, val, msg)
			assert.Equal(t, test.Data, tst, msg)
		}
	}
}

// TestString is a testing function template
func TestString(t *testing.T) {
	var tests []StringTestCase

	for index, test := range tests {
		var val interface{}
		var msg = fmt.Sprintf("expecting %v == %v", val, test.Expected)
		t.Logf("Test %v: %v\n", index+1, msg)
	}
}
