// +build all ds maps linkedlist test

package maps

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/dockerian/go-coding/ds/mathEx"
	"github.com/stretchr/testify/assert"
)

// LinkedListNumbersTestCase struct
type LinkedListNumbersTestCase struct {
	Num1 uint64
	Num2 uint64
	Sum  uint64
}

// LinkedListTestCase struct
type LinkedListTestCase struct {
	Data     interface{}
	Expected interface{}
}

// TestLinkedListNumbers tests LinkedListNumber functions/methods.
// See: https://leetcode.com/problems/add-two-numbers/
func TestLinkedListNumbers(t *testing.T) {
	tests := []LinkedListNumbersTestCase{
		{123, 321, 444},
		{0, 12345, 12345},
		{67890, 0, 67890},
		{777, 12300, 13077},
		{999999, 999999, 1999998},
		{1, 0, 1},
		{0, 0, 0},
	}

	for index, test := range tests {
		var ln1 = GetLinkedListNumber(test.Num1)
		var ln2 = GetLinkedListNumber(test.Num2)
		var sum = ln1.Add(ln2)
		var st1 = ln1.String()
		var st2 = ln2.String()
		var msg = fmt.Sprintf("expecting (%v + %v) == %v", st1, st2, test.Sum)
		var ms1 = fmt.Sprintf("%+v == %v [%+v]", test.Num1, st1, ln1)
		var ms2 = fmt.Sprintf("%+v == %v [%+v]", test.Num2, st2, ln2)
		t.Logf("Test %v: %v\n", index+1, msg)
		assert.Equal(t, strconv.FormatUint(test.Sum, 10), sum.String())
		assert.Equal(t, strconv.FormatUint(test.Num1, 10), st1, ms1)
		assert.Equal(t, strconv.FormatUint(test.Num2, 10), st2, ms2)

		var cmp = mathEx.Compare(test.Num1, test.Num2)
		msg = fmt.Sprintf("expecting compare LinkedListNumber (%v, %v) == %v",
			test.Num1, test.Num2, cmp)
		assert.Equal(t, cmp, ln1.Compare(ln2), msg)
	}
}

// TestLinkedList is a testing function template
func TestLinkedList(t *testing.T) {
	var tests []LinkedListTestCase

	for index, test := range tests {
		var val interface{}
		var msg = fmt.Sprintf("expecting %v == %v", val, test.Expected)
		t.Logf("Test %v: %v\n", index+1, msg)
	}
}
