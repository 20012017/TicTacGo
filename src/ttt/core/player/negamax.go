package player

import (
	"math"
	"ttt/core"
)

type Negamax struct {
	scores map[int]int
	rules  *core.Rules
}

func NewNegamax(rules *core.Rules) *Negamax {
	return &Negamax{make(map[int]int), rules}
}

func (negamax *Negamax) move(board core.Board, mark string) int {
	firstMove, move := negamax.firstMove(board)
	if firstMove {
		return move
	}
	opponentMark := negamax.opponentMark(mark)
	negamax.negamax(board, 0, 1, mark, opponentMark)
	move = negamax.bestMove()
	return move
}

func (negamax Negamax) opponentMark(mark string) string {
	if mark == "X" {
		return "O"
	}
	return "X"
}

func (negamax Negamax) firstMove(board core.Board) (bool, int) {
	markCount := board.CountMarks()
	if markCount == 0 {
		return true, 0
	}
	if markCount == 1 {
		return true, 4
	}
	return false, 0
}

func (negamax Negamax) getMark(color int, mark, opponentMark string) string {
	if color == -1 {
		return opponentMark
	}
	return mark
}

func (negamax Negamax) bestMove() int {
	bestMove := -999999
	bestScore := -999999
	for move, score := range negamax.scores {
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}
	return bestMove
}

func (negamax *Negamax) negamax(board core.Board, depth, colour int, mark, opponentMark string) int {
	if negamax.rules.IsOver(board, mark, opponentMark) {
		return colour * negamax.score(board, depth, mark, opponentMark)
	}
	bestValue := -9999999
	for _, move := range board.AvailableMoves() {
		mark := negamax.getMark(colour, mark, opponentMark)
		board := board.PlaceMark(move, mark)
		value := -negamax.negamax(board, depth+1, -colour, mark, opponentMark)
		bestValue = int(math.Max(float64(bestValue), float64(value)))
		if depth == 0 {
			negamax.scores[move] = value
		}
	}
	return bestValue
}

func (negamax Negamax) score(board core.Board, depth int, mark, opponentMark string) int {
	winner := negamax.rules.Winner(board, mark, opponentMark)
	if winner == mark {
		return 10 / depth
	} else if winner == opponentMark {
		return -10 / depth
	}
	return 0
}
