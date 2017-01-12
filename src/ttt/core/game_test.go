package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GameTest struct{}

var gameTest GameTest = GameTest{}
var playerOne, playerTwo Player = PlayerDouble{"X"}, PlayerDouble{"O"}
var rule *Rules = new(Rules)

func TestCanPlayAMark(t *testing.T) {
	game := gameTest.game(NewBoard(9))

	markedBoard := game.Play(0)

	assert.Equal(t, "X", markedBoard.MarkAt(0))
}

func TestCanPlayTwoMarks(t *testing.T) {
	game := gameTest.game(NewBoard(9))

	markedBoard := game.Play(0)
	markedBoard = game.Play(1)

	assert.Equal(t, "X", markedBoard.MarkAt(0))
	assert.Equal(t, "O", markedBoard.MarkAt(1))
}

func TestKnowsWhenGameIsOverWhenDrawn(t *testing.T) {
	game := gameTest.game(gameTest.fullBoard())

	assert.True(t, game.IsOver())
}

func TestKnowsWhenGameIsOverWhenWon(t *testing.T) {
	game := gameTest.game(gameTest.wonBoard())

	assert.True(t, game.IsOver())
}

func TestKnowsTheGameWinner(t *testing.T) {
	game := gameTest.game(gameTest.wonBoard())

	isWin, winner := game.Result()

	assert.Equal(t, "X", winner)
	assert.True(t, isWin)
}

func TestKnowsWhenThereIsNoWinner(t *testing.T) {
	game := gameTest.game(gameTest.fullBoard())

	isWin, winner := game.Result()

	assert.Equal(t, "", winner)
	assert.False(t, isWin)
}

func TestKnowsTheCurrentPlayer(t *testing.T) {
	game := gameTest.game(NewBoard(9))

	player := game.CurrentPlayer()

	assert.Equal(t, "X", player.Mark())
}

func (gameTest GameTest) game(tttboard TTTBoard) Game {
	return NewGame(playerOne, playerTwo, tttboard, rule)
}

func (gameTest GameTest) wonBoard() TTTBoard {
	return NewMarkedBoard([]string{
		"X", "O", "X",
		"O", "X", "O",
		"X", "O", ""})
}

func (gameTest GameTest) fullBoard() TTTBoard {
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

func (player PlayerDouble) Move(board TTTBoard) (int, error) {
	return 0, nil
}
