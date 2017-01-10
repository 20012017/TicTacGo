package board

type Line struct {
	cells []string
}

func newLine(cells ...string) Line {
	return Line{cells}
}

func (line Line) addCell(cell string) Line {
	cells := line.cells
	cells = append(cells, cell)
	return Line{cells}
}

func (line Line) reverse() Line {
	return line.reverseEachElement(len(line.cells)-1, line.cells)
}

func (line Line) at(position int) string {
	return line.cells[position]
}

func (line Line) All(condition func(string) bool) bool {
	for _, cell := range line.cells {
		if !condition(cell) {
			return false
		}
	}
	return true
}

func (line Line) reverseEachElement(reverseIndex int, row []string) Line {
	reversedLine := newLine()
	for i := 0; i < len(row); i++ {
		reversedLine = reversedLine.addCell(row[reverseIndex])
		reverseIndex--
	}
	return reversedLine
}
