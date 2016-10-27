// +build all ds string str sort test

package str

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// SortByTestCase struct
type SortByTestCase struct {
	Data     []string
	Expected []string
}

// TestSortByCaseInsensitive tests sorting strings by case-insensitivity
func TestSortByCaseInsensitive(t *testing.T) {
	var tests = []SortByTestCase{
		{[]string{"ABC", "Zhu", "Beethoven", "Docker", "Alaska", "abc", "ABBA"},
			[]string{"ABBA", "abc", "ABC", "Alaska", "Beethoven", "Docker", "Zhu"}},
		{[]string{"abc", "ABC", "abc", "abC", "Abc", "aBC"},
			[]string{"abc", "ABC", "abc", "abC", "Abc", "aBC"}},
		{[]string{""}, []string{""}},
	}

	for index, test := range tests {
		subTests := []func([]string){
			ByCaseInsensitive,
			ByCaseInsensitivity.Sort,
		}

		for n, testFunc := range subTests {
			var val = make([]string, len(test.Data))
			copy(val, test.Data)
			var msg = fmt.Sprintf("expecting %+v == %+v", val, test.Expected)
			testFunc(val)
			t.Logf("Test %2d.%d: %v\n", index+1, n, msg)
			assert.EqualValues(t, test.Expected, val, msg)
		}
	}
}

// TestSortByLength tests sorting strings by length
func TestSortByLength(t *testing.T) {
	var tests = []SortByTestCase{
		{[]string{"ABC", "a", "Bee", "Docker", "Alaska", "abc", "ABBA", "Zhu"},
			[]string{"a", "ABC", "Bee", "abc", "Zhu", "ABBA", "Docker", "Alaska"}},
		{[]string{""}, []string{""}},
	}

	for index, test := range tests {
		var val = make([]string, len(test.Data))
		copy(val, test.Data)
		var msg = fmt.Sprintf("expecting %+v == %+v", val, test.Expected)
		ByLength.Sort(val)
		t.Logf("Test %2d: %v\n", index+1, msg)
		assert.EqualValues(t, test.Expected, val, msg)
	}
}
