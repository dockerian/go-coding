package puzzle

// canWinNimGame checks who (opponent or I) will win the Nim game.
// See Nim game (https://leetcode.com/problems/nim-game/)
// There is a heap of chips on the table, each time one of players take turns to
// remove 1 to 3 chips. The one who removes the last stone will be the winner.
func canWinNimGame(n int, myTurn bool) bool {
	var opponentTurn = !myTurn
	if n > 50 {
		return myTurn && (n%4) != 0
	}
	if n <= 0 || n == 4 || n == 8 {
		return opponentTurn
	}
	if 0 < n && n < 4 {
		return myTurn
	}
	if 4 < n && n < 8 {
		return myTurn
	}

	if opponentTurn {
		return canWinNimGame(n-1, myTurn) &&
			canWinNimGame(n-2, myTurn) &&
			canWinNimGame(n-3, myTurn)
	}

	return canWinNimGame(n-1, myTurn) ||
		canWinNimGame(n-2, myTurn) ||
		canWinNimGame(n-3, myTurn)
}
