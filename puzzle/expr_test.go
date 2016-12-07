// +build all puzzle expr expression test

package puzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// CheckExpressionTestCase struct
type CheckExpressionTestCase struct {
	expression string
	expected   bool
}

// TestCheckExpression tests func CheckExpression
func TestCheckExpression(t *testing.T) {
	var tests = []CheckExpressionTestCase{
		{"(~_^) (/) (°,,°) (/) see more @ https://textfac.es/", true},
		{"(c) test case {3} :/[ ", false},
		{"b) test case {2}", false},
		{"big smile :=)", false},
		{"func (a, b) { return (a + 3) - b[4] }", true},
		{"hash := { key:'a', value:(1+(2*(3-4)/(5*(6+7/(8-9))))-foo[0]) }", true},
		{`regex: \((?:[^)(]+|\((?:[^)(]+|\([^)(]*\))*\))*\)`, false},
		{"regex = [^\\(]*(\\(.*\\))[^\\)]*", false},
		{"v := test // assignment #1", true},
		{"", true},
	}
	for index, test := range tests {
		var val = CheckExpression(test.expression)
		var msg = fmt.Sprintf("[%5v] '%v'", test.expected, test.expression)
		t.Logf("Test %2d: %v\n", index+1, msg)
		assert.Equal(t, test.expected, val, msg)
	}
}
