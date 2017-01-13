package cli

import (
	"strconv"
	"strings"
	"ttt/cli/input"
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
}

func NewMenu(display DisplayWriter, inputReader input.InputReader, playerFactory players.Factory) Menu {
	return Menu{display, inputReader, playerFactory}
}

func (menu Menu) show() core.Game {
	menu.display.Menu()
	choice := menu.inputReader.Read()
	gameChoice, err := strconv.Atoi(strings.TrimSpace(choice))
	for err != nil || gameChoice < 1 || gameChoice > 4 {
		menu.display.InvalidChoice()
		choice = menu.inputReader.Read()
		gameChoice, err = strconv.Atoi(strings.TrimSpace(choice))
	}
	playerOneType, playerTwoType := menu.getPlayerTypes(gameChoice)
	playerOne := menu.playerFactory.Create(playerOneType, core.MarkX())
	playerTwo := menu.playerFactory.Create(playerTwoType, core.MarkO())
	game := core.NewGame(playerOne, playerTwo, core.NewBoard(9), new(core.Rules))
	return game
}

func (menu Menu) getPlayerTypes(choice int) (int, int) {
	playerTypes := options[choice]
	return playerTypes[0], playerTypes[1]
}
