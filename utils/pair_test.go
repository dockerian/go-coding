// +build all utils pair test

package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type PairTestCase struct {
	Item1, Item2 interface{}
	Expected     Pair
}

func TestPair(t *testing.T) {
	tests := []PairTestCase{
		{1, 100, Pair{1, 100}},
		{"aaa", "bbb", Pair{"aaa", "bbb"}},
	}

	for index, test := range tests {
		var msg = fmt.Sprintf("expecting {%+v, %+v} == %v",
			test.Item1, test.Item2, test.Expected)
		var pair = &Pair{test.Item1, test.Item2}
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.True(t, pair.AreEqual(&test.Expected), msg)
	}
}
