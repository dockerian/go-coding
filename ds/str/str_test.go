// +build all ds str compress test

package str

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CompressStringTestCase struct {
	Data     string
	Expected string
}

// TestCompressString tests CompressString
func TestCompressString(t *testing.T) {
	tests := []CompressStringTestCase{
		{"abc", "abc"},
		{"aabbccc", "aabbccc"},
		{"aaaaaaabccccccccccc", "a*7\\bc*11\\"},
		{"c:\\windows", "c:\\\\windows"},
		{"*********stars**********", "\\**9\\stars\\**10\\"},
		{"\\\\\\\\\\\\\\\\\\\\\\\\", "\\\\*12\\"},
		{"", ""},
	}

	for index, test := range tests {
		var str = Str(test.Data)
		var exp = Str(test.Expected)
		var out = str.Compress()
		var val = exp.Decompress()
		var ms1 = fmt.Sprintf("compress '%v' ==> '%v'", test.Data, test.Expected)
		var ms2 = fmt.Sprintf("decompress '%v' ==> '%v'", test.Expected, test.Data)
		t.Logf("Test %v: %v\n", index+1, ms1)
		assert.Equal(t, test.Expected, out, ms1)
		assert.Equal(t, test.Data, val, ms2)
	}
}
