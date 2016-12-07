package heap

// IComparable interface
type IComparable interface {
	Compare(other *IComparable) int
}

// IHeap interface for max-heap or min-heap
type IHeap interface {
	Add(x *IComparable)
	Delete(x *IComparable)
	DeleteAt(i int)
	ExtractMax() *IComparable
	ExtractMin() *IComparable
	GetSize() int
	IsEmpty() bool
	PeekMax() *IComparable
	PeekMin() *IComparable
}
