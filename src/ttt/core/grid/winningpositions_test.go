package grid

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core/marks"
)

var numberedBoard []marks.Mark = []marks.Mark{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
var winningPositions WinningPositions = NewWinningPositions(
	NewPopulatedGrid(numberedBoard), 3)

func TestReturnsRows(t *testing.T) {
	expectedRows := []Line{
		NewLine("1", "2", "3"),
		NewLine("4", "5", "6"),
		NewLine("7", "8", "9")}

	assert.Equal(t, expectedRows, winningPositions.rows)
}

func TestReturnsColumns(t *testing.T) {
	expectedColumns := []Line{
		NewLine("1", "4", "7"),
		NewLine("2", "5", "8"),
		NewLine("3", "6", "9")}

	assert.Equal(t, expectedColumns, winningPositions.columns)
}

func TestReturnsDiagonals(t *testing.T) {
	expectedDiagonals := []Line{
		NewLine("1", "5", "9"),
		NewLine("3", "5", "7")}

	assert.Equal(t, expectedDiagonals, winningPositions.diagonals)
}

func TestReturnsAllPossibleWinningPositions(t *testing.T) {
	allWinningPositions := []Line{
		NewLine("1", "2", "3"),
		NewLine("4", "5", "6"),
		NewLine("7", "8", "9"),
		NewLine("1", "4", "7"),
		NewLine("2", "5", "8"),
		NewLine("3", "6", "9"),
		NewLine("1", "5", "9"),
		NewLine("3", "5", "7")}

	assert.Equal(t, allWinningPositions, winningPositions.All)
}
