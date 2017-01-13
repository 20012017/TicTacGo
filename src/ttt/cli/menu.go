package cli

import (
	"fmt"
	"ttt/cli/input"
	"ttt/cli/input/validators"
	"ttt/cli/players"
	"ttt/core"
)

var options map[int][]int = map[int][]int{
	1: []int{1, 1},
	2: []int{1, 2},
	3: []int{2, 1},
	4: []int{2, 2},
}

type Menu struct {
	display       DisplayWriter
	inputReader   input.InputReader
	playerFactory players.Factory
	menuValidator validators.Menu
}

func NewMenu(display DisplayWriter, inputReader input.InputReader, playerFactory players.Factory, menuValidator validators.Menu) Menu {
	return Menu{display, inputReader, playerFactory, menuValidator}
}

func (menu Menu) show() core.Game {
	menu.display.Clear()
	menu.display.Welcome()
	menu.display.Menu()
	choice := menu.inputReader.Read()
	valid, err := menu.menuValidator.Validate(choice)
	for valid != true {
		menu.display.Write(fmt.Sprintf("%s\n", err.Error()))
		menu.display.InvalidChoice()
		choice = menu.inputReader.Read()
		valid, err = menu.menuValidator.Validate(choice)
	}
	gameChoice := menu.menuValidator.ValidChoice(choice)
	return menu.createGame(gameChoice)
}

func (menu Menu) createGame(gameChoice int) core.Game {
	playerOneType, playerTwoType := menu.getPlayerTypes(gameChoice)
	playerOne := menu.playerFactory.Create(playerOneType, core.MarkX())
	playerTwo := menu.playerFactory.Create(playerTwoType, core.MarkO())
	return core.NewGame(playerOne, playerTwo, core.NewBoard(9), new(core.Rules))
}

func (menu Menu) getPlayerTypes(choice int) (int, int) {
	playerTypes := options[choice]
	return playerTypes[0], playerTypes[1]
}
