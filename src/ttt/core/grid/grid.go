package grid

import "ttt/core/marks"

type Grid struct {
	cells []marks.Mark
}

func NewGrid(size int) Grid {
	cells := make([]marks.Mark, size)
	return Grid{cells}
}

func NewPopulatedGrid(cells []marks.Mark) Grid {
	return Grid{cells}
}

func (grid Grid) Cells() []marks.Mark {
	return grid.cells
}

func (grid Grid) Any(mark marks.Mark) bool {
	for _, cell := range grid.cells {
		if cell == mark {
			return true
		}
	}
	return false
}

func (grid Grid) Mark(move int, mark marks.Mark) Grid {
	cells := make([]marks.Mark, 9)
	copy(cells, grid.Cells())
	cells[move] = mark
	return NewPopulatedGrid(cells)
}

func (grid Grid) split(delimiter int) []Line {
	return grid.splitLines(delimiter)
}

func (grid Grid) reverseSplit(delimiter int) []Line {
	var reversedLines []Line
	for _, line := range grid.splitLines(delimiter) {
		reversedLines = append(reversedLines, line.reverse())
	}
	return reversedLines
}

func (grid Grid) transpose(delimiter int) []Line {
	var transposedGrid []Line
	rows := grid.split(delimiter)
	for i := 0; i < delimiter; i++ {
		transposedGrid = append(transposedGrid, grid.getCellsAt(rows, i))
	}
	return transposedGrid
}

func (grid Grid) getCellsAt(lines []Line, position int) Line {
	newLine := NewLine()
	for _, line := range lines {
		newLine = newLine.addCell(line.at(position))
	}
	return newLine
}

func (grid Grid) splitCells(start, end int) Line {
	splitCells := NewLine()
	for _, cell := range grid.cells[start:end] {
		splitCells = splitCells.addCell(cell)
	}
	return splitCells
}

func (grid Grid) splitLines(delimiter int) []Line {
	var splitGrid []Line
	for i := 0; i < len(grid.cells); i = i + delimiter {
		splitGrid = append(splitGrid, grid.splitCells(i, i+delimiter))
	}
	return splitGrid
}
