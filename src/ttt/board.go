package ttt

import "math"

type Board struct {
	grid Grid
	size int
}

func NewBoard(size int) Board {
	grid := NewGrid(size)
	return Board{grid, size}
}

func newMarkedBoard(cells []string) Board {
	return Board{NewPopulatedGrid(cells), len(cells)}
}

func (board Board) PlaceMark(cell int, mark string) Board {
	updatedGrid := updateCells(board.grid, cell, mark)
	return newMarkedBoard(updatedGrid)
}

func (board Board) MarkAt(index int) string {
	return board.grid.cells[index]
}

func (board Board) RowLength() int {
	return int(math.Sqrt(float64(board.size)))
}

func updateCells(grid Grid, cell int, mark string) []string {
	cells := grid.cells
	cells[cell] = mark
	return cells
}
