package player

import (
	"ttt/core"
	"ttt/core/marks"
)

const winScore, lossScore, drawScore int = 10, -10, 0

type Negamax struct {
	scores              map[int]int
	rules               *core.Rules
	colour, alpha, beta int
}

func NewNegamax(rules *core.Rules) *Negamax {
	return &Negamax{make(map[int]int), rules, 1, -9999, 9999}
}

func (negamax *Negamax) move(board core.Board, mark string) int {
	opponentMark := negamax.opponentMark(mark)
	alpha, beta, colour := negamax.alpha, negamax.beta, negamax.colour
	_, move := negamax.negamax(board, alpha, beta, colour, mark, opponentMark)
	return move
}

func (negamax *Negamax) opponentMark(mark string) string {
	if mark == marks.X {
		return marks.O
	}
	return marks.X
}

func (negamax *Negamax) currentMark(color int, mark, opponentMark string) string {
	if color == negamax.colour {
		return mark
	}
	return opponentMark
}

func (negamax Negamax) negamax(board core.Board, alpha, beta, colour int, mark, opponentMark string) (int, int) {
	if negamax.rules.IsOver(board, mark, opponentMark) {
		return colour * negamax.score(board, mark, opponentMark), -1
	}
	var move int
	for _, cell := range board.AvailableMoves() {
		currentMark := negamax.currentMark(colour, mark, opponentMark)
		board := board.PlaceMark(cell, currentMark)
		score, _ := negamax.negamax(board, -beta, -alpha, -colour, mark, opponentMark)
		score = -score
		if score >= beta {
			return score, cell
		}
		if score > alpha {
			alpha, move = score, cell
		}
	}
	return alpha, move
}

func (negamax Negamax) score(board core.Board, mark, opponentMark string) (score int) {
	winner := negamax.rules.Winner(board, mark, opponentMark)
	if winner == mark {
		return winScore
	} else if winner == opponentMark {
		return lossScore
	}
	return drawScore
}
