package parking

// Spot struct
type Spot struct {
	id         int
	level      int
	spotNumber int
	label      string
	isOccupied bool
	isHandicap bool
	isCarpool  bool
	isCompact  bool
}

// CompareDistance compares distances to entry
func (s *Spot) CompareDistance(other *Spot) int {
	if s == nil && other == nil {
		return 0
	}
	if s == nil && other != nil {
		return -1
	}
	if s != nil && other == nil {
		return 1
	}
	if s.spotNumber > other.spotNumber {
		return 1
	}
	if s.spotNumber < other.spotNumber {
		return -1
	}
	return 0
}

// GetLabel func for Spot
func (s *Spot) GetLabel() string {
	return s.label
}

// String func for Spot
func (s *Spot) String() string {
	str := s.label
	if s.isOccupied {
		str += ": Occupied"
	} else {
		str += ": Availble"
	}
	switch {
	case s.isCarpool:
		str += " [Carpool]"
	case s.isCarpool:
		str += " [Compact]"
	case s.isHandicap:
		str += " [Handicap]"
	default:
		str += " [*]"
	}
	return str
}
