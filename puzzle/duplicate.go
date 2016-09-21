package puzzle

import (
	"fmt"

	u "github.com/dockerian/go-coding/utils"
)

// FindDuplicate looks up for one duplicated number in an integer array
// contains n numbers ranging from 0 to n-2.
// Note: There is exactly one number duplicated in the array.
func FindDuplicate(inputs []int) int {
	var sum1, sum2 int64
	var size = len(inputs)

	for i := 0; i < size-1; i++ {
		sum1 += int64(i)
	}
	for i := 0; i < size; i++ {
		sum2 += int64(inputs[i])
	}

	return int(sum2 - sum1)
}

// FindDuplicates returns a set of duplicated numbers in an integer array
// contains n numbers ranging from 0 to n-1.
// Note: There could be multiple duplicates in the array.
func FindDuplicates(inputs []int) ([]int, error) {
	var length = len(inputs)
	var result = make([]int, 0, length/2)
	var counts int
	u.Debug("\ninputs= %+v\n", inputs)
	for i := 0; i < length; i++ {
		val := inputs[i]
		if val != i {
			chk, idx := val, i
			for chk != idx {
				if chk >= length {
					err := fmt.Errorf("invalid number %v (at %v) in 0..%v array", chk, idx, length-1)
					return result, err
				}
				dat := inputs[chk]
				if dat >= length {
					err := fmt.Errorf("invalid number %v (at %v) in 0..%v array", dat, chk, length-1)
					return result, err
				}
				inputs[idx], inputs[chk] = dat, chk
				if dat != chk {
					u.Debug("\ninputs= %+v [i=%2d]\n", inputs, i)
				}
				idx = dat
				chk = inputs[idx]
			}
		}
	}
	for i := 0; i < length; i++ {
		val := inputs[i]
		if val != i && val == inputs[val] {
			result = append(result, val)
			counts++
		}
	}
	return result[0:counts], nil
}
