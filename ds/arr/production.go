// Package arr :: production.go
package arr

// GetProducts returns a new array where the value at any index has the products
// of all the other numbers in a given integer array.
func GetProducts(numbers []int) []int {
	siz := len(numbers)
	result := make([]int, siz)

	b := 1 // keep production before any index
	for x := 0; x < siz; x++ {
		result[x] = b
		b *= numbers[x]
	}

	p := 1 // keep production after the index
	for x := siz - 1; x >= 0; x-- {
		result[x] *= p
		p *= numbers[x]
	}

	return result
}
