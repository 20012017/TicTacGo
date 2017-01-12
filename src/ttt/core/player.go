package core

type Player interface {
	Mark() string

	Move(tttboard TTTBoard) (int, error)
}
