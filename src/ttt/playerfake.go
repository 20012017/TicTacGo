package ttt

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
	return move, nil
}
