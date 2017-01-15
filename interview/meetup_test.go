// +build all interview map meetup graph test

package interview

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	u "github.com/dockerian/go-coding/utils"
	"github.com/stretchr/testify/assert"
)

// MeetupTestCase struct
type MeetupTestCase struct {
	inputs   []*Point
	expected *Point
}

// BenchmarkMeetup benchmarks on func findMeetupPoint
func BenchmarkMeetup(b *testing.B) {
	u.DebugOff()

	b.Run("findMeetupPoint-without-cache", func(b *testing.B) {
		b.Logf("Benchmark findMeetupPoint (no cache)\n")
		benchmarkMeetup(b, false, false)
	})

	b.Run("findMeetupPoint-with-cache", func(b *testing.B) {
		b.Logf("Benchmark findMeetupPoint (w/ cache)\n")
		benchmarkMeetup(b, true, false)
	})

	b.Run("findMeetupPoint-by-average", func(b *testing.B) {
		b.Logf("Benchmark findMeetupPointByMidPoint\n")
		benchmarkMeetup(b, true, true)
	})

	u.DebugReset()
}

func benchmarkMeetup(b *testing.B, cache bool, useAvg bool) {
	var x = 100
	tests := make([]*Point, 0, x)
	for i := 0; i < x; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		x := r.Float64() * float64(x)
		y := r.Float64() * float64(x)
		z := r.Float64() * float64(x)
		tests = append(tests, &Point{x, y, z})
	}

	b.ResetTimer()
	if useAvg {
		for n := 0; n < b.N; n++ {
			findMeetupPointByMidPoint(tests)
		}
	} else {
		for n := 0; n < b.N; n++ {
			findMeetupPoint(tests, cache)
		}
	}
}

// TestMeetup tests func FindMeetupPoint
func TestMeetup(t *testing.T) {
	var tests = []MeetupTestCase{
		{[]*Point{
			&Point{0.0, 0.0, 0.0},
			&Point{1.1, 1.1, 1.1},
			&Point{2.2, 2.2, 2.2},
			&Point{3.3, 3.3, 3.3},
			&Point{4.4, 4.4, 4.4},
			&Point{5.5, 5.5, 5.5},
		}, &Point{2.2, 2.2, 2.2}},
		{[]*Point{
			&Point{0.0, 0.0, 0.0},
			&Point{-1.0, -2.0, 3.0},
			&Point{2.0, 4.0, 6.0},
			&Point{3.0, -5.0, 7.0},
			&Point{4.0, 3.0, 2.0},
			&Point{-55.0, -44.0, 3.0},
			&Point{11.0, 22.0, 3.0},
			&Point{-5.0, 3.0, 1.0},
			&Point{111.0, 2.2, 3.3},
		}, &Point{4.0, 3.0, 2.0}},
		{[]*Point{}, nil},
	}
	for index, test := range tests {
		var val = findMeetupPoint(test.inputs, true)
		var msg = fmt.Sprintf("meetup @ %+v for %+v", test.expected, test.inputs)
		t.Logf("Test %2d: %v\n", index+1, msg)
		assert.Equal(t, test.expected, val, msg)
		// var dst = findMeetupPointByMidPoint(test.inputs)
		// assert.Equal(t, dst, val, msg)
	}
}
