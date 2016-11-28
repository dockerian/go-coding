package stack

import "fmt"

// Str struct is a string stack implementation
type Str struct {
	top      string
	elements []string
	size     int
}

// IsEmpty returns true if stack has no element
func (s *Str) IsEmpty() bool {
	return s.size == 0
}

// Len returns the stack's length/size
func (s *Str) Len() int {
	return s.size
}

// Peek returns the top of the stack without popping out the item
func (s *Str) Peek() string {
	return s.top
}

// Pop and return the top element from the stack
// returns nil if the stack is empty
func (s *Str) Pop() string {
	if s.size > 0 {
		p := s.elements[s.size-1]
		if s.size > 1 {
			s.top = s.elements[s.size-2]
		} else {
			s.top = ""
		}
		s.elements = s.elements[:s.size-1]
		s.size--
		return p
	}
	return ""
}

// Push a new element on top of the stack
func (s *Str) Push(value string) {
	if s.size == 0 {
		s.elements = make([]string, 0, 1)
	}
	s.top = value
	s.elements = append(s.elements, value)
	s.size++
}

// String func
func (s Str) String() string {
	var ubound = s.size - 1
	var output = fmt.Sprintf("(%d) [ ", s.size)
	for i := 0; i <= ubound; i++ {
		pad, str := ", ", s.elements[i]
		if i == ubound {
			pad = ""
		}
		output = output + str + pad
	}

	return output + " ]"
}
