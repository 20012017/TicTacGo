package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var playerOne, playerTwo Player = PlayerDouble{"X"}, PlayerDouble{"O"}
var rule rules = rules{}

func TestCanPlayAMark(t *testing.T) {
	game := game(NewBoard(9))

	markedBoard := game.play(0)

	assert.Equal(t, "X", markedBoard.markAt(0))
}

func TestCanPlayTwoMarks(t *testing.T) {
	game := game(NewBoard(9))

	markedBoard := game.play(0)
	markedBoard = game.play(1)

	assert.Equal(t, "X", markedBoard.markAt(0))
	assert.Equal(t, "O", markedBoard.markAt(1))
}

func TestKnowsWhenGameIsOverWhenDrawn(t *testing.T) {
	game := game(fullBoard())

	assert.True(t, game.isOver())
}

func TestKnowsWhenGameIsOverWhenWon(t *testing.T) {
	game := game(wonBoard())

	assert.True(t, game.isOver())
}

func TestKnowsTheGameWinner(t *testing.T) {
	game := game(wonBoard())

	isWin, winner := game.result()

	assert.Equal(t, "X", winner)
	assert.True(t, isWin)
}

func TestKnowsWhenThereIsNoWinner(t *testing.T) {
	game := game(fullBoard())

	isWin, winner := game.result()

	assert.Equal(t, "", winner)
	assert.False(t, isWin)
}

func game(board Board) Game {
	return NewGame(playerOne, playerTwo, board, rule)
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
