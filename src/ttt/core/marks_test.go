package core

type Mark struct {
	X, O, Empty string
}

func NewMarks(x, o, empty string) Mark {
	return Mark{x, o, empty}
}
