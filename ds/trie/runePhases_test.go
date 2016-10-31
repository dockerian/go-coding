// +build all ds tree trie test

package trie

import (
	"fmt"
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
	var tests = []RunePhasesTestCase{}

	for index, test := range tests {
		var val = test.Expected
		var msg = fmt.Sprintf("expecting %v == %v", val, test.Expected)
		t.Logf("Test %v [%v]: %v\n", index+1, testName, msg)
		assert.Equal(t, test.Expected, val, msg)
	}
}
