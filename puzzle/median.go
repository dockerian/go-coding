package puzzle

import (
	"github.com/dockerian/go-coding/ds/heap"

	u "github.com/dockerian/go-coding/utils"
)

// FindMedian func
// Find the median (the middle value in an ordered integer list).
// For even size of the list, the median is the mean of the two middle value.
// See https://leetcode.com/problems/find-median-from-data-stream/
func FindMedian(stream []int) float64 {
	size := len(stream)
	maxH := heap.NewIxHeap(size / 2)
	minH := heap.NewIxHeap(size / 2)

	addNumber := func(v int) {
		maxH.Add(v)
		n := maxH.ExtractMin()
		minH.Add(n)
		if maxH.GetSize() < minH.GetSize() {
			n = minH.ExtractMax()
			maxH.Add(n)
		}
	}

	findMedian := func() float64 {
		if maxH.GetSize() > minH.GetSize() {
			m, _ := maxH.PeekMin()
			return float64(m)
		}
		a, _ := maxH.PeekMin()
		b, _ := minH.PeekMax()
		// u.Debug("median: (%d + %d) / 2\n", a, b)
		return (float64(a) + float64(b)) / 2.0
	}

	u.Debug("\ninputs: %+v\n", stream)
	for _, v := range stream {
		addNumber(v)
		// u.Debug("-largers: %+v\n", maxH)
		// u.Debug("-smaller: %+v\n", minH)
	}

	return findMedian()
}
