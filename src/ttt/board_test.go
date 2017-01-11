package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyBoardHasNineSpaces(t *testing.T) {
	board := newBoard()

	assert.Equal(t, 9, board.size)
}

func TestCanPlaceAMarkOnBoard(t *testing.T) {
	board := newBoard()

	board = board.placeMark(4, "X")

	assert.Equal(t, "X", board.markAt(4))
}

func TestCanPlaceTwoMarksOnBoard(t *testing.T) {
	board := newBoard()

	board = board.placeMark(4, "X")
	board = board.placeMark(5, "0")

	assert.Equal(t, "X", board.markAt(4))
	assert.Equal(t, "0", board.markAt(5))
}

func TestKnowsTheRowWidth(t *testing.T) {
	board := newBoard()

	assert.Equal(t, 3, board.rowLength())
}

func TestKnowsIfFull(t *testing.T) {
	fullBoard := []string{"X", "O", "X", "O", "X", "O", "O", "X", "O"}

	board := NewMarkedBoard(fullBoard)

	assert.True(t, board.isFull())
}

func TestKnowsAllWinningPositions(t *testing.T) {
	numberedBoard := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	board := NewMarkedBoard(numberedBoard)

	assert.Equal(t, allWinningPositions(), board.winningPositions())
}

func TestCountMarks(t *testing.T) {
	markedBoard := []string{"X", "O", "", "", "", "", "", "", ""}

	board := NewMarkedBoard(markedBoard)

	assert.Equal(t, 2, board.countMarks())
}

func allWinningPositions() []Line {
	return []Line{
		newLine("1", "2", "3"),
		newLine("4", "5", "6"),
		newLine("7", "8", "9"),
		newLine("1", "4", "7"),
		newLine("2", "5", "8"),
		newLine("3", "6", "9"),
		newLine("1", "5", "9"),
		newLine("3", "5", "7")}
}

func newBoard() Board {
	return NewBoard(9)
}
