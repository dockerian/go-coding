// +build all pkg cardgame test

package cardgame

import (
	"fmt"
	"testing"

	"github.com/dockerian/go-coding/utils"
	"github.com/stretchr/testify/assert"
)

// GameTestCase struct
type GameTestCase struct {
	Data     Game
	Expected interface{}
}

// TestGame tests Game
func TestGame(t *testing.T) {
	testName := utils.GetTestName(t)
	tests := []GameTestCase{}

	for index, test := range tests {
		var foo = test.Expected
		var val = test.Expected
		var msg = fmt.Sprintf("expecting %v == %v", val, test.Expected)
		t.Logf("Test %v [%v]: %v\n", index+1, testName, msg)
		assert.Equal(t, test.Expected, foo, msg)
		assert.Equal(t, test.Expected, val, msg)
	}
}
