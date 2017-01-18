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

	assert.True(t, rules.isAWin(board, X, O))
}

func TestKnowsIfThereIsANotAWin(t *testing.T) {
	board := NewMarkedBoard([]marks.Mark{
		X, O, X,
		O, X, O,
		O, X, O})

	assert.False(t, rules.isAWin(board, X, O))
}

func TestKnowsIfTheGameIsDrawn(t *testing.T) {
	board := NewMarkedBoard([]marks.Mark{
		X, O, X,
		O, X, O,
		O, X, O})

	assert.True(t, rules.isADraw(board, X, O))
}

func TestKnowsIfTheGameIsNotDrawn(t *testing.T) {
	board := NewMarkedBoard([]marks.Mark{
		X, X, X,
		O, X, O,
		O, X, O})

	assert.False(t, rules.isADraw(board, X, O))
}

func TestKnowsTheWinnerWhenX(t *testing.T) {
	board := NewMarkedBoard([]marks.Mark{
		X, X, X,
		O, X, O,
		O, X, O})

	assert.Equal(t, X, rules.Winner(board, X, O))
}

func TestKnowsTheWinnerWhenO(t *testing.T) {
	board := NewMarkedBoard([]marks.Mark{
		O, O, O,
		EMPTY, EMPTY, EMPTY,
		EMPTY, EMPTY, EMPTY})

	assert.Equal(t, O, rules.Winner(board, X, O))
}

func TestKnowsThereIsNoWinner(t *testing.T) {
	board := NewMarkedBoard([]marks.Mark{
		EMPTY, EMPTY, EMPTY,
		EMPTY, EMPTY, EMPTY,
		EMPTY, EMPTY, EMPTY})

	assert.Equal(t, EMPTY, rules.Winner(board, X, O))
}

func TestKnowsItIsOverWhenADraw(t *testing.T) {
	board := NewMarkedBoard([]marks.Mark{
		X, O, X,
		O, X, O,
		O, X, O})

	assert.True(t, rules.IsOver(board, X, O))
}

func TestKnowsItIsOverWhenAWin(t *testing.T) {
	board := NewMarkedBoard([]marks.Mark{
		X, X, O,
		O, X, O,
		O, X, O})

	assert.True(t, rules.IsOver(board, X, O))
}
