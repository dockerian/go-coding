// +build all puzzle matrix test

package puzzle

import (
	"fmt"
	"testing"

	"github.com/dockerian/go-coding/ds/str"
	"github.com/stretchr/testify/assert"
)

// FindAjacent1sTestCase struct
type FindAjacent1sTestCase struct {
	inputs   [][]int
	expected []string
}

// TestFindAjacent1s tests func FindAjacent1s
func TestFindAjacent1s(t *testing.T) {
	var tests = []FindAjacent1sTestCase{
		{[][]int{
			{1, 1, 0, 1, 0, 1},
			{1, 1, 1, 1, 0, 1},
			{0, 0, 0, 0, 1, 1},
			{1, 0, 1, 0, 1, 0},
		},
			[]string{"0,0", "0,1", "0,3", "0,5", "1,0", "1,1", "1,2", "1,3", "1,5", "2,4", "2,5", "3,4"},
		},
	}
	for index, test := range tests {
		var val = FindAjacent1s(test.inputs)
		var msg = fmt.Sprintf("expecting %v in %+v", test.expected, test.inputs)
		str.ByCaseInsensitive(val)
		t.Logf("Test %2d: %v\n", index+1, msg)
		assert.Equal(t, test.expected, val, msg)
	}
}
