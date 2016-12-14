package interview

import (
	"fmt"
	"math"
	// u "github.com/dockerian/go-coding/utils"
)

// Point struct represents a location on the map
type Point struct {
	x, y, z float64
}

// distanceTo returns distance to other point
func (p *Point) distanceTo(other *Point) float64 {
	return math.Abs(p.x-other.x) + math.Abs(p.y-other.y) + math.Abs(p.z-other.z)
}

// String func for Point
func (p *Point) String() string {
	return fmt.Sprintf("{%v, %v, %v}", p.x, p.y, p.z)
}

// FindMeetupPoint returns the best meetup point for all given points
func FindMeetupPoint(points []*Point) *Point {
	return findMeetupPoint(points, true)
}

// findMeetupPoint returns the best meetup point for all given points
func findMeetupPoint(points []*Point, cache bool) *Point {
	var saved = make(map[*Point]map[*Point]float64)
	var shortest = math.MaxFloat64
	var meetup *Point

	// u.Debug("points: %v\n", points)
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
				distance = s.distanceTo(d)
				if cache {
					saved[d][s] = distance
				}
				sum += distance
			}
		}
		// u.Debug("  sum = %v, from point %+v\n", sum, s)
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
