package core

import (
	"math"
	"ttt/core/grid"
	"ttt/core/marks"
)

type Board struct {
	grid grid.Grid
	size int
}

func NewBoard(size int) Board {
	grid := grid.NewGrid(size)
	return Board{grid, size}
}

func NewMarkedBoard(cells []marks.Mark) Board {
	return Board{grid.NewPopulatedGrid(cells), len(cells)}
}

func (board Board) Grid() grid.Grid {
	return board.grid
}

func (board Board) Size() int {
	return board.size
}

func (board Board) MarkAt(index int) marks.Mark {
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

func (board Board) CountMarks() int {
	count := 0
	for _, mark := range board.cells() {
		if mark != marks.EMPTY {
			count++
		}
	}
	return count
}

func (board Board) PlaceMark(cell int, mark marks.Mark) Board {
	updatedGrid := board.updateCells(cell, mark)
	return NewMarkedBoard(updatedGrid)
}

func (board Board) rowLength() int {
	return int(math.Sqrt(float64(board.size)))
}

func (board Board) isFull() bool {
	return !board.grid.Any(marks.EMPTY)
}

func (board Board) winningPositions() []grid.Line {
	return grid.NewWinningPositions(board.grid, board.rowLength()).All
}

func (board Board) updateCells(cell int, mark marks.Mark) []marks.Mark {
	grid := board.grid.Mark(cell, mark)
	return grid.Cells()
}

func (board Board) cells() []marks.Mark {
	return board.grid.Cells()
}
