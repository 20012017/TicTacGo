package ttt

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
