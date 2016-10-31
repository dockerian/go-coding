package trie

import "strings"

// -------- RunePhases, and method receivers

// RunePhases type
type RunePhases []string

// Len implements Len() in sort.Interface for RunePhases
func (rs RunePhases) Len() int {
	return len(rs)
}

// Less implements Less() in sort.Interface for RunePhases
func (rs RunePhases) Less(x, y int) bool {
	return strings.Compare(rs[x], rs[y]) < 0
}

// Swap implements Swap() in sort.Interface for RunePhases
func (rs RunePhases) Swap(x, y int) {
	rs[x], rs[y] = rs[y], rs[x]
}
