package core

type Grid struct {
	cells []string
}

func NewGrid(size int) Grid {
	cells := make([]string, size)
	return Grid{cells}
}

func NewPopulatedGrid(cells []string) Grid {
	return Grid{cells}
}

func (grid Grid) Cells() []string {
	return grid.cells
}

func (grid Grid) split(delimiter int) []Line {
	return grid.splitLines(delimiter)
}

func (grid Grid) reverseSplit(delimiter int) []Line {
	reversedLines := []Line{}
	for _, line := range grid.splitLines(delimiter) {
		reversedLines = append(reversedLines, line.reverse())
	}
	return reversedLines
}

func (grid Grid) transpose(delimiter int) []Line {
	transposedGrid, rows := []Line{}, grid.split(delimiter)
	for i := 0; i < delimiter; i++ {
		transposedGrid = append(transposedGrid, grid.getCellsAt(rows, i))
	}
	return transposedGrid
}

func (grid Grid) getCellsAt(lines []Line, position int) Line {
	newLine := newLine()
	for _, line := range lines {
		newLine = newLine.addCell(line.at(position))
	}
	return newLine
}

func (grid Grid) splitCells(start, end int) Line {
	splitCells := newLine()
	for _, cell := range grid.cells[start:end] {
		splitCells = splitCells.addCell(cell)
	}
	return splitCells
}

func (grid Grid) splitLines(delimiter int) []Line {
	splitGrid := []Line{}
	for i := 0; i < len(grid.cells); i = i + delimiter {
		splitGrid = append(splitGrid, grid.splitCells(i, i+delimiter))
	}
	return splitGrid
}

func (grid Grid) any(str string) bool {
	for _, cell := range grid.cells {
		if cell == str {
			return true
		}
	}
	return false
}
