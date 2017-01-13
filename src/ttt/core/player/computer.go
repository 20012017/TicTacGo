package player

import (
	"math"
	"ttt/core"
)

type Computer struct {
	mark, opponentMark string
	rules              *core.Rules
	scores             map[int]int
}

func NewComputer(mark, opponentMark string, rules *core.Rules) *Computer {
	return &Computer{mark, opponentMark, rules, make(map[int]int)}
}

func (computer Computer) Mark() string {
	return computer.mark
}

func (computer Computer) firstMove(board core.Board) (bool, int) {
	markCount := board.CountMarks()
	if markCount == 0 {
		return true, 0
	}
	if markCount == 1 {
		return true, 4
	}
	return false, 0
}

func (computer Computer) getMark(color int) string {
	if color == -1 {
		return computer.opponentMark
	}
	return computer.mark
}

func (computer *Computer) Move(board core.Board) (int, error) {
	firstMove, move := computer.firstMove(board)
	if firstMove {
		return move, nil
	}
	computer.negamax(board, 0, 1)
	move = computer.bestMove()
	return move, nil
}

func (computer Computer) bestMove() int {
	bestMove := -999999
	bestScore := -999999
	for move, score := range computer.scores {
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}
	return bestMove
}

func (computer *Computer) negamax(board core.Board, depth, colour int) int {
	if computer.rules.IsOver(board, computer.mark, computer.opponentMark) {
		return colour * computer.score(board, depth)
	}
	bestValue := -9999999
	for _, move := range board.AvailableMoves() {
		mark := computer.getMark(colour)
		board := board.PlaceMark(move, mark)
		value := -computer.negamax(board, depth+1, -colour)
		bestValue = int(math.Max(float64(bestValue), float64(value)))
		if depth == 0 {
			computer.scores[move] = value
		}
	}
	return bestValue
}

func (computer Computer) score(board core.Board, depth int) int {
	winner := computer.rules.Winner(board, computer.mark, computer.opponentMark)
	if winner == computer.mark {
		return 10 / depth
	} else if winner == computer.opponentMark {
		return -10 / depth
	}
	return 0
}
