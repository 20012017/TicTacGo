package ttt

const rowLength, columnWidth, firstCell int = 3, 1, 0

type WinningPositions struct {
	rows, columns, diagonals []Line
}

func NewWinningPositions(grid Grid) WinningPositions {
	return WinningPositions{
		grid.split(3),
		grid.transpose(3),
		diagonals(grid),
	}
}

func diagonals(grid Grid) []Line {
	left := getDiagonal(grid.split(3))
	right := getDiagonal(grid.reverseSplit(3))
	return []Line{left, right}
}

func (winningPostions WinningPositions) all() []Line {
	allPositions := []Line{}
	allPositions = getAllPositions(allPositions, winningPositions.rows)
	allPositions = getAllPositions(allPositions, winningPositions.columns)
	allPositions = getAllPositions(allPositions, winningPositions.diagonals)
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
