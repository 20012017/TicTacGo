package ttt

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

func (grid Grid) Split(delimiter int) line {
	splitGrid := [][]string{}
	for i := 0; i < len(grid.cells); i = i + delimiter {
		splitGrid = append(splitGrid, grid.splitCells(i, i+delimiter))
	}
	return line{splitGrid}
}

func (grid Grid) Transpose(delimiter int) [][]string {
	transposedGrid, rows := [][]string{}, grid.Split(delimiter)
	for i := 0; i < delimiter; i++ {
		transposedGrid = append(transposedGrid, grid.getCellsAt(rows, i))
	}
	return transposedGrid
}

func (grid Grid) getCellsAt(rows line, position int) []string {
	cells := []string{}
	for _, row := range rows.cells {
		cells = append(cells, row[position])
	}
	return cells
}

func (grid Grid) splitCells(start, end int) []string {
	return grid.cells[start:end]
}
