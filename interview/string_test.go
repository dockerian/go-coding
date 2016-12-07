// +build all interview string substr occurrence test

package interview

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type FindMostOccurrencesTestCase struct {
	Data           string
	SubstrLen      int
	SubstrExpected string
	Counts         int
}

// TestFindMostOccurrences tests FindMostOccurrences
func TestFindMostOccurrences(t *testing.T) {
	tests := []FindMostOccurrencesTestCase{
		{"abcde", 4, "abcd", 1},
		{"abcdabc", 4, "abc", 2},
		{"micasaestucasasenor", 5, "as", 3},
		{"", 3, "", 0},
	}

	for index, v := range tests {
		var val, counts = FindMostOccurrences(v.Data, v.SubstrLen)
		var msg = fmt.Sprintf("substr in '%s': [%2d] '%s' (might be '%s')",
			v.Data, v.Counts, v.SubstrExpected, val)
		t.Logf("Test %v: %v\n", index+1, msg)
		// assert.Equal(t, v.SubstrExpected, val, msg) // hash has no guaranteed order
		assert.Equal(t, v.Counts, counts, msg)
	}
}
