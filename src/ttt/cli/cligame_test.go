package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/cli/players"
	"ttt/core"
)

type CliGameTest struct {
}

var cliGameTest *CliGameTest = new(CliGameTest)

func TestStartsAGame(t *testing.T) {
	playerOne := players.NewFake("X", 0, 0, 1, 6, 5, 8)
	playerTwo := players.NewFake("O", 0, 4, 2, 3, 7)
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
	playerOne := players.NewFake("X", 0, 0, 1, 6, 5, 8)
	playerTwo := players.NewFake("O", 0, 4, 2, 3, 7)
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
	playerOne := players.NewFake("X", 0, 0, 1, 2)
	playerTwo := players.NewFake("O", 0, 6, 7)
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
	playerOne := players.NewFake("X", 0, -1, 0, 1, 2)
	playerTwo := players.NewFake("O", 0, 6, 7)
	displaySpy := &DisplaySpy{}
	game := core.NewGame(playerOne, playerTwo, core.NewBoard(9), new(core.Rules))
	cliGame := NewCliGame(game, displaySpy)

	cliGame.Start()

	assert.True(t, displaySpy.writeHasBeenCalled)
	assert.Equal(t, "Out of bounds\n", displaySpy.writeArgument)
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

func (displaySpy *DisplaySpy) Write(message string) {
	displaySpy.writeHasBeenCalled = true
	displaySpy.writeArgument = message
}

func (displaySpy *DisplaySpy) Welcome() {
	displaySpy.welcomeHasBeenCalled = true
}

func (displaySpy *DisplaySpy) Prompt() {
	displaySpy.promptHasBeenCalled = true
}

func (displaySpy *DisplaySpy) ShowBoard(board core.TTTBoard) {
	displaySpy.showBoardHasBeenCalled = true
}

func (displaySpy *DisplaySpy) Draw() {
	displaySpy.drawHasBeenCalled = true
}

func (displaySpy *DisplaySpy) Goodbye() {
	displaySpy.goodbyeHasBeenCalled = true
}

func (displaySpy *DisplaySpy) Win(mark string) {
	displaySpy.winHasBeenCalled = true
}

func (displaySpy *DisplaySpy) Clear() {
	displaySpy.clearHasBeenCalled = true
}
