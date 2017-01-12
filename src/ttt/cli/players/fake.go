package players

import (
	"errors"
	"ttt/core"
)

type Fake struct {
	mark        string
	currentMove int
	moves       []int
}

func NewFake(mark string, currentMove int, moves ...int) *Fake {
	return &Fake{mark, currentMove, moves}
}

func (player *Fake) Mark() string {
	return player.mark
}

func (player *Fake) Move(board core.TTTBoard) (int, error) {
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
