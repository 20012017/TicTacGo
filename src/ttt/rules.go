package ttt

import "ttt/board"

type rules struct{}

func (rule rules) isAWin(board board.Board, markOne, markTwo string) bool {
	for _, line := range board.WinningPositions() {
		if line.All(matches(markOne)) || line.All(matches(markTwo)) {
			return true
		}
	}
	return false
}

func (rule rules) isADraw(board board.Board, markOne, markTwo string) bool {
	if board.IsFull() && !rule.isAWin(board, markOne, markTwo) {
		return true
	}
	return false
}

func matches(mark string) func(string) bool {
	return func(str string) bool {
		return str == mark
	}
}
