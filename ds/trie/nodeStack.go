package trie

import (
	"fmt"

	u "github.com/dockerian/go-coding/utils"
)

// -------- NodeItem, NodeStack, and method receivers

// NodeItem struct
type NodeItem struct {
	node *Node
	data string
}

// NodeStack struct
type NodeStack struct {
	top      *NodeItem
	elements []*NodeItem
	size     int
}

// String func for NodeItem
func (e *NodeItem) String() string {
	return fmt.Sprintf("node: %+v, data: %s", e.node, e.data)
}

// Peek is a pointer method receiver for NodeStack to peek the top item
func (s *NodeStack) Peek() *NodeItem {
	if s.size > 0 {
		v := s.top
		u.Debug("stack: {%+v}, peek: {%+v}\n", s, v)
		return v
	}
	return nil
}

// Pop is a pointer method receiver for NodeStack to pop the top item
func (s *NodeStack) Pop() *NodeItem {
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

// Push is a pointer method receiver for NodeStack to push into the stack
func (s *NodeStack) Push(p *NodeItem) {
	if s.size == 0 {
		s.elements = make([]*NodeItem, 0)
	}
	// u.Debug("+push: {%+v}, to: {%+v}\n", p, s)
	s.top = p
	s.elements = append(s.elements, p)
	s.size++
	// u.Debug("stack: {%+v}, after push: {%+v}\n", s, p)
}

// Size is a pointer method receiver for NodeStack to return the stack size
func (s *NodeStack) Size() int {
	return s.size
}

// String func for NodeStack
func (s *NodeStack) String() string {
	return fmt.Sprintf("size: %d, top: {%+v}", s.size, s.top)
}
