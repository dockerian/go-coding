package exp

import "fmt"

// OpQueue struct
type OpQueue struct {
	first    *OpItem
	elements []*OpItem
	size     int
}

// IsEmpty returns true if stack has no element
func (s *OpQueue) IsEmpty() bool {
	return s.size == 0
}

// Len returns the stack's length/size
func (s *OpQueue) Len() int {
	return s.size
}

// Peek returns the first of the stack without popping out the item
func (s *OpQueue) Peek() string {
	if s.first != nil {
		return s.first.Op
	}
	return ""
}

// Dequeue returns the value of the first element; or "" if the queue is empty
func (s *OpQueue) Dequeue() string {
	p := s.DequeueItem()
	if p != nil {
		return p.Op
	}
	return ""
}

// DequeueItem returns the first element; or nil if the queue is empty
func (s *OpQueue) DequeueItem() *OpItem {
	if s.size > 0 {
		p := s.elements[0]
		if s.size > 1 {
			s.first = s.elements[1]
		} else {
			s.first = nil
		}
		s.elements = s.elements[1:s.size]
		s.size--
		return p
	}
	return nil
}

// Enqueue a new element at the end of the queue
func (s *OpQueue) Enqueue(token string) {
	if s.size == 0 {
		s.elements = make([]*OpItem, 0, 1)
	}
	item := OpItem{Op: token, Expression: token}
	s.first = &item
	s.elements = append(s.elements, &item)
	s.size++
}

// EnqueueItem adds an item at the end of the queue
func (s *OpQueue) EnqueueItem(item *OpItem) {
	if s.size == 0 {
		s.elements = make([]*OpItem, 0, 1)
	}
	s.first = item
	s.elements = append(s.elements, item)
	s.size++
}

// String func
func (s OpQueue) String() string {
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
