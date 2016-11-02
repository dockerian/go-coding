package puzzle

// Find2ndLargest returns the second largest in an array
func Find2ndLargest(x []int) int {
	var largest, secondLargest int

	for i := 0; i < len(x); i++ {
		if x[i] > largest {
			if largest > secondLargest {
				secondLargest = largest
			}
			largest = x[i]
		} else if x[i] > secondLargest {
			secondLargest = x[i]
		}
	}

	return secondLargest
}
