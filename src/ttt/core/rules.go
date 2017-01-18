package core

import (
	"ttt/core/grid"
	"ttt/core/marks"
)

type Rules struct{}

func (rule Rules) IsOver(tttboard Board, markOne, markTwo marks.Mark) bool {
	return rule.isAWin(tttboard, markOne, markTwo) ||
		rule.isADraw(tttboard, markOne, markTwo)
}

func (rule Rules) Winner(tttboard Board, markOne, markTwo marks.Mark) marks.Mark {
	for _, line := range tttboard.winningPositions() {
		if rule.winFor(line, markOne) {
			return markOne
		} else if rule.winFor(line, markTwo) {
			return markTwo
		}
	}
	return ""
}

func (rule Rules) winFor(line grid.Line, mark marks.Mark) bool {
	return line.All(matches(mark))
}

func (rule Rules) isAWin(tttboard Board, markOne, markTwo marks.Mark) bool {
	for _, line := range tttboard.winningPositions() {
		if rule.winFor(line, markOne) || rule.winFor(line, markTwo) {
			return true
		}
	}
	return false
}

func (rule Rules) isADraw(tttboard Board, markOne, markTwo marks.Mark) bool {
	if tttboard.isFull() && !rule.isAWin(tttboard, markOne, markTwo) {
		return true
	}
	return false
}

func matches(mark marks.Mark) func(marks.Mark) bool {
	return func(str marks.Mark) bool {
		return str == mark
	}
}
