package core

type Player interface {
	Mark() string

	Move(tttboard Board) (int, error)
}
