package parking

import (
	"fmt"
	"math"

	u "github.com/dockerian/go-coding/utils"
)

// Builder struct
type Builder struct {
	parking    *Parking
	regulation IParkingRegulation
	tags       map[int]string
}

// NewBuilder returns a parking builder
func NewBuilder(capacity, levels, groundLevels int, regulation IParkingRegulation) *Builder {
	capacity, levels, groundLevels =
		regulation.SetCapacities(capacity, levels, groundLevels)

	capacityCarpools := regulation.GetCarpoolCapacity()
	capacityCompacts := regulation.GetCompactCapacity()
	capacityHandicap := regulation.GetHandicapCapacity()

	tags := make(map[int]string)
	for i := groundLevels - levels; i < groundLevels; i++ {
		switch {
		case i < 0:
			tags[i] = string('A' - i - 1)
		case i > 0:
			tags[i] = string('1' + i)
		default:
			tags[0] = "L"
		}
	}

	return &Builder{
		parking: &Parking{
			capacity:          capacity,
			levels:            levels,
			groundLevels:      groundLevels,
			available:         capacity,
			availableHandicap: capacityHandicap,
			capacityCarpools:  capacityCarpools,
			capacityCompacts:  capacityCompacts,
			capacityHandicap:  capacityHandicap,
			carpoolSpots:      make(map[int][]*Spot),
			compactSpots:      make(map[int][]*Spot),
			handiSpots:        make([]*Spot, capacityHandicap, capacityHandicap),
			spots:             make(map[int][]*Spot),
		},
		regulation: regulation,
		tags:       tags,
	}
}

// Construct initializes a Parking instance
func (b *Builder) Construct() {
	levels := b.parking.levels
	capacity := b.parking.capacity
	capacityCarpools := b.parking.capacityCarpools
	capacityCompacts := b.parking.capacityCompacts
	capacityHandicap := b.parking.capacityHandicap
	capacityPerLevel := int(math.Ceil(float64(capacity)) / float64(levels))
	undergroundLevels := levels - b.parking.groundLevels

	count, factor := 0, 1

	for i := capacityPerLevel - 1; i > 0; i /= 10 {
		factor = factor * 10
	}

	totalCarpools := capacityCarpools
	totalHandicap := capacityHandicap
	assignSpots := func(l int) {
		countCarpools := capacityCarpools/levels + 1
		countCompacts := capacityCompacts / levels
		b.parking.carpoolSpots[l] = make([]*Spot, 0, countCarpools)
		b.parking.compactSpots[l] = make([]*Spot, 0, countCompacts)
		b.parking.spots[l] = make([]*Spot, 0, capacityPerLevel)
		for i := 1; i <= capacityPerLevel && count < capacity; i++ {
			var isCarpool, isCompact, isHandicap bool
			if totalHandicap > 0 {
				totalHandicap--
				isHandicap = true
			} else if countCarpools > 0 && totalCarpools > 0 {
				countCarpools--
				totalCarpools--
				isCarpool = true
			} else if i%2 == 1 && countCompacts > 0 {
				countCompacts--
				isCompact = true
			}
			spot := &Spot{
				id:         i + (undergroundLevels-l)*factor,
				spotNumber: i,
				level:      l,
				label:      fmt.Sprintf("%s-%d", b.tags[l], i),
				isOccupied: false,
				isHandicap: isHandicap,
				isCompact:  isCompact,
				isCarpool:  isCarpool,
			}
			switch {
			case isHandicap:
				x := capacityHandicap - totalHandicap - 1
				b.parking.handiSpots[x] = spot
			case isCarpool:
				b.parking.carpoolSpots[l] = append(b.parking.carpoolSpots[l], spot)
			case isCompact:
				b.parking.compactSpots[l] = append(b.parking.compactSpots[l], spot)
			}
			b.parking.spots[l] = append(b.parking.spots[l], spot)
			count--
		}
	}

	u.Debug(" Parking: %v\n", b.parking)
	u.Debug("  - tags: %+v\n", b.tags)

	for l := 0; l < b.parking.groundLevels; l++ {
		assignSpots(l)
		// u.Debug("  -Level: %v\n", b.parking.SprintLevel(l))
	}
	for l := -1; l >= -undergroundLevels; l-- {
		assignSpots(l)
		// u.Debug("  -Level: %v\n", b.parking.SprintLevel(l))
	}
}
