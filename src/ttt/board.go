package ttt

import "math"

const empty string = ""

type Board struct {
	grid Grid
	size int
}

func NewBoard(size int) Board {
	grid := NewGrid(size)
	return Board{grid, size}
}

func NewMarkedBoard(cells []string) Board {
	return Board{NewPopulatedGrid(cells), len(cells)}
}

func (board Board) placeMark(cell int, mark string) Board {
	updatedGrid := board.updateCells(board.grid, cell, mark)
	return NewMarkedBoard(updatedGrid)
}

func (board Board) markAt(index int) string {
	return board.grid.cells[index]
}

func (board Board) rowLength() int {
	return int(math.Sqrt(float64(board.size)))
}

func (board Board) isFull() bool {
	return !board.grid.any(empty)
}

func (board Board) winningPositions() []Line {
	return NewWinningPositions(board.grid).all
}

func (board Board) countMarks() int {
	count := 0
	for _, mark := range board.grid.cells {
		if mark != empty {
			count++
		}
	}
	return count
}

func (board Board) updateCells(grid Grid, cell int, mark string) []string {
	cells := grid.cells
	cells[cell] = mark
	return cells
}
