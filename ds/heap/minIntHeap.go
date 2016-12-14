package heap

import "container/heap"

// MinIntHeap is a min-heap of integers.
// - Construction  : h := &MinIntHeap{3, 1, 4}
// - Initialization: heap.Init(h) // sink down from half to top
// - Add an integer: heap.Push(h, 5) // call heap.Interface.Push then bubble up
// - Other mehotds : heap.Pop(h), heap.Fix(h, 3)
// - Peek min/max  : (*h)[0], (*h)[len(*h)-1]
type MinIntHeap []int64

// Len implements sort.Interface in heap.Interface
func (mih *MinIntHeap) Len() int { return len(*mih) }

// Les implements sort.Interface in heap.Interface
func (mih *MinIntHeap) Less(i, j int) bool { return (*mih)[i] < (*mih)[j] }

// Swap implements sort.Interface in heap.Interface
func (mih *MinIntHeap) Swap(i, j int) { (*mih)[i], (*mih)[j] = (*mih)[j], (*mih)[i] }

// ExtractMax removes the maximum item from heap
func (mih *MinIntHeap) ExtractMax() int64 {
	return heap.Remove(mih, 0).(int64)
}

// ExtractMin removes the minimum item from heap
func (mih *MinIntHeap) ExtractMin() int64 {
	return mih.Pop().(int64)
}

// Push implements heap.Interface
func (mih *MinIntHeap) Push(x interface{}) {
	// use pointer receivers to modify the slice
	*mih = append(*mih, x.(int64))
}

// Pop implements heap.Interface
func (mih *MinIntHeap) Pop() interface{} {
	h := *mih
	n := len(h)
	x := h[n-1]
	// use pointer receivers to modify the slice
	*mih = h[0 : n-1]
	return x
}
