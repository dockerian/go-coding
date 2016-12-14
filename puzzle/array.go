package puzzle

// FindAvailableSpot returns first available spot from
// an array of occupied parking spot numbers in a parking garage which
// has parking spots 0..N (N could be infinite)
// note: in a design for real parking garage,
//    a) the available spots should be thread-safe to support concurrency
//    b) or to update first available info at any time before assigned
func FindAvailableSpot(numbers []int) int {
	hash := make(map[int]bool)
	// building a hash, taking additional O(n) space
	for _, spotID := range numbers {
		hash[spotID] = true
	}
	var n int
	// taking O(n) time to check against the hash table
	for {
		if n < 0 {
			return -1 // indicate overflow
		}
		if _, occupied := hash[n]; !occupied {
			return n
		}
		n++
	}
	// if inline operation is allowed (without using extra space),
	// another algorithm is to
	// 1) sort the inputs, taking O(n log(n))
	// 2) go thru in order to find the spot
}
