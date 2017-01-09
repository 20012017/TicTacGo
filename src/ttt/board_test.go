package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyBoardHasNineSpaces(t *testing.T) {
	assert.Equal(t, 9, NewBoard().Size())
}

func TestCanPlaceAMarkOnBoard(t *testing.T) {
	board := NewBoard().PlaceMark(4, "X")
	assert.Equal(t, "X", board.MarkAt(4))
}

func TestCanPlaceTwoMarksOnBoard(t *testing.T) {
	board := NewBoard().PlaceMark(4, "X")
	board = board.PlaceMark(5, "0")
	assert.Equal(t, "X", board.MarkAt(4))
	assert.Equal(t, "0", board.MarkAt(5))
}

var numberedBoard Board = newMarkedBoard([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"})

func TestCanReturnRows(t *testing.T) {
	rows := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	assert.Equal(t, rows, numberedBoard.Rows())
}

func TestCanReturnColumns(t *testing.T) {
	columns := [][]string{{"1", "4", "7"}, {"2", "5", "8"}, {"3", "6", "9"}}
	assert.Equal(t, columns, numberedBoard.Columns())
}

func TestCanReturnDiagonals(t *testing.T) {
	diagonals := [][]string{{"1", "5", "9"}, {"3", "5", "7"}}
	assert.Equal(t, diagonals, numberedBoard.Diagonals())
}
