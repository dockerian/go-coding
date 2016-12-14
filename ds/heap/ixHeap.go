package heap

import (
	"fmt"
	"math"
)

// IxHeap struct
type IxHeap struct {
	heap     []int
	position int // bound limit of the last item in the heap
	capacity int // capacity of the heap
	capLimit int // capacity limit of the heap
}

// NewIxHeap constructs an int heap per specific capacity and optionally limit
func NewIxHeap(v ...int) *IxHeap {
	var capacity int
	var capLimit = -1
	if len(v) > 0 {
		if v[0] > 0 {
			capacity = v[0]
		}
		if len(v) > 1 && v[1] > 1 {
			capLimit = v[1]
		}
	}
	return &IxHeap{
		heap:     make([]int, 1, capacity+1),
		capacity: capacity,
		capLimit: capLimit,
		position: 0,
	}
}

// Add func
func (h *IxHeap) Add(i int) {
	siz := len(h.heap) - 1
	if h.capLimit > 1 && siz >= h.capLimit {
		if min, err := h.PeekMin(); err == nil && i < min {
			h.ExtractMin()
		} else if max, err := h.PeekMax(); err == nil && i > max {
			h.ExtractMax()
		}
	}
	h.heap = append(h.heap, i)
	h.position = h.position + 1
	if len(h.heap)-1 > h.capacity {
		h.capacity = len(h.heap) - 1
	}
	h.bubbleUp(h.position)
}

// Delete func
func (h *IxHeap) Delete(item int) int {
	for k, p := range h.heap {
		if p == item {
			h.DeleteAt(k)
			return p
		}
	}
	return 0
}

// DeleteAt func
func (h *IxHeap) DeleteAt(k int) int {
	if k > 0 && h.position > 0 {
		item := h.heap[k]
		if k < h.position {
			h.heap[k] = h.heap[h.position]
			h.sinkDown(k)
			h.bubbleUp(k)
		}
		h.heap = h.heap[0:h.position]
		h.position = h.position - 1
		return item
	}
	return 0
}

// ExtractMax func
func (h *IxHeap) ExtractMax() int {
	max := h.heap[0]
	if h.position > 0 {
		j, k := h.position-1, h.position
		if k > 2 && k%2 == 1 {
			if h.heap[j] > h.heap[k] {
				h.heap[j], h.heap[k] = h.heap[k], h.heap[j]
			}
		}
		max = h.heap[k]
		h.heap = h.heap[0:k]
		h.position = k - 1
	}
	return max
}

// ExtractMin func
func (h *IxHeap) ExtractMin() int {
	min := h.heap[0]
	if h.position > 0 {
		min = h.heap[1]
		h.heap[1] = h.heap[h.position]
		h.heap = h.heap[0:h.position]
		h.position = h.position - 1
		h.sinkDown(1)
	}
	return min
}

// GetSize func returns the position of the heap
func (h *IxHeap) GetSize() int {
	return h.position
}

// IsEmpty checks if the heap is empty
func (h *IxHeap) IsEmpty() bool {
	return h.position == 0
}

// PeekMax func
func (h *IxHeap) PeekMax() (int, error) {
	if h.position > 0 {
		return h.heap[h.position], nil
	}
	return math.MinInt32, fmt.Errorf("Empty heap")
}

// PeekMin func
func (h *IxHeap) PeekMin() (int, error) {
	if h.position > 0 {
		return h.heap[1], nil
	}
	return math.MaxInt32, fmt.Errorf("Empty heap")
}

// String func
func (h *IxHeap) String() string {
	var str string
	for i, item := range h.heap {
		if i == 0 {
			str = "heap: ["
		} else {
			str = fmt.Sprintf("%s, %v", str, item)
		}
	}
	return str + " ]"
}

// bubbleUp the last item to correct position in the heap
func (h *IxHeap) bubbleUp(end int) {
	if 1 <= end && end <= h.position {
		pos := end
		for pos > 1 {
			i := pos / 2
			if h.heap[i] > h.heap[pos] {
				h.heap[pos], h.heap[i] = h.heap[i], h.heap[pos]
				if pos%2 == 1 && h.heap[pos] < h.heap[pos-1] {
					h.heap[pos], h.heap[pos-1] = h.heap[pos-1], h.heap[pos]
				}
			}
			pos = i
		}
	}
}

// sinkDown the item to correct position in the heap
func (h *IxHeap) sinkDown(k int) {
	pos := k
	pc1, pc2 := 2*k, 2*k+1            // position for left and right child
	if h.position >= pc1 && pc1 > 0 { // check potential overflow
		if h.heap[pos] > h.heap[pc1] {
			pos = pc1
		}
		if h.position >= pc2 && pc2 > 0 && h.heap[pos] > h.heap[pc2] {
			pos = pc2
		}
	}
	if pos != k {
		h.heap[k], h.heap[pos] = h.heap[pos], h.heap[k]
		h.sinkDown(pos)
	}
}
