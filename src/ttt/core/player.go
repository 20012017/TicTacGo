package core

type Player interface {
	Mark() string

	Move(board Board) (int, error)
}
