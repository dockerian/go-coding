package heap

import "fmt"

// Heap struct
type Heap struct {
	heap     []int
	position int // bound limit of the last item in the heap
	capacity int // capacity of the heap
}

// NewHeap constructs a heap per specific capacity
func NewHeap(capacity int) *Heap {
	if capacity < 0 {
		capacity = 0
	}
	return &Heap{
		heap:     make([]int, 1, capacity+1),
		capacity: capacity,
		position: 0,
	}
}

// Add func
func (h *Heap) Add(i int) {
	h.heap = append(h.heap, i)
	h.position = h.position + 1
	if len(h.heap) > h.capacity {
		h.capacity = len(h.heap) - 1
	}
	h.bubbleUp()
}

// Delete func
func (h *Heap) Delete(item int) int {
	for k, p := range h.heap {
		if p == item {
			h.DeleteAt(k)
			return p
		}
	}
	return 0
}

// DeleteAt func
func (h *Heap) DeleteAt(k int) int {
	if k > 0 && h.position > 0 {
		item := h.heap[k]
		if k < h.position {
			h.heap[k] = h.heap[h.position]
			h.sinkDown(k)
		}
		h.heap = h.heap[0:h.position]
		h.position = h.position - 1
		return item
	}
	return 0
}

// ExtractMax func
func (h *Heap) ExtractMax() int {
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
func (h *Heap) ExtractMin() int {
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
func (h *Heap) GetSize() int {
	return h.position
}

// IsEmpty checks if the heap is empty
func (h *Heap) IsEmpty() bool {
	return h.position == 0
}

// PeekMax func
func (h *Heap) PeekMax() (int, error) {
	if h.position > 0 {
		return h.heap[h.position], nil
	}
	return 0, fmt.Errorf("Empty heap")
}

// PeekMin func
func (h *Heap) PeekMin() (int, error) {
	if h.position > 0 {
		return h.heap[1], nil
	}
	return 0, fmt.Errorf("Empty heap")
}

// String func
func (h *Heap) String() string {
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
func (h *Heap) bubbleUp() {
	if h.position > 1 {
		pos := h.position
		for pos > 1 && h.heap[pos/2] > h.heap[pos] {
			h.heap[pos], h.heap[pos/2] = h.heap[pos/2], h.heap[pos]
			if pos%2 == 1 && h.heap[pos] < h.heap[pos-1] {
				h.heap[pos], h.heap[pos-1] = h.heap[pos-1], h.heap[pos]
			}
			pos = pos / 2
		}
	}
}

// sinkDown the item to correct position in the heap
func (h *Heap) sinkDown(k int) {
	pos := k
	if 2*k <= h.position && h.heap[pos] > h.heap[2*k] {
		pos = 2 * k
	}
	if 2*k+1 <= h.position && h.heap[pos] > h.heap[2*k+1] {
		pos = 2 * k
	}
	if pos != k {
		h.heap[k], h.heap[pos] = h.heap[pos], h.heap[k]
		h.sinkDown(pos)
	}
}
