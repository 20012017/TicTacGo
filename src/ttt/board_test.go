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