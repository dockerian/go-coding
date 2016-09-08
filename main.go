package main

import (
	"fmt"
	"math/rand"
)

var packageVar = 0

func main() {
	var result = getSum(rand.Intn(10))
	fmt.Println("Result:", result)

}

func getXplus(n int) int {
	packageVar += n
	return packageVar
}

func getSum(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += getXplus(i)
	}
	return sum
}
