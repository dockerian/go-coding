// +build all demo chan test

package demo

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestMergeChannels tests MergeChannels
func TestMergeChannels(t *testing.T) {
	data := []float32{}
	cins := []chan float32{
		make(chan float32),
		make(chan float32),
		make(chan float32),
		make(chan float32),
		make(chan float32),
	}
	cout := make(chan float32)
	size := len(cins)

	part := 1
	for f := size; f > 0; f = f / 10 {
		part *= 10
	}

	for i, c := range cins {
		// create and start a goroutine
		go func(x int, cx chan float32) {
			for n := 1; n <= size; n++ {
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				v := r.Int63n(50) + 75 // random range between [75, 125)
				timeout := time.Duration(v)
				time.Sleep(timeout * time.Millisecond)
				cx <- float32(x+1) + float32(n)/float32(part)
			}
			close(cx)
		}(i, c)
	}

	// start a goroutine to merge
	go MergeChannels(cout, cins)

	done := false
	for !done {
		select {
		case v, ok := <-cout:
			if ok {
				t.Logf("cout gets [chan.seq] %v\n", v)
				// Note: This only to demo data can append when it starts with []
				data = append(data, v)
				continue // for loop
			}
			done = true
			// Note: This only to demo `goto` statement
			goto assertion
		}
	}

assertion:
	assert.Equal(t, size*size, len(data))
}
