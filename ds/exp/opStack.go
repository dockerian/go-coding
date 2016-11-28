package exp

import "fmt"

// OpStack struct
type OpStack struct {
	top      *OpItem
	elements []*OpItem
	size     int
}

// IsEmpty returns true if stack has no element
func (s *OpStack) IsEmpty() bool {
	return s.size == 0
}

// Len returns the stack's length/size
func (s *OpStack) Len() int {
	return s.size
}

// Peek returns the top of the stack without popping out the item
func (s *OpStack) Peek() string {
	if s.top != nil {
		return s.top.Op
	}
	return ""
}

// Pop returns the value of the top element; or "" if the stack is empty
func (s *OpStack) Pop() string {
	p := s.PopItem()
	if p != nil {
		return p.Op
	}
	return ""
}

// PopItem returns the top element from the stack; or nil if the stack is empty
func (s *OpStack) PopItem() *OpItem {
	if s.size > 0 {
		p := s.elements[s.size-1]
		if s.size > 1 {
			s.top = s.elements[s.size-2]
		} else {
			s.top = nil
		}
		s.elements = s.elements[:s.size-1]
		s.size--
		return p
	}
	return nil
}

// Push a new element on top of the stack
func (s *OpStack) Push(token string) {
	if s.size == 0 {
		s.elements = make([]*OpItem, 0, 1)
	}
	item := OpItem{Op: token, Expression: token}
	s.top = &item
	s.elements = append(s.elements, &item)
	s.size++
}

// PushItem adds an item on top of the stack
func (s *OpStack) PushItem(item *OpItem) {
	if s.size == 0 {
		s.elements = make([]*OpItem, 0, 1)
	}
	s.top = item
	s.elements = append(s.elements, item)
	s.size++
}

// String func
func (s OpStack) String() string {
	var ubound = s.size - 1
	var output = fmt.Sprintf("(%d) [ ", s.size)
	for i := 0; i <= ubound; i++ {
		pad, str := ", ", s.elements[i].Op
		if i == ubound {
			pad = ""
		}
		output = output + str + pad
	}

	return output + " ]"
}
