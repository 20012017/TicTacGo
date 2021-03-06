package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/cli/display"
	"ttt/cli/players"
	"ttt/core"
	"ttt/core/marks"
)

type gameTests struct{}

var cliGameTest *gameTests = new(gameTests)

var playerPrompter Prompter = NewPrompter(GameOptions)

func TestStartsAGame(t *testing.T) {
	playerOne := players.NewFake(marks.X, 0, 0, 1, 6, 5, 8)
	playerTwo := players.NewFake(marks.O, 0, 4, 2, 3, 7)
	displaySpy := new(display.Spy)
	game := cliGameTest.createGame(playerOne, playerTwo)
	cliGame := NewCliGame(&game, displaySpy, playerPrompter, 1)

	cliGame.Start()

	assert.True(t, displaySpy.ClearHasBeenCalled)
	assert.True(t, displaySpy.WelcomeHasBeenCalled)
	assert.True(t, displaySpy.ShowBoardHasBeenCalled)
	assert.True(t, displaySpy.HumanPromptHasBeenCalled)
}

func TestPlaysADrawnGame(t *testing.T) {
	playerOne := players.NewFake(marks.X, 0, 0, 1, 6, 5, 8)
	playerTwo := players.NewFake(marks.O, 0, 4, 2, 3, 7)
	displaySpy := new(display.Spy)
	game := cliGameTest.createGame(playerOne, playerTwo)
	cliGame := NewCliGame(&game, displaySpy, playerPrompter, 1)

	cliGame.Start()

	assert.True(t, game.IsOver())
	assert.True(t, game.IsADraw())
	assert.True(t, displaySpy.DrawHasBeenCalled)
}

func TestPlaysAWonGame(t *testing.T) {
	playerOne := players.NewFake(marks.X, 0, 0, 1, 2)
	playerTwo := players.NewFake(marks.O, 0, 6, 7)
	displaySpy := new(display.Spy)
	game := cliGameTest.createGame(playerOne, playerTwo)
	cliGame := NewCliGame(&game, displaySpy, playerPrompter, 1)

	cliGame.Start()

	assert.True(t, game.IsOver())
	assert.True(t, game.IsAWin())
	assert.True(t, displaySpy.WinHasBeenCalled)
}

func TestDisplayErrorsForInvalidMove(t *testing.T) {
	playerOne := players.NewFake(marks.X, 0, -1, 0, 1, 2)
	playerTwo := players.NewFake(marks.O, 0, 6, 7)
	displaySpy := new(display.Spy)
	game := cliGameTest.createGame(playerOne, playerTwo)
	cliGame := NewCliGame(&game, displaySpy, playerPrompter, 1)

	cliGame.Start()

	assert.True(t, displaySpy.WriteHasBeenCalled)
	assert.Equal(t, "Out of bounds\n", displaySpy.WriteArgument)
}

func TestPromptsTheCorrectPlayers(t *testing.T) {
	playerOne := players.NewFake(marks.X, 0, -1, 0, 1, 2)
	playerTwo := players.NewFake(marks.O, 0, 6, 7)
	displaySpy := new(display.Spy)
	game := cliGameTest.createGame(playerOne, playerTwo)
	cliGame := NewCliGame(&game, displaySpy, playerPrompter, 2)

	cliGame.Start()

	assert.True(t, displaySpy.HumanPromptHasBeenCalled)
	assert.True(t, displaySpy.ComputerPromptHasBeenCalled)
}

func (gameTest gameTests) createGame(playerOne, playerTwo core.Player) core.Game {
	return core.NewGame(playerOne, playerTwo, core.NewBoard(9), new(core.Rules))
}
