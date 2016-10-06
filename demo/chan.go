package demo

import (
	u "github.com/dockerian/go-coding/utils"
)

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
