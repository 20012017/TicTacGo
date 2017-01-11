package ttt

type Player interface {
	Mark() string

	move(board Board) (int, error)
}
