package ttt

type WinningPositions struct {
	rows, columns, diagonals [][]string
}

func NewWinningPositions(cells []string) WinningPositions {
	return WinningPositions{rows(cells), columns(cells), diagonals(cells)}
}

func rows(cells []string) [][]string {
	row := getRow
	return [][]string{row(0, cells), row(3, cells), row(6, cells)}
}

func columns(cells []string) [][]string {
	column := getColumn
	return [][]string{column(0, cells), column(1, cells), column(2, cells)}
}

func diagonals(cells []string) [][]string {
	return [][]string{getLeftDiagonal(cells), getRightDiagonal(cells)}
}

func getRow(rowStart int, cells []string) []string {
	return getCellRange(cells, rowStart, rowStart+2, 1)
}

func getColumn(columnStart int, cells []string) []string {
	return getCellRange(cells, columnStart, columnStart+6, 3)
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
