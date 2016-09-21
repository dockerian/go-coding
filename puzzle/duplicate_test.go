// +build all puzzle duplicate test

package puzzle

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// DuplicateTestCase struct
type DuplicateTestCase struct {
	inputs   []int
	expected int
}

// DuplicatesTestCase struct
type DuplicatesTestCase struct {
	inputs        []int
	dups          []int
	canRaiseError bool
}

// TestFindDuplicate tests func FindDuplicate
func TestFindDuplicate(t *testing.T) {
	var tests = []DuplicateTestCase{
		{[]int{0, 2, 1, 3, 2}, 2},
		{[]int{5, 0, 3, 7, 4, 1, 2, 3, 6, 8}, 3},
		{[]int{9, 3, 8, 7, 6, 5, 4, 9, 2, 1, 0}, 9},
	}
	for index, test := range tests {
		var val = FindDuplicate(test.inputs)
		var msg = fmt.Sprintf("expecting '%v' in %+v", test.expected, test.inputs)
		t.Logf("Test %2d: %v\n", index+1, msg)
		assert.Equal(t, test.expected, val, msg)
	}
}

// TestFindDuplicates tests func FindDuplicates
func TestFindDuplicates(t *testing.T) {
	var tests = []DuplicatesTestCase{
		{[]int{0, 1, 2, 3, 4}, []int{}, false},
		{[]int{0, 2, 1, 3, 2}, []int{2}, false},
		{[]int{6, 5, 4, 2, 3, 5, 5}, []int{5, 5}, false},
		{[]int{5, 0, 3, 7, 4, 0, 2, 3, 6, 8}, []int{0, 3}, false},
		{[]int{7, 1, 2, 1, 4, 2, 6, 7}, []int{1, 2, 7}, false},
		{[]int{7, 1, 2, 1, 4, 2, 6, 17}, []int{}, true},
		{[]int{0, 1, 1, 3, 3, 5, 6, 7, 9, 9}, []int{1, 3, 9}, false},
		{[]int{4, 4, 3, 3, 1, 1}, []int{1, 3, 4}, false},
	}
	for index, test := range tests {
		var result, err = FindDuplicates(test.inputs)
		var msg = fmt.Sprintf("expecting %+v in %+v", test.dups, test.inputs)
		t.Logf("Test %2d: %v\n", index+1, msg)
		if err != nil {
			if !test.canRaiseError {
				t.Errorf("unexpected error %v\n", err)
			}
		} else {
			sort.Ints(result)
			sort.Ints(test.dups)
			assert.Equal(t, test.dups, result, msg)
		}
	}
}
