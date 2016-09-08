// +build all puzzle sequence test

package puzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SequenceTestCase struct {
	Expected int
	Data     []int
}

// TestLongestConsecutiveIncrease tests getLongestConsecutiveIncrease
func TestLongestConsecutiveIncrease(t *testing.T) {
	tests := []SequenceTestCase{
		{3, []int{10, 9, 2, 5, 3, 7, 101, 18}},
		{4, []int{1, 2, 3, 4}},
		{3, []int{1, 2, 3, 0}},
		{2, []int{-1, 0, 0, 3, -10, 11}},
		{0, []int{9, 9, 9}},
	}

	for index, v := range tests {
		var actual, slice = getLongestConsecutiveIncrease(v.Data)
		var msg = fmt.Sprintf("expecting %v from %+v = %+v", v.Expected, v.Data, slice)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.Equal(t, v.Expected, actual, msg)
	}
}

// TestLongestIncrease tests getLongestIncrease
func TestLongestIncrease(t *testing.T) {
	tests := []SequenceTestCase{
		{4, []int{10, 9, 2, 5, 3, 7, 101, 18}},
		{4, []int{1, 2, 3, 4}},
		{5, []int{1, 2, 3, 0, 9, 99}},
		{9, []int{-11, -10, 0, -15, -14, -12, -17, -11, 0, -9, -1, 0, 3, -10, 11}},
		{6, []int{-11, 0, -18, -12, -11, -10, -21, 0, -17, -16, -15, -14, 0}},
		{0, []int{-7, -7, -7}},
	}

	for index, v := range tests {
		var actual, slice = getLongestIncrease(v.Data)
		var msg = fmt.Sprintf("expecting %v from %+v = %+v", v.Expected, v.Data, slice)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.Equal(t, v.Expected, actual, msg)
	}
}
