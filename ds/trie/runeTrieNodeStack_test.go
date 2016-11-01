// +build all ds tree trie test

package trie

import (
	"fmt"
	"testing"

	"github.com/dockerian/go-coding/utils"
	"github.com/stretchr/testify/assert"
)

// RuneTrieNodeStackTestCase struct
type RuneTrieNodeStackTestCase struct {
	Data     *RuneTrieNodeItem  `json:"data,omitempty"`
	Expected *RuneTrieNodeItem  `json:"expected,omitempty"`
	Stack    *RuneTrieStack `stack:"a,omitempty"`
}

// TestRuneTrieNodeStack tests RuneTrieNodeStack
func TestRuneTrieNodeStack(t *testing.T) {
	testName := utils.GetTestName(t)
	var tests = []RuneTrieNodeStackTestCase{}

	for index, test := range tests {
		var val = test.Expected
		var msg = fmt.Sprintf("expecting %v == %v", val, test.Expected)
		t.Logf("Test %v [%v]: %v\n", index+1, testName, msg)
		assert.Equal(t, test.Expected, val, msg)
	}
}
