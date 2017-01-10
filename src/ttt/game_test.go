package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var playerOne, playerTwo Player = PlayerDouble{"X"}, PlayerDouble{"O"}
var rule rules = rules{}

func TestCanPlayAMark(t *testing.T) {
	var board Board = NewBoard(9)
	game := NewGame(playerOne, playerTwo, board, rule)

	markedBoard := game.play(0)

	assert.Equal(t, "X", markedBoard.markAt(0))
}

func TestCanPlayTwoMarks(t *testing.T) {
	var board Board = NewBoard(9)
	game := NewGame(playerOne, playerTwo, board, rule)

	markedBoard := game.play(0)
	markedBoard = game.play(1)

	assert.Equal(t, "X", markedBoard.markAt(0))
	assert.Equal(t, "O", markedBoard.markAt(1))
}

func TestKnowsWhenGameIsOverWhenDrawn(t *testing.T) {
	game := NewGame(playerOne, playerTwo, fullBoard(), rule)

	assert.True(t, game.isOver())
}

func TestKnowsWhenGameIsOverWhenWon(t *testing.T) {
	game := NewGame(playerOne, playerTwo, wonBoard(), rule)

	assert.True(t, game.isOver())
}

func TestKnowsTheGameWinner(t *testing.T) {
	game := NewGame(playerOne, playerTwo, wonBoard(), rule)

	isWin, winner := game.result()
	assert.Equal(t, "X", winner)
	assert.True(t, isWin)
}

func wonBoard() Board {
	return NewMarkedBoard([]string{
		"X", "O", "X",
		"O", "X", "O",
		"X", "O", ""})
}

func fullBoard() Board {
	return NewMarkedBoard([]string{
		"X", "O", "X",
		"O", "X", "O",
		"O", "X", "O"})
}

type PlayerDouble struct {
	playerMark string
}

func (player PlayerDouble) mark() string {
	return player.playerMark
}
