// +build all ds math gauss test

package mathex

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGaussSum
func TestGaussSum(t *testing.T) {
	for index, test := range []struct {
		Start    uint16
		Interval uint16
		Num      uint16
	}{
		{1, 1, 99},
		{3, 2, 101},
		{9005, 1000, 10001},
		{math.MaxUint16, math.MaxUint16, math.MaxUint16},
		{math.MaxUint16, 1, math.MaxUint16},
	} {
		result := GaussSum(test.Start, test.Interval, test.Num)
		expected := uint64(test.Start)
		x := uint64(test.Start)
		v := uint64(test.Interval)
		n := uint64(test.Num)
		if test.Num > 0 {
			for i := uint64(1); i <= n; i++ {
				expected += x + v*i
			}
		}
		var msg = fmt.Sprintf("sum of [%d, %d, ..., %d] => %d",
			x, x+v, x+(n*v), expected)
		t.Logf("Test %2d: %v\n", index+1, msg)
		assert.Equal(t, expected, result, msg)
	}
}
