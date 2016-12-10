package puzzle

// FindBestParkingSpot returns first available spot from
// an array of occupied parking spot numbers in a parking garage which
// has parking spots 0..N (N could be infinite)
// note: in a design for real parking garage,
//    a) the available spots should be thread-safe to support concurrency
//    b) or to update first available info at any time before assigned
func FindBestParkingSpot(numbers []int) int {
	hash := make(map[int]bool)
	for _, spotID := range numbers {
		hash[spotID] = true
	}
	var n int
	for {
		if _, occupied := hash[n]; !occupied {
			return n
		}
		n++
	}
}
