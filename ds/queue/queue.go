// Package queue :: queue.go
package queue

// Element respresents an item in Queue
type Element interface{}

// Queue is a slice of Element
type Queue []Element

// NewQueue constructs a new and empty Queue.
func NewQueue() Queue {
	return Queue{}
}

// Dequeue remove the first element from queue
func (q *Queue) Dequeue() Element {
	if len(*q) > 0 {
		v := (*q)[0]
		*q = (*q)[1:]
		return v
	}
	return nil
}

// Enqueue adds (appends) an element into queue
func (q *Queue) Enqueue(v Element) {
	*q = append(*q, v)
}
