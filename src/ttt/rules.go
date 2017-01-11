package ttt

type rules struct{}

func (rule rules) isAWin(board Board, markOne, markTwo string) bool {
	for _, line := range board.winningPositions() {
		if rule.winFor(line, markOne) || rule.winFor(line, markTwo) {
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

func (rule rules) winner(board Board, markOne, markTwo string) string {
	for _, line := range board.winningPositions() {
		if rule.winFor(line, markOne) {
			return markOne
		} else if rule.winFor(line, markTwo) {
			return markTwo
		}
	}
	return ""
}

func (rule rules) winFor(line Line, mark string) bool {
	return line.all(matches(mark))
}

func matches(mark string) func(string) bool {
	return func(str string) bool {
		return str == mark
	}
}
