package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLineCreatedWithCells(t *testing.T) {
	line := newLine("1", "2", "3")
	expectedCells := []string{"1", "2", "3"}

	assert.Equal(t, expectedCells, line.cells)
}

func TestEmptyNewLine(t *testing.T) {
	line := newLine()

	assert.Equal(t, []string(nil), line.cells)
}

func TestAddCellToLine(t *testing.T) {
	line := newLine("1")
	expectedLine := newLine("1", "2")

	assert.Equal(t, expectedLine, line.addCell("2"))
}

func TestLinesReverse(t *testing.T) {
	line := newLine("1", "2", "3")
	reversedLine := newLine("3", "2", "1")

	assert.Equal(t, reversedLine, line.reverse())
}

func TestLinesAt(t *testing.T) {
	line := newLine("1")

	assert.Equal(t, "1", line.at(0))
}

func TestLinesAll(t *testing.T) {
	isAOne := func(str string) bool {
		return str == "1"
	}
	sameLine := newLine("1", "1", "1")
	differentLine := newLine("1", "2", "3")

	assert.True(t, sameLine.all(isAOne))
	assert.False(t, differentLine.all(isAOne))
}
