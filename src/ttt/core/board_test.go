package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core/grid"
	"ttt/core/marks"
)

type boardTests struct{}

var boardTest boardTests = boardTests{}

func TestEmptyBoardHasNineSpaces(t *testing.T) {
	board := boardTest.newBoard()

	assert.Equal(t, 9, board.size)
}

func TestCanPlaceAMarkOnBoard(t *testing.T) {
	board := boardTest.newBoard()

	board = board.PlaceMark(4, X)

	assert.Equal(t, X, board.MarkAt(4))
}

func TestCanPlaceTwoMarksOnBoard(t *testing.T) {
	board := boardTest.newBoard()

	board = board.PlaceMark(4, X)
	board = board.PlaceMark(5, O)

	assert.Equal(t, X, board.MarkAt(4))
	assert.Equal(t, O, board.MarkAt(5))
}

func TestBoardIsImmutable(t *testing.T) {
	board := boardTest.newBoard()

	board = board.PlaceMark(4, X)
	board.PlaceMark(5, O)

	assert.Equal(t, X, board.MarkAt(4))
	assert.Equal(t, EMPTY, board.MarkAt(5))
}

func TestKnowsTheRowWidth(t *testing.T) {
	board := boardTest.newBoard()

	assert.Equal(t, 3, board.rowLength())
}

func TestKnowsIfFull(t *testing.T) {
	fullBoard := []marks.Mark{X, O, X, O, X, O, O, X, O}

	board := NewMarkedBoard(fullBoard)

	assert.True(t, board.isFull())
}

func TestKnowsAllWinningPositions(t *testing.T) {
	numberedBoard := []marks.Mark{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	board := NewMarkedBoard(numberedBoard)

	assert.Equal(t, boardTest.allWinningPositions(), board.winningPositions())
}

func TestCountMarks(t *testing.T) {
	markedBoard := []marks.Mark{X, O, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY}

	board := NewMarkedBoard(markedBoard)

	assert.Equal(t, 2, board.CountMarks())
}

func TestAvailableMovesOnEmptyBoard(t *testing.T) {
	board := boardTest.newBoard()

	availablePositions := board.AvailableMoves()

	assert.Equal(t, 9, len(availablePositions))
}

func TestAvailableMovesOnAMarkedBoard(t *testing.T) {
	markedBoard := []marks.Mark{X, O, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY}
	board := NewMarkedBoard(markedBoard)

	availablePositions := board.AvailableMoves()

	assert.Equal(t, 7, len(availablePositions))
}

func TestAvailableMovesDecreaseAfterMark(t *testing.T) {
	markedBoard := []marks.Mark{X, O, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY}
	board := NewMarkedBoard(markedBoard)

	availablePositions := board.AvailableMoves()
	board = board.PlaceMark(8, X)
	newAvailablePositions := board.AvailableMoves()

	assert.Equal(t, 7, len(availablePositions))
	assert.Equal(t, 6, len(newAvailablePositions))
}

func (boardTest boardTests) allWinningPositions() []grid.Line {
	return []grid.Line{
		boardTest.newLine("1", "2", "3"),
		boardTest.newLine("4", "5", "6"),
		boardTest.newLine("7", "8", "9"),
		boardTest.newLine("1", "4", "7"),
		boardTest.newLine("2", "5", "8"),
		boardTest.newLine("3", "6", "9"),
		boardTest.newLine("1", "5", "9"),
		boardTest.newLine("3", "5", "7")}
}

func (boardTest boardTests) newLine(cell1, cell2, cell3 marks.Mark) grid.Line {
	return grid.NewLine(cell1, cell2, cell3)
}

func (boardTest boardTests) newBoard() Board {
	return NewBoard(9)
}
