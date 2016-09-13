// +build all puzzle sum test

package puzzle

import (
	"fmt"
	"testing"

	u "github.com/dockerian/go-coding/utils"
	"github.com/stretchr/testify/assert"
)

// SumTestCase struct
type SumTestCase struct {
	ExpectedIndex1 int
	ExpectedIndex2 int
	Data           []int
	Sum            int
}

// TestFindMatchedSum tests FindMatchedSum function.
// Problems: Given an array of integers, find indices of the two numbers
//           such that they add up to a specific sum
// Keywords: array, hash, sum
func TestFindMatchedSum(t *testing.T) {
	tests := []SumTestCase{
		{0, 1, []int{2, 7, 3, 6, 9, 0, -2, 11}, 9},
		{0, 5, []int{3, 11, -3, 5, 0, 6, 12, 9, 4, -2, 11}, 9},
		{2, 3, []int{11, 0, -22, -1, -23, 314, 7, -11}, -23},
		{4, 8, []int{8, 7, 6, -1, -2, -3, 5, 0, -7, -9}, -9},
		{0, 5, []int{13, -7, -5, -3, -2, 0, 0, 0, 7, 6}, 13},
		{5, 6, []int{13, -7, -5, -3, -2, 0, 0, 0, 7, 6}, 0},
		{-1, -1, []int{9, 0, -22, -1, -23, 314, 7, -11}, 100},
	}

	for index, test := range tests {
		t.Logf("Finding sum=%v in %+v\n", test.Sum, test.Data)
		result1, result2 := FindMatchedSum(test.Data, test.Sum)
		expected := &u.Pair{Item1: test.ExpectedIndex1, Item2: test.ExpectedIndex2}
		pair := &u.Pair{Item1: result1, Item2: result2}
		result := pair.AreEqual(expected)
		var msg = fmt.Sprintf("expecting %v == %v", *pair, *expected)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.True(t, result, msg)
	}
}
