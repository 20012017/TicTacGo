package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core/board"
)

var boardTest BoardTest = BoardTest{}

type BoardTest struct{}

func TestEmptyBoardHasNineSpaces(t *testing.T) {
	board := boardTest.newBoard()

	assert.Equal(t, 9, board.size)
}

func TestCanPlaceAMarkOnBoard(t *testing.T) {
	board := boardTest.newBoard()

	board = board.PlaceMark(4, "X")

	assert.Equal(t, "X", board.MarkAt(4))
}

func TestCanPlaceTwoMarksOnBoard(t *testing.T) {
	board := boardTest.newBoard()

	board = board.PlaceMark(4, "X")
	board = board.PlaceMark(5, "0")

	assert.Equal(t, "X", board.MarkAt(4))
	assert.Equal(t, "0", board.MarkAt(5))
}

func TestBoardIsImmutable(t *testing.T) {
	board := boardTest.newBoard()

	board = board.PlaceMark(4, "X")
	board.PlaceMark(5, "0")

	assert.Equal(t, "X", board.MarkAt(4))
	assert.Equal(t, "", board.MarkAt(5))
}

func TestKnowsTheRowWidth(t *testing.T) {
	board := boardTest.newBoard()

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

	assert.Equal(t, boardTest.allWinningPositions(), board.winningPositions())
}

func TestCountMarks(t *testing.T) {
	markedBoard := []string{"X", "O", "", "", "", "", "", "", ""}

	board := NewMarkedBoard(markedBoard)

	assert.Equal(t, 2, board.CountMarks())
}

func TestAvailableMovesOnEmptyBoard(t *testing.T) {
	board := boardTest.newBoard()

	availablePositions := board.AvailableMoves()

	assert.Equal(t, 9, len(availablePositions))
}

func TestAvailableMovesOnAMarkedBoard(t *testing.T) {
	markedBoard := []string{"X", "O", "", "", "", "", "", "", ""}
	board := NewMarkedBoard(markedBoard)

	availablePositions := board.AvailableMoves()

	assert.Equal(t, 7, len(availablePositions))
}

func (boardTest BoardTest) allWinningPositions() []board.Line {
	return []board.Line{
		board.NewLine("1", "2", "3"),
		board.NewLine("4", "5", "6"),
		board.NewLine("7", "8", "9"),
		board.NewLine("1", "4", "7"),
		board.NewLine("2", "5", "8"),
		board.NewLine("3", "6", "9"),
		board.NewLine("1", "5", "9"),
		board.NewLine("3", "5", "7")}
}

func (boardTest BoardTest) newBoard() TTTBoard {
	return NewBoard(9)
}
