package cli

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"ttt/cli/display"
	"ttt/cli/input"
	"ttt/cli/input/validators"
	"ttt/cli/players"
	"ttt/core"
)

type MenuTest struct{}

var menuTest MenuTest = MenuTest{}

var playerFactory players.Factory = players.NewPlayerFactory(input.NewInputReaderSpy("1\n"))
var displaySpy *display.Spy = new(display.Spy)
var menuValidator validators.Input = validators.NewInputValidator(1, 4)

func TestPrintsAMenuOfGameChoices(t *testing.T) {
	menu := menuTest.newMenuWithInput("1\n")

	menu.show()

	assert.True(t, displaySpy.MenuWasCalled)
}

func TestReadsInput(t *testing.T) {
	inputReader := input.NewInputReaderSpy("1\n")
	menu := NewMenu(displaySpy, inputReader, playerFactory, menuValidator)

	menu.show()

	assert.True(t, inputReader.WasCalled)
}

func TestReturnsTheCorrectMarksForPlayers(t *testing.T) {
	menu := menuTest.newMenuWithInput("1\n")

	game := menu.show()
	playerOne, playerTwo := menuTest.getPlayers(game)

	assert.Equal(t, "X", playerOne.Mark())
	assert.Equal(t, "O", playerTwo.Mark())
}

func TestReturnsAHumanVHumanGame(t *testing.T) {
	menu := menuTest.newMenuWithInput("1\n")

	game := menu.show()
	playerOne, playerTwo := menuTest.getPlayers(game)
	playerOneType, playerTwoType := menuTest.getPlayerTypes(playerOne, playerTwo)

	assert.Equal(t, "*players.Human", playerOneType)
	assert.Equal(t, "*players.Human", playerTwoType)
}

func TestReturnsAHumanVComputerGame(t *testing.T) {
	menu := menuTest.newMenuWithInput("2\n")

	game := menu.show()
	playerOne, playerTwo := menuTest.getPlayers(game)
	playerOneType, playerTwoType := menuTest.getPlayerTypes(playerOne, playerTwo)

	assert.Equal(t, "*players.Human", playerOneType)
	assert.Equal(t, "*player.Computer", playerTwoType)
}

func TestReturnsAComputerVHumanGame(t *testing.T) {
	menu := menuTest.newMenuWithInput("3\n")

	game := menu.show()
	playerOne, playerTwo := menuTest.getPlayers(game)
	playerOneType, playerTwoType := menuTest.getPlayerTypes(playerOne, playerTwo)

	assert.Equal(t, "*player.Computer", playerOneType)
	assert.Equal(t, "*players.Human", playerTwoType)
}

func TestReturnsAComputerVComputerGame(t *testing.T) {
	menu := menuTest.newMenuWithInput("4\n")

	game := menu.show()
	playerOne, playerTwo := menuTest.getPlayers(game)
	playerOneType, playerTwoType := menuTest.getPlayerTypes(playerOne, playerTwo)

	assert.Equal(t, "*player.Computer", playerOneType)
	assert.Equal(t, "*player.Computer", playerTwoType)
}

func TestLoopsForANumber(t *testing.T) {
	menu := menuTest.newMenuWithInput("hello\n")

	menu.show()

	assert.True(t, displaySpy.InvalidChoiceWasCalled)
}

func TestLoopsForAValidNumber(t *testing.T) {
	menu := menuTest.newMenuWithInput("10\n")

	menu.show()

	assert.True(t, displaySpy.InvalidChoiceWasCalled)
}

func TestDisplaysTheError(t *testing.T) {
	menu := menuTest.newMenuWithInput("10\n")

	menu.show()

	assert.True(t, displaySpy.WriteHasBeenCalled)
	assert.Equal(t, "Not between 1 and 4\n", displaySpy.WriteArgument)
}

func TestClearsTheScreen(t *testing.T) {
	menu := menuTest.newMenuWithInput("4\n")

	menu.show()

	assert.True(t, displaySpy.ClearHasBeenCalled)
}

func TestWelcomesThePlayer(t *testing.T) {
	menu := menuTest.newMenuWithInput("4\n")

	menu.show()

	assert.True(t, displaySpy.WelcomeHasBeenCalled)
}

func (menuTest MenuTest) newMenuWithInput(userInput string) Menu {
	inputReader := input.NewInputReaderSpy(userInput, "1\n")
	return NewMenu(displaySpy, inputReader, playerFactory, menuValidator)
}

func (menuTest MenuTest) getPlayers(game core.Game) (playerOne, playerTwo core.Player) {
	playerOne = game.CurrentPlayer()
	game.Play(0)
	playerTwo = game.CurrentPlayer()
	return playerOne, playerTwo
}

func (menuTest MenuTest) getPlayerTypes(playerOne, playerTwo core.Player) (string, string) {
	return menuTest.getPlayerType(playerOne),
		menuTest.getPlayerType(playerTwo)
}

func (menuTest MenuTest) getPlayerType(player core.Player) string {
	playerType := reflect.TypeOf(player)
	return fmt.Sprint(playerType)
}