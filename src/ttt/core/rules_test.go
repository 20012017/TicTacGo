package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var rules *Rules = new(Rules)

func TestKnowsIfThereIsAWin(t *testing.T) {
	board := NewMarkedBoard([]string{
		"X", "X", "X",
		"O", "X", "O",
		"O", "X", "O"})

	assert.True(t, rules.isAWin(board, "X", "O"))
}

func TestKnowsIfThereIsANotAWin(t *testing.T) {
	board := NewMarkedBoard([]string{
		"X", "O", "X",
		"O", "X", "O",
		"O", "X", "O"})

	assert.False(t, rules.isAWin(board, "X", "O"))
}

func TestKnowsIfTheGameIsDrawn(t *testing.T) {
	board := NewMarkedBoard([]string{
		"X", "O", "X",
		"O", "X", "O",
		"O", "X", "O"})

	assert.True(t, rules.isADraw(board, "X", "O"))
}

func TestKnowsIfTheGameIsNotDrawn(t *testing.T) {
	board := NewMarkedBoard([]string{
		"X", "X", "X",
		"O", "X", "O",
		"O", "X", "O"})

	assert.False(t, rules.isADraw(board, "X", "O"))
}

func TestKnowsTheWinnerWhenX(t *testing.T) {
	board := NewMarkedBoard([]string{
		"X", "X", "X",
		"O", "X", "O",
		"O", "X", "O"})

	assert.Equal(t, "X", rules.Winner(board, "X", "O"))
}

func TestKnowsTheWinnerWhenO(t *testing.T) {
	board := NewMarkedBoard([]string{
		"O", "O", "O",
		"", "", "",
		"", "", ""})

	assert.Equal(t, "O", rules.Winner(board, "X", "O"))
}

func TestKnowsThereIsNoWinner(t *testing.T) {
	board := NewMarkedBoard([]string{
		"", "", "",
		"", "", "",
		"", "", ""})

	assert.Equal(t, "", rules.Winner(board, "X", "O"))
}

func TestKnowsItIsOverWhenADraw(t *testing.T) {
	board := NewMarkedBoard([]string{
		"X", "O", "X",
		"O", "X", "O",
		"O", "X", "O"})

	assert.True(t, rules.IsOver(board, "X", "O"))
}

func TestKnowsItIsOverWhenAWin(t *testing.T) {
	board := NewMarkedBoard([]string{
		"X", "X", "O",
		"O", "X", "O",
		"O", "X", "O"})

	assert.True(t, rules.IsOver(board, "X", "O"))
}
