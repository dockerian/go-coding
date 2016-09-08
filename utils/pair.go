// +build all utils pair

package utils

import (
	"fmt"
)

// Pair struct respresents a pair of anything
type Pair struct {
	Item1, Item2 interface{}
}

func (p *Pair) AreEqual(other *Pair) bool {
	return p.Item1 == other.Item1 && p.Item2 == other.Item2
}

func (p *Pair) string() string {
	return fmt.Sprintf("{%+v, %+v}", p.Item1, p.Item2)
}
