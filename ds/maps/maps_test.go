// +build all ds maps test

package maps

import (
	"fmt"
	"testing"
)

// MapTestCase struct
type MapTestCase struct {
	Data     interface{}
	Expected interface{}
}

// TestMap is a testing function template
func TestMap(t *testing.T) {
	var tests []MapTestCase

	for index, test := range tests {
		var val interface{}
		var msg = fmt.Sprintf("expecting %v == %v", val, test.Expected)
		t.Logf("Test %v: %v\n", index+1, msg)
	}
}
