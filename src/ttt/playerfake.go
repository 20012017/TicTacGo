package ttt

import "errors"

type PlayerFake struct {
	mark        string
	currentMove int
	moves       []int
}

func newPlayerFake(mark string, currentMove int, moves ...int) *PlayerFake {
	return &PlayerFake{mark, currentMove, moves}
}

func (player *PlayerFake) Mark() string {
	return player.mark
}

func (player *PlayerFake) move(board Board) (int, error) {
	move := player.moves[player.currentMove]
	player.currentMove = player.currentMove + 1
	return move, player.setError(move)
}

func (player *PlayerFake) setError(move int) error {
	if move > 9 || move < 0 {
		return errors.New("Out of bounds")
	}
	return nil
}
