package ttt

type rules struct{}

func (rule rules) isAWin(board Board, markOne, markTwo string) bool {
	for _, line := range board.winningPositions() {
		if line.all(matches(markOne)) || line.all(matches(markTwo)) {
			return true
		}
	}
	return false
}

func (rule rules) isADraw(board Board, markOne, markTwo string) bool {
	if board.isFull() && !rule.isAWin(board, markOne, markTwo) {
		return true
	}
	return false
}

func matches(mark string) func(string) bool {
	return func(str string) bool {
		return str == mark
	}
}
