package parking

import "math"

const (
	// COMPACTPERCENT specifies percentage of compact spots
	COMPACTPERCENT = 20.0
	// CARPOOLPERCENT specifies percentage of carpool spots
	CARPOOLPERCENT = 7.5
	// DEFAULTCAPACITY specifies minimum total capacity for a parking garage
	DEFAULTCAPACITY = 100
	// DEFAULTHANDICAPPERCENT sepcifies percentage of handicap spots
	DEFAULTHANDICAPPERCENT = 5.0
	// DEFAULTGROUNDLEVELS specifies minimum number of ground levels
	DEFAULTGROUNDLEVELS = 0
	// DEFAULTLEVELS specifies minimum number of parking garage levels
	DEFAULTLEVELS = 1
)

// IParkingRegulation interface
type IParkingRegulation interface {
	SetCapacities(int, int, int) (int, int, int)
	GetCarpoolCapacity() int
	GetCompactCapacity() int
	GetHandicapCapacity() int
}

// CityRegulation implements IParkingRegulation
type CityRegulation struct {
	capacity     int
	groundLevels int
	levels       int
}

// GetCityRegulation returns a city regulation
func GetCityRegulation(capacity, levels, groundLevels int) CityRegulation {
	r := &CityRegulation{
		capacity:     capacity,
		levels:       levels,
		groundLevels: groundLevels,
	}
	r.SetCapacities(capacity, levels, groundLevels)
	return *r
}

// SetCapacities returns capacity, levels, and groundLevels
func (r *CityRegulation) SetCapacities(capacity, levels, groundLevels int) (int, int, int) {
	r.capacity, r.levels, r.groundLevels = capacity, levels, groundLevels

	if capacity < DEFAULTCAPACITY {
		r.capacity = DEFAULTCAPACITY
	}
	if levels < DEFAULTLEVELS {
		r.levels = DEFAULTLEVELS
	}
	if groundLevels < DEFAULTGROUNDLEVELS || groundLevels > levels {
		r.groundLevels = DEFAULTGROUNDLEVELS
	}

	return r.capacity, r.levels, r.groundLevels
}

// GetCarpoolCapacity returns carpool capacity per the regulation
func (r *CityRegulation) GetCarpoolCapacity() int {
	capacityCarpools := int(math.Ceil(float64(r.capacity) * float64(CARPOOLPERCENT) / 100))
	if capacityCarpools <= 0 {
		capacityCarpools = 1
	}
	return capacityCarpools
}

// GetCompactCapacity returns compact capacity per the regulation
func (r *CityRegulation) GetCompactCapacity() int {
	capacityCompacts := int(math.Ceil(float64(r.capacity) * float64(COMPACTPERCENT) / 100))
	return capacityCompacts
}

// GetHandicapCapacity returns handicap capacity per the regulation
func (r *CityRegulation) GetHandicapCapacity() int {
	capacityHandiCap := int(math.Ceil(float64(r.capacity) * float64(DEFAULTHANDICAPPERCENT) / 100))
	if capacityHandiCap <= 0 {
		capacityHandiCap = 1
	}
	return capacityHandiCap
}
