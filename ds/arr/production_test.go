// +build all ds array production

package arr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetProducts is a testing function GetProducts
func TestGetProducts(t *testing.T) {
	var tests = []struct {
		numbers  []int
		expected []int
	}{
		{
			[]int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24},
		},
		{
			[]int{3, 4, 1, 5, 2}, []int{40, 30, 120, 24, 60},
		},
	}

	for index, test := range tests {
		var msg = fmt.Sprintf("expecting %v", test.expected)
		var result = GetProducts(test.numbers)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.EqualValues(t, result, test.expected, msg)
	}
}
