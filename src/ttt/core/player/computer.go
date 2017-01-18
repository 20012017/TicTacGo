package player

import (
	"ttt/core"
	"ttt/core/marks"
)

type Computer struct {
	mark    marks.Mark
	negamax *Negamax
}

func NewComputer(mark marks.Mark, negamax *Negamax) *Computer {
	return &Computer{mark, negamax}
}

func (computer Computer) Mark() marks.Mark {
	return computer.mark
}

func (computer *Computer) Move(board core.Board) (int, error) {
	return computer.negamax.move(board, computer.mark), nil
}
