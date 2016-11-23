// stack implementations
// see https://gist.github.com/bemasher/1777766

package stack

// Element struct
type Element struct {
	value interface{}
	next  *Element
}

// Stack struct
type Stack struct {
	top      *Element
	elements []*Element
	size     int
}

// IsEmpty returns true if stack has no element
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// Len returns the stack's length/size
func (s *Stack) Len() int {
	return s.size
}

// Peek returns the top of the stack without popping out the item
func (s *Stack) Peek() interface{} {
	if s.size > 0 {
		peek := s.top.value
		return peek
	}
	return nil
}

// Pop the top element from the stack and return it's value
// returns nil if the stack is empty
func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		s.top, value = s.top.next, s.top.value
		s.elements = s.elements[:s.size-1]
		s.size--
		return
	}
	return nil
}

// Push a new element on top of the stack
func (s *Stack) Push(value interface{}) {
	if s.size == 0 {
		s.elements = make([]*Element, 0, 1)
	}
	s.top = &Element{value, s.top}
	s.elements = append(s.elements, s.top)
	s.size++
}
