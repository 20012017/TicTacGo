package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type CliGameTest struct {
}

var cliGameTest *CliGameTest = new(CliGameTest)

func TestStartsAGame(t *testing.T) {
	playerOne := cliGameTest.newPlayerStub("X", 0, 1, 6, 5, 8)
	playerTwo := cliGameTest.newPlayerStub("O", 4, 2, 3, 7)
	displaySpy := &DisplaySpy{}
	cliGame := NewCliGame(Game{}, playerOne, playerTwo, displaySpy)

	cliGame.start()

	assert.True(t, displaySpy.welcomeHasBeenCalled)
	assert.True(t, displaySpy.showBoardHasBeenCalled)
	assert.True(t, displaySpy.promptHasBeenCalled)
}

func TestPlaysADrawnGame(t *testing.T) {
	playerOne := cliGameTest.newPlayerStub("X", 0, 1, 6, 5, 8)
	playerTwo := cliGameTest.newPlayerStub("O", 4, 2, 3, 7)
	displaySpy := &DisplaySpy{}

	game := Game{playerOne, playerTwo, NewBoard(9), rules{}}
	cliGame := NewCliGame(game, playerOne, playerTwo, displaySpy)

	cliGame.start()

	assert.True(t, game.isOver())
	assert.True(t, game.isADraw())
	assert.True(t, displaySpy.drawHasBeenCalled)
	assert.True(t, displaySpy.goodbyeHasBeenCalled)
}

func (cliGameTest CliGameTest) newPlayerStub(mark string, moves ...int) PlayerStub {
	return PlayerStub{mark, moves, 0}
}

type PlayerStub struct {
	mark        string
	moves       []int
	currentMove int
}

func (player PlayerStub) Mark() string {
	return player.mark
}

func (player PlayerStub) move(board Board) (int, error) {
	move := player.moves[player.currentMove]
	player.currentMove = player.currentMove + 1
	return move, nil
}

type DisplaySpy struct {
	welcomeHasBeenCalled   bool
	showBoardHasBeenCalled bool
	promptHasBeenCalled    bool
	drawHasBeenCalled      bool
	goodbyeHasBeenCalled   bool
}

func (displaySpy DisplaySpy) write(message string) {
}

func (displaySpy *DisplaySpy) welcome() {
	displaySpy.welcomeHasBeenCalled = true
}

func (displaySpy *DisplaySpy) prompt() {
	displaySpy.promptHasBeenCalled = true
}

func (displaySpy *DisplaySpy) showBoard(board Board) {
	displaySpy.showBoardHasBeenCalled = true
}

func (displaySpy *DisplaySpy) draw() {
	displaySpy.drawHasBeenCalled = true
}

func (displaySpy *DisplaySpy) goodbye() {
	displaySpy.goodbyeHasBeenCalled = true
}
