package puzzle

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
