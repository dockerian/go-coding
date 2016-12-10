// +build all pkg parking test

package parking

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ParkingEntry struct
type ParkingEntry struct {
	isHandicap bool
	isCarpool  bool
	isCompact  bool
	entryLevel int
	expected   Spot
}

// ParkingTestCase struct
type ParkingTestCase struct {
	capacity     int
	levels       int
	groundLevels int
	expected     *Parking
	entries      []ParkingEntry
}

// String func for ParkingEntry
func (e *ParkingEntry) String() string {
	s := "*"
	switch {
	case e.isHandicap:
		s = "H"
	case e.isCarpool:
		s = "P"
	case e.isCompact:
		s = "C"
	}
	return fmt.Sprintf("%2d [%s]", e.entryLevel, s)
}

// TestParking tests Parking functions
func TestParking(t *testing.T) {
	var reg IParkingRegulation
	var tests = []ParkingTestCase{
		{103, 5, 3, &Parking{
			levels:            5,
			groundLevels:      3,
			capacity:          103,
			capacityCarpools:  8,
			capacityCompacts:  21,
			capacityHandicap:  6,
			available:         103,
			availableHandicap: 6,
			carpoolSpots:      make(map[int][]*Spot),
			compactSpots:      make(map[int][]*Spot),
			handiSpots:        make([]*Spot, 6, 6),
			spots:             make(map[int][]*Spot)},
			[]ParkingEntry{
				{isHandicap: false, isCarpool: false, isCompact: true, entryLevel: 0,
					// id|level|spotNumber|label|isOccupied|isHandicap|isCarpool|isCompact
					expected: Spot{209, 0, 9, "L-9", false, false, false, true},
				},
				{true, false, false, 5,
					Spot{201, 0, 1, "L-1", false, true, false, false},
				},
				{false, false, false, 0,
					Spot{210, 0, 10, "L-10", false, false, false, false},
				},
				{false, true, false, -1,
					Spot{301, -1, 1, "A-1", false, false, true, false},
				},
			},
		},
	}
	reg = &CityRegulation{}

	for i, p := range tests {
		var bdr = NewBuilder(p.capacity, p.levels, p.groundLevels, reg)
		var obj = bdr.parking
		var tst = fmt.Sprintf("expecting %+v", *(p.expected))
		t.Logf("Test %2d: %v\n", i+1, obj)
		assert.Equal(t, p.expected, obj, tst)

		bdr.Construct()

		for ix, test := range p.entries {
			var val = obj.FindAvailableSpot(
				test.isHandicap, test.isCarpool, test.isCompact, test.entryLevel)
			var msg = fmt.Sprintf("entry:%v => (%+v)", &test, &(test.expected))

			t.Logf("%s", obj.SprintMap())
			t.Logf("Test %2d.%02d: %v\n", i+1, ix+1, msg)
			assert.Equal(t, test.expected, *val, msg)

			obj.Park(val)
		}
	}
}
