package ttt

const rowLength, columnWidth, firstCell int = 3, 1, 0

type WinningPositions struct {
	rows, columns, diagonals [][]string
}

func NewWinningPositions(cells []string) WinningPositions {
	return WinningPositions{
		rows(cells),
		columns(cells),
		diagonals(cells),
	}
}

func rows(cells []string) [][]string {
	return getRows(firstCell, cells)
}

func columns(cells []string) [][]string {
	return getColumns(firstCell, cells)
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

func getRows(rowStart int, cells []string) [][]string {
	rows := [][]string{}
	for rowStart < (rowLength * rowLength) {
		row := getCellRange(cells, rowStart, rowStart+(rowLength-1), columnWidth)
		rows = append(rows, row)
		rowStart = rowStart + rowLength
	}
	return rows
}

func getColumns(columnStart int, cells []string) [][]string {
	columns := [][]string{}
	for columnStart < rowLength {
		column := getCellRange(cells, columnStart, columnStart+rowLength*2, rowLength)
		columns = append(columns, column)
		columnStart = columnStart + columnWidth
	}
	return columns
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
