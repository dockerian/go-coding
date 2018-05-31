// Package demo :: chan.go
package demo

import (
	u "github.com/dockerian/go-coding/utils"
)

// GetMergedChannels merges two channels to one.
// See https://medium.com/justforfunc/why-are-there-nil-channels-in-go-9877cc0b2308
// Notes: in common pattern, each channel (e.g. `ch` as `chan<- int`)
//   - receives data in a lengthy producer (e.g. `ch <- v`)
//   - closes at the end of the producer
func GetMergedChannels(ch1, ch2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for ch1 != nil || ch2 != nil {
			select {
			case v, ok := <-ch1:
				if !ok { // checking if ch1 is closed
					ch1 = nil // prevent from busy loop on closed channel which never blocks but nil will
					// u.Debug("channel ch1 (%v) is closed.\n", &ch1)
					continue
				}
				out <- v
			case v, ok := <-ch2:
				if !ok { // checking if ch2 is closed
					ch2 = nil // prevent from busy loop on closed channel which never blocks but nil will
					// u.Debug("channel ch2 (%v) is closed.\n", &ch2)
					continue
				}
				out <- v
			}
		}
	}()
	return out
}

// MergeChannels merges inputs to one out channel
func MergeChannels(out chan float32, inputs []chan float32) {
	var a, b chan float32
	switch len(inputs) {
	case 2:
		b = inputs[1]
		fallthrough
	case 1:
		a = inputs[0]
	case 0:
	default:
		a = make(chan float32)
		b = make(chan float32)
		half := len(inputs) / 2
		go MergeChannels(a, inputs[:half])
		go MergeChannels(b, inputs[half:])
	}

	mergeChan(out, a, b)
}

func mergeChan(out chan<- float32, a, b <-chan float32) {
	for a != nil || b != nil {
		select {
		case v, ok := <-a:
			if !ok {
				a = nil
				u.Debug("channel a (%v) is closed.\n", &a)
				continue
			}
			out <- v
		case v, ok := <-b:
			if !ok {
				b = nil
				u.Debug("channel b (%v) is closed.\n", &b)
				continue
			}
			out <- v
		}
	}
	close(out)
}
