// +build all ds exp test

package exp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ExpTestCase struct {
	Data       string
	Expression string
	Result     float64
}

// TestExp tests Exp functions
func TestExp(t *testing.T) {
	var tests = []ExpTestCase{
		{"1.2 + 34.5 + 56.7", "(1.2 + (34.5 + 56.7))", 92.4},
		{"7 + 15 / 5 * 6", "(7 + ((15 / 5) * 6))", 25.0},
		{"11*(55-5)-50*3", "((11 * (55 - 5)) - (50 * 3))", 400},
		{"50/5 + 3.33-8+(3*25)", "((50 / 5) + ((3.33 - 8) + (3 * 25)))", 80.33},
		{"", "", 0.0},
	}

	for index, test := range tests {
		var ctx = Exp{context: test.Data}
		var exp = ctx.String()
		var val = ctx.Eval()
		var msg = fmt.Sprintf("expecting '%v' => %v", test.Data, test.Expression)
		var out = fmt.Sprintf("expecting '%v' == %v", test.Data, test.Result)
		t.Logf("\nTest %v: %v = {%v}\n", index+1, msg, test.Result)
		assert.Equal(t, test.Expression, exp, msg)
		assert.Equal(t, test.Result, val, out)
	}
}
