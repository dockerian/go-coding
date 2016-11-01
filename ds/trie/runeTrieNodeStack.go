package trie

import (
	"fmt"

	u "github.com/dockerian/go-coding/utils"
)

// -------- RuneTrieNodeItem, and method receivers

// RuneTrieNodeItem struct
type RuneTrieNodeItem struct {
	node *RuneTrieNode
	data string
}

// String func for RuneTrieItem
func (e *RuneTrieNodeItem) String() string {
	return fmt.Sprintf("node: %+v, data: %s", e.node, e.data)
}

// -------- RuneTrieStack, and method receivers

// RuneTrieStack struct
type RuneTrieStack struct {
	top      *RuneTrieNodeItem
	elements []*RuneTrieNodeItem
	size     int
}

// Peek is a pointer method receiver for RuneTrieStack to peek the top item
func (s *RuneTrieStack) Peek() *RuneTrieNodeItem {
	if s.size > 0 {
		v := s.top
		u.Debug("stack: {%+v}, peek: {%+v}\n", s, v)
		return v
	}
	return nil
}

// Pop is a pointer method receiver for RuneTrieStack to pop the top item
func (s *RuneTrieStack) Pop() *RuneTrieNodeItem {
	if s.size > 0 {
		v := s.top
		// u.Debug("- pop: {%+v}, from: {%+v}\n", v, s)
		s.elements = s.elements[0 : s.size-1]
		s.size--
		if s.size > 0 {
			s.top = s.elements[s.size-1]
		} else {
			s.top = nil
		}
		// u.Debug("stack: {%+v}, after - pop: {%+v}\n", s, v)
		return v
	}
	return nil
}

// Push is a pointer method receiver for RuneTrieStack to push into the stack
func (s *RuneTrieStack) Push(p *RuneTrieNodeItem) {
	if s.size == 0 {
		s.elements = make([]*RuneTrieNodeItem, 0)
	}
	// u.Debug("+push: {%+v}, to: {%+v}\n", p, s)
	s.top = p
	s.elements = append(s.elements, p)
	s.size++
	// u.Debug("stack: {%+v}, after push: {%+v}\n", s, p)
}

// Size is a pointer method receiver for RuneTrieStack to return the stack size
func (s *RuneTrieStack) Size() int {
	return s.size
}

// String func for RuneTrieStack
func (s *RuneTrieStack) String() string {
	return fmt.Sprintf("size: %d, top: {%+v}", s.size, s.top)
}
