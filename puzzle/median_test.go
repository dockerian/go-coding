// +build all puzzle median test

package puzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// FindMedianTestCase struct
type FindMedianTestCase struct {
	inputs   []int
	expected float64
}

// TestFindMedian tests func FindMedian
// See https://leetcode.com/problems/find-median-from-data-stream/
func TestFindMedian(t *testing.T) {
	var tests = []FindMedianTestCase{
		{[]int{0, 2, 1, 4, 3, 6, 5}, 3.0},
		{[]int{3, 10, 2, 9, 18, 11, 5, 101, 3, -9, 0}, 5.0},
		{[]int{19, -11, 0, 7, 33, -1024, 0, 11}, 3.5},
		{[]int{9, 9, 9}, 9},
	}
	for index, test := range tests {
		var val = FindMedian(test.inputs)
		var msg = fmt.Sprintf("median in %+v: %+v", test.inputs, test.expected)
		t.Logf("Test %2d: %v\n", index+1, msg)
		assert.Equal(t, test.expected, val, msg)
	}
}
