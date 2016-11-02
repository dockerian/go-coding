// +build all puzzle integer test

package puzzle

import (
	"fmt"
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
