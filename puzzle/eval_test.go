// +build all puzzle eval test

package puzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// EvalTestCase struct
type EvalTestCase struct {
	Expression string
	Expected   float64
}

// TestEval tests eval func
func TestEval(t *testing.T) {
	var tests = []EvalTestCase{
		{"1.1", 1.1},
		{"3 - 1", 2},
		{"123 - 4 * 5", 103},
		{"1.2 * 3.4 - 5.6 * 7.8 + 9", -30.6},
		{"33.3 / 3 * 4.5 - 6.7 + 8.9 / 10", 44.14},
		{"3.1415926535897932 * 1.414 / 2.0", 2.221106},
		{"1 ^ 2 + 3 ^ 4 - 5 ^ 6 * 7 - 8 * 9", -108873},
		{"-5.5 * -3.3 * -1.0", -18.15},
		{"- 9.0", -9},
	}

	for index, test := range tests {
		val, _ := eval(test.Expression)
		var msg = fmt.Sprintf("evaluating '%v' ==> %v", test.Expression, test.Expected)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.Equal(t, test.Expected, val, msg)
	}
}
