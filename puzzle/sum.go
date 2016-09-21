package puzzle

import (
	u "github.com/dockerian/go-coding/utils"
)

// FindMatchedSum returns index in input array of two items match to the sum.
// See: https://leetcode.com/problems/two-sum/
// Problems: Given an array of integers, find indices of the two numbers
//           such that they add up to a specific sum
// Keywords: array, hash, sum
func FindMatchedSum(inputs []int, sum int) (int, int) {
	hashmap := make(map[int]int)

	u.Debug("- BEGIN inputs%+v\n", inputs)
	for i := 0; i < len(inputs); i++ {
		var value = inputs[i]
		var target = sum - value
		if index, ok := hashmap[target]; ok {
			u.Debug("- FOUND inputs[%v] == %v, in hash %+v\n\n", i, value, hashmap)
			return index, i
		}

		hashmap[value] = i
		u.Debug("- ADDED inputs[%v] == %v, to hash %+v\n", i, value, hashmap)
	}

	u.Debug("\n")
	return -1, -1
}
