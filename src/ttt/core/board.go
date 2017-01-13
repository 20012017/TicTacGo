package core

import (
	"math"
	"ttt/core/grid"
)

const empty string = ""

type Board struct {
	grid grid.Grid
	size int
}

func NewBoard(size int) Board {
	grid := grid.NewGrid(size)
	return Board{grid, size}
}

func NewMarkedBoard(cells []string) Board {
	return Board{grid.NewPopulatedGrid(cells), len(cells)}
}

func (board Board) Grid() grid.Grid {
	return board.grid
}

func (board Board) Size() int {
	return board.size
}

func (board Board) MarkAt(index int) string {
	return board.cells()[index]
}

func (board Board) AvailableMoves() []int {
	availableMoves := []int{}
	for index, cell := range board.grid.Cells() {
		if cell == "" {
			availableMoves = append(availableMoves, index)
		}
	}
	return availableMoves
}

func (board Board) PlaceMark(cell int, mark string) Board {
	updatedGrid := board.updateCells(cell, mark)
	return NewMarkedBoard(updatedGrid)
}

func (board Board) rowLength() int {
	return int(math.Sqrt(float64(board.size)))
}

func (board Board) isFull() bool {
	return !board.grid.Any(empty)
}

func (board Board) winningPositions() []grid.Line {
	return grid.NewWinningPositions(board.grid, board.rowLength()).All
}

func (board Board) CountMarks() int {
	count := 0
	for _, mark := range board.cells() {
		if mark != empty {
			count++
		}
	}
	return count
}

func (board Board) updateCells(cell int, mark string) []string {
	grid := board.grid.Mark(cell, mark)
	return grid.Cells()
}

func (board Board) cells() []string {
	return board.grid.Cells()
}
