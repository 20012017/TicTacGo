package board

const rowLength, columnWidth, firstCell int = 3, 1, 0

type WinningPositions struct {
	rows, columns, diagonals, all []Line
}

func NewWinningPositions(grid Grid) WinningPositions {
	return WinningPositions{
		grid.split(3),
		grid.transpose(3),
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
	allPositions := []Line{}
	allPositions = getAllPositions(allPositions, grid.split(3))
	allPositions = getAllPositions(allPositions, grid.transpose(3))
	allPositions = getAllPositions(allPositions, diagonals(grid))
	return allPositions
}

func getAllPositions(allPositions, positions []Line) []Line {
	for _, position := range positions {
		allPositions = append(allPositions, position)
	}
	return allPositions
}

func getDiagonal(lines []Line) Line {
	diagonal := newLine()
	for index, line := range lines {
		diagonal = diagonal.addCell(line.at(index))
	}
	return diagonal
}
