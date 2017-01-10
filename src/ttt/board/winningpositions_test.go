package board

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var numberedBoard []string = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
var winningPositions WinningPositions = NewWinningPositions(
	NewPopulatedGrid(numberedBoard))

func TestReturnsRows(t *testing.T) {
	expectedRows := []Line{
		newLine("1", "2", "3"),
		newLine("4", "5", "6"),
		newLine("7", "8", "9")}
	assert.Equal(t, expectedRows, winningPositions.rows)
}

func TestReturnsColumns(t *testing.T) {
	expectedColumns := []Line{
		newLine("1", "4", "7"),
		newLine("2", "5", "8"),
		newLine("3", "6", "9")}
	assert.Equal(t, expectedColumns, winningPositions.columns)
}

func TestReturnsDiagonals(t *testing.T) {
	expectedDiagonals := []Line{
		newLine("1", "5", "9"),
		newLine("3", "5", "7")}
	assert.Equal(t, expectedDiagonals, winningPositions.diagonals)
}

func TestReturnsAllPossibleWinningPositions(t *testing.T) {
	allWinningPositions := []Line{
		newLine("1", "2", "3"),
		newLine("4", "5", "6"),
		newLine("7", "8", "9"),
		newLine("1", "4", "7"),
		newLine("2", "5", "8"),
		newLine("3", "6", "9"),
		newLine("1", "5", "9"),
		newLine("3", "5", "7")}
	assert.Equal(t, allWinningPositions, winningPositions.all)
}
