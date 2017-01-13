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
var moveValidator validators.Menu = validators.NewMenuValidator(1, 4)

func TestPrintsAMenuOfGameChoices(t *testing.T) {
	inputReader := input.NewInputReaderSpy("1\n")
	menu := NewMenu(displaySpy, inputReader, playerFactory, moveValidator)

	menu.show()

	assert.True(t, displaySpy.MenuWasCalled)
}

func TestReadsInput(t *testing.T) {
	inputReader := input.NewInputReaderSpy("1\n")
	menu := NewMenu(displaySpy, inputReader, playerFactory, moveValidator)

	menu.show()

	assert.True(t, inputReader.WasCalled)
}

func TestReturnsTheCorrectMarksForPlayers(t *testing.T) {
	inputReader := input.NewInputReaderSpy("1\n")
	menu := NewMenu(displaySpy, inputReader, playerFactory, moveValidator)

	game := menu.show()
	playerOne, playerTwo := menuTest.getPlayers(game)

	assert.Equal(t, "X", playerOne.Mark())
	assert.Equal(t, "O", playerTwo.Mark())
}

func TestReturnsAHumanVHumanGame(t *testing.T) {
	inputReader := input.NewInputReaderSpy("1\n")
	menu := NewMenu(displaySpy, inputReader, playerFactory, moveValidator)

	game := menu.show()
	playerOne, playerTwo := menuTest.getPlayers(game)
	playerOneType, playerTwoType := menuTest.getPlayerTypes(playerOne, playerTwo)

	assert.Equal(t, "*players.Human", playerOneType)
	assert.Equal(t, "*players.Human", playerTwoType)
}

func TestReturnsAHumanVComputerGame(t *testing.T) {
	inputReader := input.NewInputReaderSpy("2\n")
	menu := NewMenu(displaySpy, inputReader, playerFactory, moveValidator)

	game := menu.show()
	playerOne, playerTwo := menuTest.getPlayers(game)
	playerOneType, playerTwoType := menuTest.getPlayerTypes(playerOne, playerTwo)

	assert.Equal(t, "*players.Human", playerOneType)
	assert.Equal(t, "*player.Computer", playerTwoType)
}

func TestReturnsAComputerVHumanGame(t *testing.T) {
	inputReader := input.NewInputReaderSpy("3\n")
	menu := NewMenu(displaySpy, inputReader, playerFactory, moveValidator)

	game := menu.show()
	playerOne, playerTwo := menuTest.getPlayers(game)
	playerOneType, playerTwoType := menuTest.getPlayerTypes(playerOne, playerTwo)

	assert.Equal(t, "*player.Computer", playerOneType)
	assert.Equal(t, "*players.Human", playerTwoType)
}

func TestReturnsAComputerVComputerGame(t *testing.T) {
	inputReader := input.NewInputReaderSpy("4\n")
	menu := NewMenu(displaySpy, inputReader, playerFactory, moveValidator)

	game := menu.show()
	playerOne, playerTwo := menuTest.getPlayers(game)
	playerOneType, playerTwoType := menuTest.getPlayerTypes(playerOne, playerTwo)

	assert.Equal(t, "*player.Computer", playerOneType)
	assert.Equal(t, "*player.Computer", playerTwoType)
}

func TestLoopsForANumber(t *testing.T) {
	inputReader := input.NewInputReaderSpy("hello\n", "4\n")
	menu := NewMenu(displaySpy, inputReader, playerFactory, moveValidator)

	menu.show()

	assert.True(t, displaySpy.InvalidChoiceWasCalled)
}

func TestLoopsForAValidNumber(t *testing.T) {
	inputReader := input.NewInputReaderSpy("10\n", "4\n")
	menu := NewMenu(displaySpy, inputReader, playerFactory, moveValidator)

	menu.show()

	assert.True(t, displaySpy.InvalidChoiceWasCalled)
}

func TestDisplaysTheError(t *testing.T) {
	inputReader := input.NewInputReaderSpy("10\n", "4\n")
	menu := NewMenu(displaySpy, inputReader, playerFactory, moveValidator)

	menu.show()

	assert.True(t, displaySpy.WriteHasBeenCalled)
	assert.Equal(t, "Not in menu", displaySpy.WriteArgument)
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
