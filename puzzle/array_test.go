// +build all puzzle array parking test

package puzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ParkingTestCase struct
type ParkingSpotTestCase struct {
	inputs   []int
	expected int
}

// TestFindAvailableSpot tests func FindAvailableSpot
func TestFindAvailableSpot(t *testing.T) {
	var tests = []ParkingSpotTestCase{
		{[]int{0, 2, 7, 3, 1}, 4},
		{[]int{5, 0, 3, 17, 4, 1, 2, 13, 6, 8}, 7},
		{[]int{9, 3, 8, 7, 6, 5, 4, 19, 2, 1, 0}, 10},
	}
	for index, test := range tests {
		var val = FindAvailableSpot(test.inputs)
		var msg = fmt.Sprintf("expecting '%v' in %+v", test.expected, test.inputs)
		t.Logf("Test %2d: %v\n", index+1, msg)
		assert.Equal(t, test.expected, val, msg)
	}
}
