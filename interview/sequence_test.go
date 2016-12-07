// +build all interview sequence test

package interview

import (
	"fmt"
	"testing"

	u "github.com/dockerian/go-coding/utils"
	"github.com/stretchr/testify/assert"
)

type SequenceStringTestCase struct {
	Data      string
	Decending bool
	Sequence  string
}

type SequenceTestCase struct {
	Expected int
	Data     []int
}

type SumSequenceTestCase struct {
	Data     []int
	MaxLen   int // maximum length of sequence to test
	Expected []int
	Sum      int64
}

// TestLongestConsecutiveIncrease tests GetLongestConsecutiveIncrease
func TestLongestConsecutiveIncrease(t *testing.T) {
	tests := []SequenceTestCase{
		{3, []int{10, 9, 2, 5, 3, 7, 101, 18}},
		{4, []int{1, 2, 3, 4}},
		{3, []int{1, 2, 3, 0}},
		{2, []int{-1, 0, 0, 3, -10, 11}},
		{0, []int{9, 9, 9}},
	}

	for index, v := range tests {
		var actual, slice = GetLongestConsecutiveIncrease(v.Data)
		var msg = fmt.Sprintf("expecting %v from %+v = %+v", v.Expected, v.Data, slice)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.Equal(t, v.Expected, actual, msg)
	}
}

// TestLongestIncrease tests GetLongestIncrease
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
		var actual, slice = GetLongestIncrease(v.Data)
		var msg = fmt.Sprintf("expecting %v from %+v = %+v", v.Expected, v.Data, slice)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.Equal(t, v.Expected, actual, msg)
	}
}

// TestLongestSequence tests GetLongestSequence
func TestLongestSequence(t *testing.T) {
	tests := []SequenceStringTestCase{
		{"0123.abcdefg.456789", false, "abcdefg"},
		{"nothing-in-sequential-but-stuvw", false, "stuvw"},
		{"hijk-9876543210-tsrqpon", false, "hijk"},
		{"hijk-9876543210-tsrqpon", true, "9876543210"},
		{"zyx--", true, "zyx"},
		{"", true, ""},
	}

	for index, test := range tests {
		var val = GetLongestSequence(test.Data, test.Decending)
		var seq = u.GetTernary(test.Decending, "decending", "acending")
		var msg = fmt.Sprintf("finding longest %v sequence: '%v' in '%v'",
			seq, test.Sequence, test.Data)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.Equal(t, test.Sequence, val, msg)
	}
}

// TestMaxSumSequence tests GetMaxSumSequence
func TestMaxSumSequence(t *testing.T) {
	tests := []SumSequenceTestCase{
		{[]int{1, 2, 3, 4, -1, 2, 3, -1, 5}, 0, []int{1, 2, 3, 4, -1, 2, 3, -1, 5}, 18},
		{[]int{-1, 0, 1, 2, 3, 4, -1, 2, 3, -1, 5}, 4, []int{0, 1, 2, 3, 4}, 10},
		{[]int{1, 2, 3, 4, -10, 2, 3, -11}, 0, []int{1, 2, 3, 4}, 10},
		{[]int{}, 0, []int{}, 0},
	}

	for index, test := range tests {
		var sub, val = GetMaxSumSequence(test.Data, test.MaxLen)
		var msg = fmt.Sprintf("max sum in %+v: %+v => %d",
			test.Data, test.Expected, test.Sum)
		t.Logf("Test %v: %v\n", index+1, msg)
		var sum int64
		for _, n := range sub {
			sum += int64(n)
		}
		assert.Equal(t, test.Sum, val, msg)
		assert.Equal(t, test.Sum, sum, msg)
	}
}
