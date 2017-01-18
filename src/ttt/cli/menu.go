package cli

import (
	"fmt"
	"ttt/cli/input"
	"ttt/cli/input/validators"
	"ttt/cli/players"
	"ttt/core"
)

type Menu struct {
	display       DisplayWriter
	inputReader   input.InputReader
	playerFactory players.Factory
	menuValidator validators.Input
	gameOptions   map[int][]int
}

func NewMenu(display DisplayWriter,
	inputReader input.InputReader,
	playerFactory players.Factory,
	menuValidator validators.Input,
	gameOptions map[int][]int) Menu {
	return Menu{display, inputReader, playerFactory, menuValidator, gameOptions}
}

func (menu Menu) show() CliGame {
	menu.welcome()
	choice := menu.validGameChoice()
	return menu.createGame(choice)
}

func (menu Menu) replay() bool {
	menu.display.Replay()
	input := menu.readInput()
	replay := input == "yes\n"
	if !replay {
		menu.display.Goodbye()
	}
	return replay
}

func (menu Menu) validGameChoice() int {
	choice := menu.readInput()
	gameChoice, err := menu.menuValidator.Validate(choice)
	for err != nil {
		menu.promptForValidInput(err)
		choice = menu.readInput()
		gameChoice, err = menu.menuValidator.Validate(choice)
	}
	return gameChoice
}

func (menu Menu) promptForValidInput(err error) {
	menu.display.Write(fmt.Sprintf("%s\n", err.Error()))
	menu.display.InvalidChoice()
}

func (menu Menu) readInput() string {
	return menu.inputReader.Read()
}

func (menu Menu) welcome() {
	menu.display.Clear()
	menu.display.Welcome()
	menu.display.Menu()
}

func (menu Menu) createGame(gameChoice int) Game {
	playerOneType, playerTwoType := menu.getPlayerTypes(gameChoice)
	playerOne, playerTwo := menu.createPlayers(playerOneType, playerTwoType)
	game := core.NewGame(playerOne, playerTwo, core.NewBoard(9), new(core.Rules))
	return NewCliGame(&game, menu.display, NewPrompter(menu.gameOptions), gameChoice)
}

func (menu Menu) createPlayers(playerOneType, playerTwoType int) (core.Player, core.Player) {
	playerOne := menu.playerFactory.Create(playerOneType, core.MarkX())
	playerTwo := menu.playerFactory.Create(playerTwoType, core.MarkO())
	return playerOne, playerTwo
}

func (menu Menu) getPlayerTypes(choice int) (int, int) {
	playerTypes := menu.gameOptions[choice]
	return playerTypes[0], playerTypes[1]
}
