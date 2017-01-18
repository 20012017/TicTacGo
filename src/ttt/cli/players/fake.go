package players

import (
	"errors"
	"ttt/core"
	"ttt/core/marks"
)

type Fake struct {
	mark        marks.Mark
	currentMove int
	moves       []int
}

func NewFake(mark marks.Mark, currentMove int, moves ...int) *Fake {
	return &Fake{mark, currentMove, moves}
}

func (player *Fake) Mark() marks.Mark {
	return player.mark
}

func (player *Fake) Move(board core.Board) (int, error) {
	move := player.moves[player.currentMove]
	player.currentMove = player.currentMove + 1
	return move, player.setError(move)
}

func (player *Fake) setError(move int) error {
	if move > 9 || move < 0 {
		return errors.New("Out of bounds")
	}
	return nil
}
