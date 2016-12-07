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
	maxH := heap.NewHeap(size / 2)
	minH := heap.NewHeap(size / 2)

	addNumber := func(v int) {
		maxH.Add(v)
		n := 0 - maxH.ExtractMin()
		minH.Add(n)
		if maxH.GetSize() < minH.GetSize() {
			n = 0 - minH.ExtractMin()
			maxH.Add(n)
		}
	}

	findMedian := func() float64 {
		if maxH.GetSize() > minH.GetSize() {
			m, _ := maxH.PeekMin()
			return float64(m)
		}
		a, _ := maxH.PeekMin()
		b, _ := minH.PeekMin()
		return float64(a+b) / float64(2)
	}

	u.Debug("\ninputs: %+v\n", stream)
	for _, v := range stream {
		addNumber(v)
	}

	// u.Debug("-largers: %+v\n", maxH)
	// u.Debug("-smaller: %+v\n", minH)
	return findMedian()
}
