// +build all puzzle integer test

package puzzle

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Find2ndLargestTestCase struct
type Find2ndLargestTestCase struct {
	inputs   []int
	expected int
}

// TestFind2ndLargest tests func Find2ndLargest
func TestFind2ndLargest(t *testing.T) {
	var tests = []Find2ndLargestTestCase{
		{[]int{0, 2, 1, 3, 2}, 2},
		{[]int{3, 10, 2, 9, 18, 11, 5}, 11},
		{[]int{9, 9, 9}, 9},
	}
	for index, test := range tests {
		var val = Find2ndLargest(test.inputs)
		var msg = fmt.Sprintf("expecting '%v' in %+v", test.expected, test.inputs)
		t.Logf("Test %2d: %v\n", index+1, msg)
		assert.Equal(t, test.expected, val, msg)
	}
}

// TestTranslate tests Translate function
func TestTranslate(t *testing.T) {
	for idx, test := range []struct {
		number   uint64
		commaStr string
		expected string
	}{
		{
			0, "0", "zero",
		},
		{
			10, "10", "ten",
		},
		{
			100, "100", "one hundred",
		},
		{
			7000000, "7,000,000", "seven millions",
		},
		{
			39000000009, "39,000,000,009", "thirty nine billions nine",
		},
		{
			math.MaxUint32,
			"4,294,967,295",
			"four billions two hundreds ninety four millions nine hundreds sixty seven thousands two hundreds ninety five",
		},
		{
			math.MaxUint64,
			"18,446,744,073,709,551,615",
			"eighteen quintillions four hundreds forty six quadrillions seven hundreds forty four trillions seventy three billions seven hundreds nine millions five hundreds fifty one thousands six hundreds fifteen",
		},
	} {
		result := Translate(test.number)
		msg := fmt.Sprintf("Test %2d: %d ==> [%s] %s\n",
			idx, test.number, test.commaStr, test.expected)
		assert.Equal(t, test.expected, result, msg)
	}
}
