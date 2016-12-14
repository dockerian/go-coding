package heap

import "container/heap"

// Item defines a heap item
type Item struct {
	index    int
	priority int
	value    interface{}
}

// Heap is a list of *Item
type Heap []*Item

// New returns a new Heap
func New() *Heap {
	h := make(Heap, 0)
	return &h
}

// Len implements sort.Interface in heap.Interface
func (h *Heap) Len() int { return len(*h) }

// Less implements sort.Interface in heap.Interface
func (h *Heap) Less(x, y int) bool {
	return (*h)[x].priority > (*h)[y].priority
}

// Swap implements sort.Interface in heap.Interface
func (h *Heap) Swap(x, y int) {
	(*h)[x], (*h)[y] = (*h)[y], (*h)[x]
}

// ExtractMax removes the maximum item from heap
func (h *Heap) ExtractMax() *Item {
	return heap.Remove(h, 0).(*Item)
}

// ExtractMin removes the minimum item from heap
func (h *Heap) ExtractMin() *Item {
	return h.Pop().(*Item)
}

// PeekMax returns the maximum item from heap
func (h *Heap) PeekMax() *Item {
	return (*h)[0]
}

// PeekMin returns the minimum item from heap
func (h *Heap) PeekMin() *Item {
	return (*h)[len(*h)-1]
}

// Pop implements heap.Interface
func (h *Heap) Pop() interface{} {
	q := *h
	size := len(q)
	item := q[size-1]
	item.index = -1 // indicate it is out of the heap
	*h = q[0 : size-1]
	return item
}

// Push implements heap.Interface
func (h *Heap) Push(x interface{}) {
	size := len(*h)
	item := x.(*Item)
	item.index = size
	*h = append(*h, item)
}

// update modifies the priority and value of an Item in the queue.
func (h *Heap) update(item *Item, value interface{}, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(h, item.index)
}
