package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/board"
)

func TestKnowsIfThereIsAWin(t *testing.T) {
	rules := rules{}
	board := board.NewMarkedBoard([]string{
		"X", "X", "X",
		"O", "X", "O",
		"O", "X", "O"})
	assert.True(t, rules.isAWin(board, "X", "O"))
}

func TestKnowsIfThereIsANotAWin(t *testing.T) {
	rules := rules{}
	board := board.NewMarkedBoard([]string{
		"X", "O", "X",
		"O", "X", "O",
		"O", "X", "O"})
	assert.False(t, rules.isAWin(board, "X", "O"))
}

func TestKnowsIfTheGameIsDrawn(t *testing.T) {
	rules := rules{}
	board := board.NewMarkedBoard([]string{
		"X", "O", "X",
		"O", "X", "O",
		"O", "X", "O"})
	assert.True(t, rules.isADraw(board, "X", "O"))
}

func TestKnowsIfTheGameIsNotDrawn(t *testing.T) {
	rules := rules{}
	board := board.NewMarkedBoard([]string{
		"X", "X", "X",
		"O", "X", "O",
		"O", "X", "O"})
	assert.False(t, rules.isADraw(board, "X", "O"))
}
