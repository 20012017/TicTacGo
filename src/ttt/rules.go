package ttt

type rules struct{}

func (rule rules) isAWin(board Board) bool {
	for _, line := range board.winningPositions() {
		if line.all(matches("X")) || line.all(matches("O")) {
			return true
		}
	}
	return false
}

func (rule rules) isADraw(board Board) bool {
	if board.isFull() && !rule.isAWin(board) {
		return true
	}
	return false
}

func matches(mark string) func(string) bool {
	return func(str string) bool {
		return str == mark
	}
}
