package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GameTest struct{}

var gameTest GameTest = GameTest{}
var playerOne, playerTwo Player = PlayerDouble{"X"}, PlayerDouble{"O"}
var rule rules = rules{}

func TestCanPlayAMark(t *testing.T) {
	game := gameTest.game(NewBoard(9))

	markedBoard := game.play(0)

	assert.Equal(t, "X", markedBoard.markAt(0))
}

func TestCanPlayTwoMarks(t *testing.T) {
	game := gameTest.game(NewBoard(9))

	markedBoard := game.play(0)
	markedBoard = game.play(1)

	assert.Equal(t, "X", markedBoard.markAt(0))
	assert.Equal(t, "O", markedBoard.markAt(1))
}

func TestKnowsWhenGameIsOverWhenDrawn(t *testing.T) {
	game := gameTest.game(gameTest.fullBoard())

	assert.True(t, game.isOver())
}

func TestKnowsWhenGameIsOverWhenWon(t *testing.T) {
	game := gameTest.game(gameTest.wonBoard())

	assert.True(t, game.isOver())
}

func TestKnowsTheGameWinner(t *testing.T) {
	game := gameTest.game(gameTest.wonBoard())

	isWin, winner := game.result()

	assert.Equal(t, "X", winner)
	assert.True(t, isWin)
}

func TestKnowsWhenThereIsNoWinner(t *testing.T) {
	game := gameTest.game(gameTest.fullBoard())

	isWin, winner := game.result()

	assert.Equal(t, "", winner)
	assert.False(t, isWin)
}

func (gameTest GameTest) game(board Board) Game {
	return NewGame(playerOne, playerTwo, board, rule)
}

func (gameTest GameTest) wonBoard() Board {
	return NewMarkedBoard([]string{
		"X", "O", "X",
		"O", "X", "O",
		"X", "O", ""})
}

func (gameTest GameTest) fullBoard() Board {
	return NewMarkedBoard([]string{
		"X", "O", "X",
		"O", "X", "O",
		"O", "X", "O"})
}

type PlayerDouble struct {
	mark string
}

func (player PlayerDouble) Mark() string {
	return player.mark
}

func (player PlayerDouble) move(board Board) (int, error) {
	return 0, nil
}
