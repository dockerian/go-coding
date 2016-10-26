// +build all demo foo test

package demo

import (
	"fmt"
	"testing"

	"github.com/dockerian/go-coding/utils"
	"github.com/stretchr/testify/assert"
)

// FooTestCase struct
type FooTestCase struct {
	Data     Foo
	Expected interface{}
}

// TestFoo is a testing function template
func TestFoo(t *testing.T) {
	testName := utils.GetTestName(t)
	tests := []FooTestCase{
		{Foo{100}, 100},
		{Foo{"aaa"}, "aaa"},
	}

	for index, test := range tests {
		var foo = &Foo{test.Expected}
		var val = test.Data.GetAnything()
		var msg = fmt.Sprintf("expecting %v == %v", val, test.Expected)
		t.Logf("Test %v [%v]: %v\n", index+1, testName, msg)
		assert.Equal(t, test.Expected, foo.GetAnything(), msg)
		assert.Equal(t, test.Expected, val, msg)
	}
}
