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

func (grid Grid) Split(delimiter int) [][]string {
	splitGrid := [][]string{}
	for i := 0; i < len(grid.cells); i = i + delimiter {
		splitGrid = append(splitGrid, grid.splitCells(i, i+delimiter))
	}
	return splitGrid
}

func (grid Grid) Transpose(delimiter int) [][]string {
	transposedGrid, rows := [][]string{}, grid.Split(delimiter)
	for i := 0; i < delimiter; i++ {
		transposedGrid = append(transposedGrid, grid.getCellsAt(rows, i))
	}
	return transposedGrid
}

func (grid Grid) Reverse() []string {
	reverseIndex, index, startIndex := len(grid.cells)-1, 0, 0
	return reverseAllElements(index, reverseIndex, startIndex)
}

func reverseAllElements(index, reverseIndex, startIndex int) []string {
	reversedGrid := make([]string, len(grid.cells))
	for reverseIndex >= 0 {
		reversedGrid[index] = grid.cells[reverseIndex]
		reverseIndex--
		index++
	}
	return reversedGrid
}

func (grid Grid) getCellsAt(rows [][]string, position int) []string {
	cells := []string{}
	for _, row := range rows {
		cells = append(cells, row[position])
	}
	return cells
}

func (grid Grid) splitCells(start, end int) []string {
	return grid.cells[start:end]
}
