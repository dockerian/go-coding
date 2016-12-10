package puzzle

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
