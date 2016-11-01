// +build all ds tree trie test

package trie

import (
	"fmt"
	"sort"
	"testing"

	"github.com/dockerian/go-coding/utils"
	"github.com/stretchr/testify/assert"
)

// RunePhasesTestCase struct
type RunePhasesTestCase struct {
	Data     []string
	Expected []string
}

// TestRunePhases tests RunePhases
func TestRunePhases(t *testing.T) {
	testName := utils.GetTestName(t)
	var tests = []RunePhasesTestCase{
		{[]string{"a", "A", "about", "123", "xyz", "zhu"},
			[]string{"123", "A", "a", "about", "xyz", "zhu"}},
		{[]string{""}, []string{""}},
	}

	for index, test := range tests {
		var val = make([]string, len(test.Data))
		copy(val, test.Data)
		sort.Sort(RunePhases(val))
		var msg = fmt.Sprintf("expecting %v == %v", val, test.Expected)
		t.Logf("Test %v [%v]: %v\n", index+1, testName, msg)
		assert.Equal(t, test.Expected, val, msg)
	}
}
