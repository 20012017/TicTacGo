package ttt

const rowLength, columnWidth, firstCell int = 3, 1, 0

type WinningPositions struct {
	rows, columns, diagonals [][]string
}

func NewWinningPositions(cells []string, grid Grid) WinningPositions {
	return WinningPositions{
		grid.Split(3),
		grid.Transpose(3),
		diagonals(cells),
	}
}

func diagonals(cells []string) [][]string {
	return [][]string{getLeftDiagonal(cells), getRightDiagonal(cells)}
}

func (winningPositions WinningPositions) All() [][]string {
	allPositions := [][]string{}
	allPositions = getAllPositions(allPositions, winningPositions.rows)
	allPositions = getAllPositions(allPositions, winningPositions.columns)
	allPositions = getAllPositions(allPositions, winningPositions.diagonals)
	return allPositions
}

func getAllPositions(allPositions [][]string, positions [][]string) [][]string {
	for _, row := range positions {
		allPositions = append(allPositions, row)
	}
	return allPositions
}

func getLeftDiagonal(cells []string) []string {
	return getCellRange(cells, 0, 8, 4)
}

func getRightDiagonal(cells []string) []string {
	return getCellRange(cells, 2, 6, 2)
}

func getCellRange(cells []string, rangeStart, rangeEnd, incrementer int) []string {
	cellRange := []string{}
	for rangeStart <= rangeEnd {
		cellRange = append(cellRange, cells[rangeStart])
		rangeStart = rangeStart + incrementer
	}
	return cellRange
}
