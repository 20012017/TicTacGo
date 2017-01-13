package grid

const rowLength, columnWidth, firstCell int = 3, 1, 0

type WinningPositions struct {
	rows, columns, diagonals, All []Line
}

func NewWinningPositions(grid Grid, delimiter int) WinningPositions {
	return WinningPositions{
		grid.split(delimiter),
		grid.transpose(delimiter),
		diagonals(grid),
		all(grid),
	}
}

func diagonals(grid Grid) []Line {
	left := getDiagonal(grid.split(3))
	right := getDiagonal(grid.reverseSplit(3))
	return []Line{left, right}
}

func all(grid Grid) []Line {
	allPositions := getPositions([]Line{}, grid.split(3))
	allPositions = getPositions(allPositions, grid.transpose(3))
	return getPositions(allPositions, diagonals(grid))
}

func getPositions(allPositions, positions []Line) []Line {
	for _, position := range positions {
		allPositions = append(allPositions, position)
	}
	return allPositions
}

func getDiagonal(lines []Line) Line {
	diagonal := NewLine()
	for index, line := range lines {
		diagonal = diagonal.addCell(line.at(index))
	}
	return diagonal
}
