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

func (board Board) Rows() [][]string {
	row := board.getRow
	return [][]string{row(0), row(3), row(6)}
}

func (board Board) Columns() [][]string {
	column := board.getColumn
	return [][]string{column(0), column(1), column(2)}
}

func (board Board) Diagonals() [][]string {
	return [][]string{board.getLeftDiagonal(), board.getRightDiagonal()}
}

func updateCells(cells []string, cell int, mark string) []string {
	cells[cell] = mark
	return cells
}

func (board Board) getLeftDiagonal() []string {
	return board.getCellRange(0, 8, 4)
}

func (board Board) getRightDiagonal() []string {
	return board.getCellRange(2, 6, 2)
}

func (board Board) getColumn(columnStart int) []string {
	return board.getCellRange(columnStart, columnStart+6, 3)
}

func (board Board) getRow(rowStart int) []string {
	return board.getCellRange(rowStart, rowStart+2, 1)
}

func (board Board) getCellRange(rangeStart, rangeEnd, incrementer int) []string {
	cellRange := []string{}
	for rangeStart <= rangeEnd {
		cellRange = append(cellRange, board.cells[rangeStart])
		rangeStart = rangeStart + incrementer
	}
	return cellRange
}
