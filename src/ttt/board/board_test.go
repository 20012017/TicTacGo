package board

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyBoardHasNineSpaces(t *testing.T) {
	assert.Equal(t, 9, NewBoard(9).size)
}

func TestCanPlaceAMarkOnBoard(t *testing.T) {
	board := NewBoard(9).placeMark(4, "X")
	assert.Equal(t, "X", board.markAt(4))
}

func TestCanPlaceTwoMarksOnBoard(t *testing.T) {
	board := NewBoard(9).placeMark(4, "X")
	board = board.placeMark(5, "0")
	assert.Equal(t, "X", board.markAt(4))
	assert.Equal(t, "0", board.markAt(5))
}

func TestKnowsTheRowWidth(t *testing.T) {
	board := NewBoard(9)
	assert.Equal(t, 3, board.rowLength())
}

func TestKnowsIfFull(t *testing.T) {
	fullBoard := []string{"X", "O", "X", "O", "X", "O", "O", "X", "O"}
	board := NewMarkedBoard(fullBoard)
	assert.True(t, board.IsFull())
}

func TestKnowsAllWinningPositions(t *testing.T) {
	numberedBoard := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	board := NewMarkedBoard(numberedBoard)
	allWinningPositions := []Line{
		newLine("1", "2", "3"),
		newLine("4", "5", "6"),
		newLine("7", "8", "9"),
		newLine("1", "4", "7"),
		newLine("2", "5", "8"),
		newLine("3", "6", "9"),
		newLine("1", "5", "9"),
		newLine("3", "5", "7")}
	assert.Equal(t, allWinningPositions, board.WinningPositions())
}
