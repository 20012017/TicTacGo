package grid

import "ttt/core/marks"

type Line struct {
	cells []marks.Mark
}

func NewLine(cells ...marks.Mark) Line {
	return Line{cells}
}

func (line Line) All(condition func(marks.Mark) bool) bool {
	for _, cell := range line.cells {
		if !condition(cell) {
			return false
		}
	}
	return true
}

func (line Line) addCell(cell marks.Mark) Line {
	cells := line.cells
	cells = append(cells, cell)
	return Line{cells}
}

func (line Line) reverse() Line {
	return line.reverseEachElement(len(line.cells)-1, line.cells)
}

func (line Line) at(position int) marks.Mark {
	return line.cells[position]
}

func (line Line) reverseEachElement(reverseIndex int, row []marks.Mark) Line {
	reversedLine := NewLine()
	for i := 0; i < len(row); i++ {
		reversedLine = reversedLine.addCell(row[reverseIndex])
		reverseIndex--
	}
	return reversedLine
}
