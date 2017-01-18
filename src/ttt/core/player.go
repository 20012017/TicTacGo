package core

import "ttt/core/marks"

type Player interface {
	Mark() marks.Mark

	Move(tttboard Board) (int, error)
}
