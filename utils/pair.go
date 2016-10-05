package utils

import (
	"fmt"
)

// Pair struct respresents a pair of anything
type Pair struct {
	Item1, Item2 interface{}
}

// AreEqual method receiver compares 'this' Pair to 'other' Pair
func (p *Pair) AreEqual(other *Pair) bool {
	if p == nil || other == nil {
		return false
	}
	return p.Item1 == other.Item1 && p.Item2 == other.Item2
}

// String method receiver for Pair struct
func (p *Pair) String() string {
	if p == nil {
		return ""
	}
	return fmt.Sprintf("{%+v, %+v}", p.Item1, p.Item2)
}
