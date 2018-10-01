// Package mathex :: gauss.go
package mathex

// GaussSum calculates sum of a series of integers from specific `start` with
// an `interval` in next `num` sequence (as total `num + 1`).
//
// Per the idea of Gauss formula, a sum of any sequencial numbers has pattern
//
//    ```
//    x, x+i, x+2i, ..., x+Ni
//    x+Ni, ..., x+2i, x+i, x
//    ```
//
// Adding above 2 lines, the sum would be `(N + 1) * (x + (x + N*i)) / 2`,
// and simplified as `(N+1) * (2*x + N*i) / 2`.
//
// Note: overflow needs be handled.
//
func GaussSum(start, interval, num uint16) uint64 {
	x := float64(start)
	i := float64(interval)
	n := float64(num)
	// TODO [jason_zhuyx]: possibility to handle overflow for uint32 inputs
	part1 := (n + 1) * x
	part2 := (n + 1) / 2 * n * i

	return uint64(part1 + part2)
}
