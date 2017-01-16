package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/cli/display"
	"ttt/cli/players"
	"ttt/core"
)

type GameTest struct{}

var cliGameTest *GameTest = new(GameTest)

func TestStartsAGame(t *testing.T) {
	playerOne := players.NewFake("X", 0, 0, 1, 6, 5, 8)
	playerTwo := players.NewFake("O", 0, 4, 2, 3, 7)
	displaySpy := new(display.Spy)
	game := core.NewGame(playerOne, playerTwo, core.NewBoard(9), new(core.Rules))
	cliGame := NewCliGame(&game, displaySpy)

	cliGame.Start()

	assert.True(t, displaySpy.ClearHasBeenCalled)
	assert.True(t, displaySpy.WelcomeHasBeenCalled)
	assert.True(t, displaySpy.ShowBoardHasBeenCalled)
	assert.True(t, displaySpy.PromptHasBeenCalled)
}

func TestPlaysADrawnGame(t *testing.T) {
	playerOne := players.NewFake("X", 0, 0, 1, 6, 5, 8)
	playerTwo := players.NewFake("O", 0, 4, 2, 3, 7)
	displaySpy := new(display.Spy)
	game := core.NewGame(playerOne, playerTwo, core.NewBoard(9), new(core.Rules))
	cliGame := NewCliGame(&game, displaySpy)

	cliGame.Start()

	assert.True(t, game.IsOver())
	assert.True(t, game.IsADraw())
	assert.True(t, displaySpy.DrawHasBeenCalled)
}

func TestPlaysAWonGame(t *testing.T) {
	playerOne := players.NewFake("X", 0, 0, 1, 2)
	playerTwo := players.NewFake("O", 0, 6, 7)
	displaySpy := &display.Spy{}
	game := core.NewGame(playerOne, playerTwo, core.NewBoard(9), new(core.Rules))
	cliGame := NewCliGame(&game, displaySpy)

	cliGame.Start()

	assert.True(t, game.IsOver())
	assert.True(t, game.IsAWin())
	assert.True(t, displaySpy.WinHasBeenCalled)
}

func TestDisplayErrorsForInvalidMove(t *testing.T) {
	playerOne := players.NewFake("X", 0, -1, 0, 1, 2)
	playerTwo := players.NewFake("O", 0, 6, 7)
	displaySpy := &display.Spy{}
	game := core.NewGame(playerOne, playerTwo, core.NewBoard(9), new(core.Rules))
	cliGame := NewCliGame(&game, displaySpy)

	cliGame.Start()

	assert.True(t, displaySpy.WriteHasBeenCalled)
	assert.Equal(t, "Out of bounds\n", displaySpy.WriteArgument)
}
