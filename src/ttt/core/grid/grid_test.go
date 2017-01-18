package grid

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core/marks"
)

type gridTests struct{}

var gridTest gridTests = gridTests{}
var numberedSlice []marks.Mark = []marks.Mark{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func TestEmptyGrid(t *testing.T) {
	newGrid := NewGrid(9)

	assert.Equal(t, 9, len(newGrid.cells))
}

func TestSplit(t *testing.T) {
	splitGrid := []Line{
		NewLine("1", "2", "3"),
		NewLine("4", "5", "6"),
		NewLine("7", "8", "9")}

	assert.Equal(t, splitGrid, gridTest.grid().split(3))
}

func TestTranspose(t *testing.T) {
	transposedGrid := []Line{
		NewLine("1", "4", "7"),
		NewLine("2", "5", "8"),
		NewLine("3", "6", "9")}

	assert.Equal(t, transposedGrid, gridTest.grid().transpose(3))
}

func TestReverseSplit(t *testing.T) {
	reversedSplit := []Line{
		NewLine("3", "2", "1"),
		NewLine("6", "5", "4"),
		NewLine("9", "8", "7")}

	assert.Equal(t, reversedSplit, gridTest.grid().reverseSplit(3))
}

func TestAnyEmpty(t *testing.T) {
	newGrid := NewGrid(9)

	assert.True(t, newGrid.Any(""))
}

func TestFull(t *testing.T) {
	grid := gridTest.grid()

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

func (gridTest gridTests) grid() Grid {
	return NewPopulatedGrid(numberedSlice)
}
