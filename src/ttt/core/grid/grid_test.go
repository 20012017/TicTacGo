package grid

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core/marks"
)

var numberedSlice []marks.Mark = []marks.Mark{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
var grid Grid = NewPopulatedGrid(numberedSlice)

func TestEmptyGrid(t *testing.T) {
	newGrid := NewGrid(9)

	assert.Equal(t, 9, len(newGrid.cells))
}

func TestSplit(t *testing.T) {
	expectedSplitGrid := []Line{
		NewLine("1", "2", "3"),
		NewLine("4", "5", "6"),
		NewLine("7", "8", "9")}

	splitGrid := grid.split(3)

	assert.Equal(t, expectedSplitGrid, splitGrid)
}

func TestTranspose(t *testing.T) {
	expectedTransposedGrid := []Line{
		NewLine("1", "4", "7"),
		NewLine("2", "5", "8"),
		NewLine("3", "6", "9")}

	transposedGrid := grid.transpose(3)

	assert.Equal(t, expectedTransposedGrid, transposedGrid)
}

func TestReverseSplit(t *testing.T) {
	expectedReverseSplit := []Line{
		NewLine("3", "2", "1"),
		NewLine("6", "5", "4"),
		NewLine("9", "8", "7")}

	reverseSplit := grid.reverseSplit(3)

	assert.Equal(t, expectedReverseSplit, reverseSplit)
}

func TestAnyEmpty(t *testing.T) {
	newGrid := NewGrid(9)

	assert.True(t, newGrid.Any(""))
}

func TestFull(t *testing.T) {
	assert.False(t, grid.Any(""))
}

func TestMark(t *testing.T) {
	grid := NewGrid(9)

	markedGrid := grid.Mark(0, marks.X)

	assert.Equal(t,
		[]marks.Mark{"", "", "", "", "", "", "", "", ""},
		grid.Cells())
	assert.Equal(t,
		[]marks.Mark{marks.X, "", "", "", "", "", "", "", ""},
		markedGrid.Cells())
}
