package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core/marks"
)

var rules *Rules = new(Rules)
var X, O, EMPTY marks.Mark = marks.X, marks.O, marks.EMPTY

func TestKnowsIfThereIsAWin(t *testing.T) {
	board := NewMarkedBoard([]marks.Mark{
		X, X, X,
		O, X, O,
		O, X, O})

	isAWin := rules.isAWin(board, X, O)

	assert.True(t, isAWin)
}

func TestKnowsIfThereIsANotAWin(t *testing.T) {
	board := NewMarkedBoard([]marks.Mark{
		X, O, X,
		O, X, O,
		O, X, O})

	isAWin := rules.isAWin(board, X, O)

	assert.False(t, isAWin)
}

func TestKnowsIfTheGameIsDrawn(t *testing.T) {
	board := NewMarkedBoard([]marks.Mark{
		X, O, X,
		O, X, O,
		O, X, O})

	isADraw := rules.isADraw(board, X, O)

	assert.True(t, isADraw)
}

func TestKnowsIfTheGameIsNotDrawn(t *testing.T) {
	board := NewMarkedBoard([]marks.Mark{
		X, X, X,
		O, X, O,
		O, X, O})

	isADraw := rules.isADraw(board, X, O)

	assert.False(t, isADraw)
}

func TestKnowsTheWinnerWhenX(t *testing.T) {
	board := NewMarkedBoard([]marks.Mark{
		X, X, X,
		O, X, O,
		O, X, O})

	winner := rules.Winner(board, X, O)

	assert.Equal(t, X, winner)
}

func TestKnowsTheWinnerWhenO(t *testing.T) {
	board := NewMarkedBoard([]marks.Mark{
		O, O, O,
		EMPTY, EMPTY, EMPTY,
		EMPTY, EMPTY, EMPTY})

	winner := rules.Winner(board, X, O)

	assert.Equal(t, O, winner)
}

func TestKnowsThereIsNoWinner(t *testing.T) {
	board := NewMarkedBoard([]marks.Mark{
		EMPTY, EMPTY, EMPTY,
		EMPTY, EMPTY, EMPTY,
		EMPTY, EMPTY, EMPTY})

	winner := rules.Winner(board, X, O)

	assert.Equal(t, EMPTY, winner)
}

func TestKnowsItIsOverWhenADraw(t *testing.T) {
	board := NewMarkedBoard([]marks.Mark{
		X, O, X,
		O, X, O,
		O, X, O})

	isOver := rules.IsOver(board, X, O)

	assert.True(t, isOver)
}

func TestKnowsItIsOverWhenAWin(t *testing.T) {
	board := NewMarkedBoard([]marks.Mark{
		X, X, O,
		O, X, O,
		O, X, O})

	isOver := rules.IsOver(board, X, O)

	assert.True(t, isOver)
}
