// +build all ds str reverse test

package str

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ReverseWordsTestCase struct
type ReverseWordsTestCase struct {
	Data     string
	Expected string
}

// TestReverseWords tests reversing string functions
func TestReverseWords(t *testing.T) {
	var tests = []ReverseWordsTestCase{
		{"", ""},
		{"mine is yours", "yours is mine"},
		{"帘 卷 晚 晴 天", "天 晴 晚 卷 帘"},
		{" 照影  窗纱  红 ", "红 窗纱 照影"},
		{"palindrome deified civic", "civic deified palindrome"},
		{"radar tenet", "tenet radar"},
	}

	for index, test := range tests {
		var val = ReverseWords(test.Data)
		var msg = fmt.Sprintf("'%v'=>'%v'", test.Data, test.Expected)
		t.Logf("Test %2d: %v\n", index+1, msg)
		assert.Equal(t, test.Expected, val, msg)
	}
}
