package puzzle

// Queens type
type Queens [][]bool

// String func for Queens
func (q *Queens) String() string {
	str := ""
	for _, row := range *q {
		for _, v := range row {
			if v {
				str += "Q "
			} else {
				str += "_ "
			}
		}
		str += "\n"
	}
	return str
}

// PlaceQueens func
// Place N queens on NxN chessboard so that no two of them attack each other
// See https://developers.google.com/optimization/puzzles/queens
func PlaceQueens(n int) []Queens {
	var queens []Queens
	return queens
}
