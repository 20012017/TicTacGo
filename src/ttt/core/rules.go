package core

import "ttt/core/board"

type Rules struct{}

func (rule Rules) IsOver(tttboard TTTBoard, markOne, markTwo string) bool {
	return rule.isAWin(tttboard, markOne, markTwo) ||
		rule.isADraw(tttboard, markOne, markTwo)
}

func (rule Rules) Winner(tttboard TTTBoard, markOne, markTwo string) string {
	for _, line := range tttboard.winningPositions() {
		if rule.winFor(line, markOne) {
			return markOne
		} else if rule.winFor(line, markTwo) {
			return markTwo
		}
	}
	return ""
}

func (rule Rules) winFor(line board.Line, mark string) bool {
	return line.All(matches(mark))
}

func (rule Rules) isAWin(tttboard TTTBoard, markOne, markTwo string) bool {
	for _, line := range tttboard.winningPositions() {
		if rule.winFor(line, markOne) || rule.winFor(line, markTwo) {
			return true
		}
	}
	return false
}

func (rule Rules) isADraw(tttboard TTTBoard, markOne, markTwo string) bool {
	if tttboard.isFull() && !rule.isAWin(tttboard, markOne, markTwo) {
		return true
	}
	return false
}

func matches(mark string) func(string) bool {
	return func(str string) bool {
		return str == mark
	}
}
