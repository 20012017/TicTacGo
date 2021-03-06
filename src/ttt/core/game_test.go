package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core/marks"
)

type gameTests struct{}

var gameTest gameTests = gameTests{}

var playerOne, playerTwo Player = playerStub{X}, playerStub{O}
var rule *Rules = new(Rules)

func TestCanPlayAMark(t *testing.T) {
	game := gameTest.game(NewBoard(9))

	markedBoard := game.Play(0)

	assert.Equal(t, X, markedBoard.MarkAt(0))
}

func TestCanPlayTwoMarks(t *testing.T) {
	game := gameTest.game(NewBoard(9))

	game.Play(0)
	game.Play(1)

	assert.Equal(t, X, game.board.MarkAt(0))
	assert.Equal(t, O, game.board.MarkAt(1))
}

func TestKnowsWhenGameIsOverWhenDrawn(t *testing.T) {
	game := gameTest.game(gameTest.fullBoard())

	assert.True(t, game.IsOver())
}

func TestCanDrawAGame(t *testing.T) {
	game := gameTest.game(NewBoard(9))

	game.Play(0)
	game.Play(3)
	game.Play(1)
	game.Play(4)
	game.Play(5)
	game.Play(2)
	game.Play(6)
	game.Play(7)
	board := game.Play(8)

	assert.True(t, board.isFull())
	assert.True(t, game.IsOver())
	assert.True(t, game.IsADraw())
}

func TestKnowsWhenGameIsOverWhenWon(t *testing.T) {
	game := gameTest.game(gameTest.wonBoard())

	assert.True(t, game.IsOver())
}

func TestKnowsTheGameWinner(t *testing.T) {
	game := gameTest.game(gameTest.wonBoard())

	isWin, winner := game.Result()

	assert.Equal(t, X, winner)
	assert.True(t, isWin)
}

func TestKnowsWhenThereIsNoWinner(t *testing.T) {
	game := gameTest.game(gameTest.fullBoard())

	isWin, winner := game.Result()

	assert.Equal(t, EMPTY, winner)
	assert.False(t, isWin)
}

func TestKnowsTheCurrentPlayer(t *testing.T) {
	game := gameTest.game(NewBoard(9))

	player := game.CurrentPlayer()

	assert.Equal(t, X, player.Mark())
}

func TestSwitchesThePlayer(t *testing.T) {
	game := gameTest.game(NewBoard(9))

	game.Play(0)
	player := game.CurrentPlayer()
	game.Play(1)

	assert.Equal(t, O, player.Mark())
	assert.Equal(t, X, game.board.MarkAt(0))
	assert.Equal(t, O, game.board.MarkAt(1))
}

func (gameTest gameTests) game(tttboard Board) Game {
	return NewGame(playerOne, playerTwo, tttboard, rule)
}

func (gameTest gameTests) wonBoard() Board {
	return NewMarkedBoard([]marks.Mark{
		X, O, X,
		O, X, O,
		X, O, EMPTY})
}

func (gameTest gameTests) fullBoard() Board {
	return NewMarkedBoard([]marks.Mark{
		X, O, X,
		O, X, O,
		O, X, O})
}

type playerStub struct {
	mark marks.Mark
}

func (player playerStub) Mark() marks.Mark {
	return player.mark
}

func (player playerStub) Move(board Board) (int, error) {
	return 0, nil
}
