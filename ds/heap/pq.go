package heap

import "container/heap"

// PriorityQueue is a list of *Item
type PriorityQueue []*Item

// Len implements sort.Interface in heap.Interface
func (pq PriorityQueue) Len() int { return len(pq) }

// Less implements sort.Interface in heap.Interface
func (pq PriorityQueue) Less(x, y int) bool {
	return pq[x].priority > pq[y].priority
}

// Swap implements sort.Interface in heap.Interface
func (pq PriorityQueue) Swap(x, y int) {
	pq[x], pq[y] = pq[y], pq[x]
}

// Pop implements heap.Interface
func (pq *PriorityQueue) Pop() interface{} {
	q := *pq
	size := len(q)
	item := q[size-1]
	item.index = -1 // indicate it is out of the queue
	*pq = q[0 : size-1]
	return item
}

// Push implements heap.Interface
func (pq *PriorityQueue) Push(x interface{}) {
	size := len(*pq)
	item := x.(*Item)
	item.index = size
	*pq = append(*pq, item)
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value interface{}, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
