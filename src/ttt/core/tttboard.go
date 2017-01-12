package core

import (
	"math"
	"ttt/core/board"
)

const empty string = ""

type TTTBoard struct {
	grid board.Grid
	size int
}

func NewBoard(size int) TTTBoard {
	grid := board.NewGrid(size)
	return TTTBoard{grid, size}
}

func NewMarkedBoard(cells []string) TTTBoard {
	return TTTBoard{board.NewPopulatedGrid(cells), len(cells)}
}

func (tttboard TTTBoard) Grid() board.Grid {
	return tttboard.grid
}

func (tttboard TTTBoard) Size() int {
	return tttboard.size
}

func (tttboard TTTBoard) MarkAt(index int) string {
	return tttboard.cells()[index]
}

func (tttboard TTTBoard) placeMark(cell int, mark string) TTTBoard {
	updatedGrid := tttboard.updateCells(cell, mark)
	return NewMarkedBoard(updatedGrid)
}

func (tttboard TTTBoard) rowLength() int {
	return int(math.Sqrt(float64(tttboard.size)))
}

func (tttboard TTTBoard) isFull() bool {
	return !tttboard.grid.Any(empty)
}

func (tttboard TTTBoard) winningPositions() []board.Line {
	return board.NewWinningPositions(tttboard.grid, tttboard.rowLength()).All
}

func (tttboard TTTBoard) countMarks() int {
	count := 0
	for _, mark := range tttboard.cells() {
		if mark != empty {
			count++
		}
	}
	return count
}

func (tttboard TTTBoard) updateCells(cell int, mark string) []string {
	cells := tttboard.cells()
	cells[cell] = mark
	return cells
}

func (tttboard TTTBoard) cells() []string {
	return tttboard.grid.Cells()
}
