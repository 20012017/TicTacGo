package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core"
)

type CliGameTest struct {
}

var cliGameTest *CliGameTest = new(CliGameTest)

func TestStartsAGame(t *testing.T) {
	playerOne := cliGameTest.newPlayerFake("X", 0, 1, 6, 5, 8)
	playerTwo := cliGameTest.newPlayerFake("O", 4, 2, 3, 7)
	displaySpy := &DisplaySpy{}
	game := core.NewGame(playerOne, playerTwo, core.NewBoard(9), new(core.Rules))
	cliGame := NewCliGame(game, displaySpy)

	cliGame.Start()

	assert.True(t, displaySpy.clearHasBeenCalled)
	assert.True(t, displaySpy.welcomeHasBeenCalled)
	assert.True(t, displaySpy.showBoardHasBeenCalled)
	assert.True(t, displaySpy.promptHasBeenCalled)
}

func TestPlaysADrawnGame(t *testing.T) {
	playerOne := cliGameTest.newPlayerFake("X", 0, 1, 6, 5, 8)
	playerTwo := cliGameTest.newPlayerFake("O", 4, 2, 3, 7)
	displaySpy := &DisplaySpy{}
	game := core.NewGame(playerOne, playerTwo, core.NewBoard(9), new(core.Rules))
	cliGame := NewCliGame(game, displaySpy)

	cliGame.Start()

	assert.True(t, game.IsOver())
	assert.True(t, game.IsADraw())
	assert.True(t, displaySpy.drawHasBeenCalled)
	assert.True(t, displaySpy.goodbyeHasBeenCalled)
}

func TestPlaysAWonGame(t *testing.T) {
	playerOne := cliGameTest.newPlayerFake("X", 0, 1, 2)
	playerTwo := cliGameTest.newPlayerFake("O", 6, 7)
	displaySpy := &DisplaySpy{}
	game := core.NewGame(playerOne, playerTwo, core.NewBoard(9), new(core.Rules))
	cliGame := NewCliGame(game, displaySpy)

	cliGame.Start()

	assert.True(t, game.IsOver())
	assert.True(t, game.IsAWin())
	assert.True(t, displaySpy.winHasBeenCalled)
	assert.True(t, displaySpy.goodbyeHasBeenCalled)
}

func TestDisplayErrorsForInvalidMove(t *testing.T) {
	playerOne := cliGameTest.newPlayerFake("X", -1, 0, 1, 2)
	playerTwo := cliGameTest.newPlayerFake("O", 6, 7)
	displaySpy := &DisplaySpy{}
	game := core.NewGame(playerOne, playerTwo, core.NewBoard(9), new(core.Rules))
	cliGame := NewCliGame(game, displaySpy)

	cliGame.Start()

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

func (displaySpy *DisplaySpy) showBoard(board core.TTTBoard) {
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
