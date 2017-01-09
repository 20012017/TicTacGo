package ttt

type Board struct {
	cells []string
}

func NewBoard() Board {
	return Board{make([]string, 9)}
}

func newMarkedBoard(cells []string) Board {
	return Board{cells}
}

func (board Board) Size() int {
	return len(board.cells)
}

func (board Board) PlaceMark(cell int, mark string) Board {
	return newMarkedBoard(updateCells(board.cells, cell, mark))
}

func (board Board) MarkAt(index int) string {
	return board.cells[index]
}

func updateCells(cells []string, cell int, mark string) []string {
	cells[cell] = mark
	return cells
}
