package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var numberedSlice []string = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func TestEmptyGrid(t *testing.T) {
	newGrid := NewGrid(9)

	assert.Equal(t, 9, len(newGrid.cells))
}

func TestSplit(t *testing.T) {
	splitGrid := []Line{
		newLine("1", "2", "3"),
		newLine("4", "5", "6"),
		newLine("7", "8", "9")}

	assert.Equal(t, splitGrid, grid().split(3))
}

func TestTranspose(t *testing.T) {
	transposedGrid := []Line{
		newLine("1", "4", "7"),
		newLine("2", "5", "8"),
		newLine("3", "6", "9")}

	assert.Equal(t, transposedGrid, grid().transpose(3))
}

func TestReverseSplit(t *testing.T) {
	reversedSplit := []Line{
		newLine("3", "2", "1"),
		newLine("6", "5", "4"),
		newLine("9", "8", "7")}

	assert.Equal(t, reversedSplit, grid().reverseSplit(3))
}

func TestAnyEmpty(t *testing.T) {
	newGrid := NewGrid(9)

	assert.True(t, newGrid.any(""))
}

func TestFull(t *testing.T) {
	grid := grid()

	assert.False(t, grid.any(""))
}

func grid() Grid {
	return NewPopulatedGrid(numberedSlice)
}
