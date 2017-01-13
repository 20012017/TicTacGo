package player

import (
	"ttt/core"
)

type Computer struct {
	mark    string
	negamax *Negamax
}

func NewComputer(mark string, negamax *Negamax) *Computer {
	return &Computer{mark, negamax}
}

func (computer Computer) Mark() string {
	return computer.mark
}

func (computer *Computer) Move(board core.Board) (int, error) {
	return computer.negamax.move(board, computer.mark), nil
}
