package interview

import (
	"fmt"
	"math"

	u "github.com/dockerian/go-coding/utils"
)

// Point struct represents a location on the map
type Point struct {
	x, y, z float64
}

// PointDistance struct
type PointDistance struct {
	point    *Point
	distance float64
}

// ClosestPoints is a max heap
// see https://golang.org/pkg/container/heap/
type ClosestPoints []*PointDistance

// Len implements sort.Interface in heap.Interface
func (cp ClosestPoints) Len() int {
	return len(cp)
}

// Less implements sort.Interface in heap.Interface
func (cp ClosestPoints) Less(i, j int) bool {
	return cp[i].distance > cp[j].distance
}

// Swap implements sort.Interface in heap.Interface
func (cp ClosestPoints) Swap(i, j int) {
	cp[i], cp[j] = cp[j], cp[i]
}

// Peek returns the top item on the heap
func (cp *ClosestPoints) Peek() interface{} {
	h := *cp
	item := h[len(h)-1]
	return item
}

// Pop implements help.Interface
func (cp *ClosestPoints) Pop() interface{} {
	old := *cp
	n := len(old)
	item := old[n-1]
	*cp = old[0 : n-1]
	return item
}

// Push implements help.Interface
func (cp *ClosestPoints) Push(x interface{}) {
	item := x.(*PointDistance)
	*cp = append(*cp, item)
}

// DistanceTo returns distance to other point
func (p *Point) DistanceTo(other *Point) float64 {
	return math.Abs(p.x-other.x) + math.Abs(p.y-other.y) + math.Abs(p.z-other.z)
}

// GetClosest returns closest k points
// Give a point/location p and a list of many many other locations
// calculate K locations that are closest to the point p
func (p *Point) GetClosest(others []*Point, k int) []*Point {
	heap := &ClosestPoints{}

	if len(others) >= k {
		return others
	}

	for _, other := range others {
		heapSize := heap.Len()
		otherDistance := p.DistanceTo(other)

		if heapSize >= k {
			maxp := heap.Peek().(*PointDistance)

			if maxp.distance > otherDistance {
				heap.Pop()
				heap.Push(other)
			}
		} else {
			item := &PointDistance{point: other, distance: otherDistance}
			heap.Push(item)
		}
	}

	// convert heap to []*Point
	result := make([]*Point, k)
	for i := 0; i < k; i++ {
		item := heap.Pop().(*PointDistance)
		result[i] = item.point
	}

	return result
}

// String func for Point
func (p *Point) String() string {
	return fmt.Sprintf("{%v, %v, %v}", p.x, p.y, p.z)
}

// FindMeetupPoint returns the best meetup point for all given points
func FindMeetupPoint(points []*Point) *Point {
	return findMeetupPoint(points, true)
}

// findMeetupPointByMidPoint finds the meetup point by mid point
func findMeetupPointByMidPoint(points []*Point) *Point {
	count := len(points)
	if count == 0 {
		return nil
	}
	var sumX, sumY, sumZ float64
	for _, s := range points {
		sumX += s.x
		sumY += s.y
		sumZ += s.z
	}
	countf := float64(count)
	midPoint := &Point{x: sumX / countf, y: sumY / countf, z: sumZ / countf}
	u.Debug("\npoints: %v\n", points)
	u.Debug("  mid point: %+v\n", midPoint)

	var shortest = math.MaxFloat64
	var destPoint *Point
	for _, s := range points {
		distance := s.DistanceTo(midPoint)
		u.Debug("  distance to mid point: %+v [%+v]\n", distance, s)
		if distance < shortest {
			shortest = distance
			destPoint = s
		}
	}
	return destPoint
}

// findMeetupPoint returns the best meetup point for all given points
func findMeetupPoint(points []*Point, cache bool) *Point {
	var saved = make(map[*Point]map[*Point]float64)
	var shortest = math.MaxFloat64
	var meetup *Point

	u.Debug("\npoints: %v\n", points)
	for _, s := range points {
		var sum float64
		for _, d := range points {
			ds, ok := saved[d]
			var distance float64
			var dvisited bool
			if cache {
				if ok {
					if distance, dvisited = ds[s]; dvisited {
						sum += distance
					}
				} else {
					saved[d] = make(map[*Point]float64)
				}
			}
			if !dvisited {
				distance = s.DistanceTo(d)
				if cache {
					saved[d][s] = distance
				}
				sum += distance
			}
		}
		u.Debug("  sum = %v, distances to point %+v\n", sum, s)
		if sum < shortest {
			shortest = sum
			meetup = s
		}
	}

	return meetup
}

// FindShortest returns shortest paths between start and end, in a matrix
// NxN array [][]int, with each row i as distance from i to other [0..N)
// value -1 indicates no path
func FindShortest(matrix [][]int, start, end int) []int {
	var sizeX = len(matrix)
	var paths = make([]int, 0)
	var prevs = make(map[int]int)
	var visit = make([]bool, sizeX)
	var queue = make([]int, 0, sizeX)
	var distL = math.MaxInt32

	queue = append(queue, start)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		visit[curr] = true

		if curr == end {
			currPath := make([]int, 0, sizeX)
			currLength := 0
			prev, ok := prevs[curr]
			for ok {
				if prev != start {
					currPath = append(currPath, prev)
				}
				currLength += matrix[prev][curr]
				curr = prev
				prev, ok = prevs[curr]
			}

			if currLength > 0 && currLength < distL {
				paths = currPath[0:]
				distL = currLength
			}
		}

		var next = curr + 1
		for next < sizeX {
			if matrix[curr][next] != -1 && !visit[next] {
				prevs[next] = curr
				queue = append(queue, next)
				break
			}
			next++
		}
	}

	return paths
}
