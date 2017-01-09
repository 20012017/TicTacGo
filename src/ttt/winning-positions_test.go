package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var numberedBoard []string = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
var winningPositions WinningPositions = NewWinningPositions(numberedBoard, NewPopulatedGrid(numberedBoard))

func TestReturnsRows(t *testing.T) {
	expectedRows := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	assert.Equal(t, expectedRows, winningPositions.rows)
}

func TestReturnsColumns(t *testing.T) {
	expectedColumns := [][]string{{"1", "4", "7"}, {"2", "5", "8"}, {"3", "6", "9"}}
	assert.Equal(t, expectedColumns, winningPositions.columns)
}

func TestReturnsDiagonals(t *testing.T) {
	diagonals := [][]string{{"1", "5", "9"}, {"3", "5", "7"}}
	assert.Equal(t, diagonals, winningPositions.diagonals)
}

func TestReturnsAllPossibleWinningPositions(t *testing.T) {
	allWinningPositions := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"},
		{"1", "4", "7"}, {"2", "5", "8"}, {"3", "6", "9"},
		{"1", "5", "9"}, {"3", "5", "7"}}
	assert.Equal(t, allWinningPositions, winningPositions.All())
}
