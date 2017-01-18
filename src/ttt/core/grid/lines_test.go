package grid

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core/marks"
)

func TestLineCreatedWithCells(t *testing.T) {
	line := NewLine("1", "2", "3")
	expectedCells := []marks.Mark{"1", "2", "3"}

	assert.Equal(t, expectedCells, line.cells)
}

func TestEmptyNewLine(t *testing.T) {
	line := NewLine()

	assert.Equal(t, []marks.Mark(nil), line.cells)
}

func TestAddCellToLine(t *testing.T) {
	line := NewLine("1")
	expectedLine := NewLine("1", "2")

	assert.Equal(t, expectedLine, line.addCell("2"))
}

func TestLinesReverse(t *testing.T) {
	line := NewLine("1", "2", "3")
	reversedLine := NewLine("3", "2", "1")

	assert.Equal(t, reversedLine, line.reverse())
}

func TestLinesAt(t *testing.T) {
	line := NewLine("1")

	assert.Equal(t, marks.Mark("1"), line.at(0))
}

func TestLinesAll(t *testing.T) {
	isAOne := func(str marks.Mark) bool {
		return str == "1"
	}
	sameLine := NewLine("1", "1", "1")
	differentLine := NewLine("1", "2", "3")

	assert.True(t, sameLine.All(isAOne))
	assert.False(t, differentLine.All(isAOne))
}
