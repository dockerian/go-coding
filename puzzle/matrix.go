package puzzle

import (
	"fmt"
)

// SumMaxtrixDiagonal returns both diagonal sum of
// up-left to down-right and up-right to down-left
func SumMaxtrixDiagonal(matrix [][]int) (int64, int64) {
	var sum1, sum2 int64
	var size = len(matrix)
	for i := 0; i < size; i++ {
		j := size - i + 1
		sum1 += int64(matrix[i][i])
		sum2 += int64(matrix[i][j])
	}
	return sum1, sum2
}

// FindAjacent1s returns all coordinates "x,y" of adjacent 1s, where
// the position has adjacent 1s horizontally or vertically, from a matrix
// contains only 0s and 1s
func FindAjacent1s(matrix [][]int) []string {
	var rows = len(matrix)
	var maps = make([]string, 0)
	var hash = make(map[string]bool)

	checkAndAdd := func(x, y int) {
		key := fmt.Sprintf("%d,%d", x, y)
		if _, ok := hash[key]; !ok {
			maps = append(maps, key)
			hash[key] = true
		}
	}

	for x := 0; x < rows; x++ {
		var cols = len(matrix[x])
		for y := 0; y < cols; y++ {
			if matrix[x][y] == 1 {
				switch {
				case x > 0 && matrix[x-1][y] == 1:
					checkAndAdd(x, y)
					checkAndAdd(x-1, y)
				case y > 0 && matrix[x][y-1] == 1:
					checkAndAdd(x, y)
					checkAndAdd(x, y-1)
				case x < rows-1 && matrix[x+1][y] == 1:
					checkAndAdd(x, y)
					checkAndAdd(x+1, y)
				case y < cols-1 && matrix[x][y+1] == 1:
					checkAndAdd(x, y)
					checkAndAdd(x, y+1)
				}
			}
		}
	}
	return maps
}
