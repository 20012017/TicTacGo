package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKnowsIfThereIsAWin(t *testing.T) {
	rules := rules{}
	board := NewMarkedBoard([]string{
		"X", "X", "X",
		"O", "X", "O",
		"O", "X", "O"})
	assert.True(t, rules.isAWin(board))
}

func TestKnowsIfThereIsANotAWin(t *testing.T) {
	rules := rules{}
	board := NewMarkedBoard([]string{
		"X", "O", "X",
		"O", "X", "O",
		"O", "X", "O"})
	assert.False(t, rules.isAWin(board))
}

func TestKnowsIfTheGameIsDrawn(t *testing.T) {
	rules := rules{}
	board := NewMarkedBoard([]string{
		"X", "O", "X",
		"O", "X", "O",
		"O", "X", "O"})
	assert.True(t, rules.isADraw(board))
}

func TestKnowsIfTheGameIsNotDrawn(t *testing.T) {
	rules := rules{}
	board := NewMarkedBoard([]string{
		"X", "X", "X",
		"O", "X", "O",
		"O", "X", "O"})
	assert.False(t, rules.isADraw(board))
}
