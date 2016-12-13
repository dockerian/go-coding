package parking

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

var parkingMutex sync.Mutex

// Parking struct
type Parking struct {
	levels            int
	groundLevels      int
	capacity          int
	capacityCarpools  int
	capacityCompacts  int
	capacityHandicap  int
	available         int
	availableHandicap int
	carpoolSpots      map[int][]*Spot
	compactSpots      map[int][]*Spot
	handiSpots        []*Spot
	spots             map[int][]*Spot
}

// New create a Parking instance with capacity, levels, and groundLevels
func New(capacity, levels, groundLevels int, regulation IParkingRegulation) *Parking {
	builder := NewBuilder(capacity, levels, groundLevels, regulation)
	builder.Construct()
	return builder.parking
}

// FindAvailableSpot returns the first best available to the entryLevel
// per specific handicap, carpool, or compact property of the vehical
// note: the caculation for the best spot assumes for the walking distance to
//       the entry level, instead of driving distance to the spot, since the
//       the parking does not specify which level is the garage entry;
//       otherwise, the best/first available spot would be different.
func (p *Parking) FindAvailableSpot(
	isHandicap, isCarpool, isCompact bool, entryLevel int) *Spot {
	parkingMutex.Lock()
	defer parkingMutex.Unlock()

	if p.available <= 0 || isHandicap && p.availableHandicap <= 0 {
		return nil
	}

	bottom := p.groundLevels - p.levels
	if entryLevel > p.groundLevels || entryLevel < bottom {
		entryLevel = p.levels + bottom - 1
	}

	top, down, up := p.groundLevels, entryLevel-1, entryLevel+1

	// u.Debug("  Level: bottom, down [%d, %d]; up, top: [%d, %d]\n", bottom, down, up, top)

	findSpots := func(level int) *Spot {
		// u.Debug("  Check: %+v\n", p.SprintLevel(level))
		if bottom <= level || level < top {
			for _, spot := range p.spots[level] {
				if !spot.isOccupied {
					switch {
					case spot.isCarpool && isCarpool == spot.isCarpool:
						return spot
					case spot.isCompact && isCompact == spot.isCompact:
						return spot
					case !(spot.isCarpool || spot.isCompact || spot.isHandicap):
						return spot
					}
				}
			}
		}
		return nil
	}

	switch {
	case isHandicap:
		// handicapLevel := p.handiSpots[0].levl
		// u.Debug("  Check: %+v\n", p.SprintLevel(handicapLevel))
		for _, spot := range p.handiSpots {
			if !spot.isOccupied {
				return spot
			}
		}
	default:
		if spot := findSpots(entryLevel); spot != nil {
			return spot
		}
		for up < top || down >= bottom {
			spot1 := findSpots(up)
			spot2 := findSpots(down)
			dcomp := spot1.CompareDistance(spot2)
			if spot1 != nil && dcomp >= 0 {
				return spot1
			}
			if spot2 != nil && dcomp < 0 {
				return spot2
			}
			down--
			up++
		}
	}
	return nil
}

// Park occupies a specific spot
func (p *Parking) Park(spot *Spot) {
	parkingMutex.Lock()
	defer parkingMutex.Unlock()

	spot.isOccupied = true
	if spot.isHandicap {
		p.availableHandicap--
	}
	p.available--
}

// SprintLevel returns a parking map for specific level
func (p *Parking) SprintLevel(level int) string {
	s := fmt.Sprintf("%2d: ", level)
	for _, spot := range p.spots[level] {
		mark := "_"
		switch {
		case spot.isCarpool:
			mark = "P"
		case spot.isCompact:
			mark = "C"
		case spot.isHandicap:
			mark = "H"
		}
		if spot.isOccupied {
			if mark == "_" {
				mark = "*"
			}
			mark = fmt.Sprintf("(%s)", mark)
		} else {
			mark = fmt.Sprintf("_%s|", mark)
		}
		s += mark
	}
	return s
}

// SprintMap returns a parking map
func (p *Parking) SprintMap() string {
	x := 1 + int(math.Ceil(float64(p.capacity))/float64(p.levels))
	s := fmt.Sprintf("%s\n L: ", p.String())
	for i := 1; i <= x/10; i++ {
		s += fmt.Sprintf("%s %d ", strings.Repeat("   ", 9), i)
	}
	s += "\n----"
	for i := 1; i < x; i++ {
		s += fmt.Sprintf("-%d-", i%10)
	}
	for i := 0; i < p.levels; i++ {
		level := p.levels - p.groundLevels - i
		s += fmt.Sprintf("\n%s", p.SprintLevel(level))
	}
	return s
}

// String func for Parking
func (p *Parking) String() string {
	s := fmt.Sprintf(
		"Capacity: %d / %d, Handicap: %d / %d, Carpool: %d, Compact: %d",
		p.available,
		p.capacity,
		p.availableHandicap,
		p.capacityHandicap,
		p.capacityCarpools,
		p.capacityCompacts)
	return s
}
