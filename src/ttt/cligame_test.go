package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type CliGameTest struct {
}

var cliGameTest *CliGameTest = new(CliGameTest)

func TestStartsAGame(t *testing.T) {
	playerOne := cliGameTest.newPlayerFake("X", 0, 1, 6, 5, 8)
	playerTwo := cliGameTest.newPlayerFake("O", 4, 2, 3, 7)
	displaySpy := &DisplaySpy{}
	game := Game{playerOne, playerTwo, NewBoard(9), rules{}}
	cliGame := NewCliGame(game, displaySpy)

	cliGame.start()

	assert.True(t, displaySpy.clearHasBeenCalled)
	assert.True(t, displaySpy.welcomeHasBeenCalled)
	assert.True(t, displaySpy.showBoardHasBeenCalled)
	assert.True(t, displaySpy.promptHasBeenCalled)
}

func TestPlaysADrawnGame(t *testing.T) {
	playerOne := cliGameTest.newPlayerFake("X", 0, 1, 6, 5, 8)
	playerTwo := cliGameTest.newPlayerFake("O", 4, 2, 3, 7)
	displaySpy := &DisplaySpy{}
	game := Game{playerOne, playerTwo, NewBoard(9), rules{}}
	cliGame := NewCliGame(game, displaySpy)

	cliGame.start()

	assert.True(t, game.isOver())
	assert.True(t, game.isADraw())
	assert.True(t, displaySpy.drawHasBeenCalled)
	assert.True(t, displaySpy.goodbyeHasBeenCalled)
}

func TestPlaysAWonGame(t *testing.T) {
	playerOne := cliGameTest.newPlayerFake("X", 0, 1, 2)
	playerTwo := cliGameTest.newPlayerFake("O", 6, 7)
	displaySpy := &DisplaySpy{}
	game := Game{playerOne, playerTwo, NewBoard(9), rules{}}
	cliGame := NewCliGame(game, displaySpy)

	cliGame.start()

	assert.True(t, game.isOver())
	assert.True(t, game.isAWin())
	assert.True(t, displaySpy.winHasBeenCalled)
	assert.True(t, displaySpy.goodbyeHasBeenCalled)
}

func TestDisplayErrorsForInvalidMove(t *testing.T) {
	playerOne := cliGameTest.newPlayerFake("X", -1, 0, 1, 2)
	playerTwo := cliGameTest.newPlayerFake("O", 6, 7)
	displaySpy := &DisplaySpy{}
	game := Game{playerOne, playerTwo, NewBoard(9), rules{}}
	cliGame := NewCliGame(game, displaySpy)

	cliGame.start()

	assert.True(t, displaySpy.writeHasBeenCalled)
	assert.Equal(t, "Out of bounds\n", displaySpy.writeArgument)
}

func (cliGameTest CliGameTest) newPlayerFake(mark string, moves ...int) *PlayerFake {
	return &PlayerFake{mark, 0, moves}
}

type DisplaySpy struct {
	welcomeHasBeenCalled   bool
	showBoardHasBeenCalled bool
	promptHasBeenCalled    bool
	drawHasBeenCalled      bool
	goodbyeHasBeenCalled   bool
	winHasBeenCalled       bool
	clearHasBeenCalled     bool
	writeHasBeenCalled     bool
	writeArgument          string
}

func (displaySpy *DisplaySpy) write(message string) {
	displaySpy.writeHasBeenCalled = true
	displaySpy.writeArgument = message
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

func (displaySpy *DisplaySpy) win(mark string) {
	displaySpy.winHasBeenCalled = true
}

func (displaySpy *DisplaySpy) clear() {
	displaySpy.clearHasBeenCalled = true
}
