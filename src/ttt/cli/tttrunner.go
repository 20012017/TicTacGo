package cli

import (
	"bufio"
	"os"
	"ttt/cli/display"
	"ttt/cli/input"
	"ttt/cli/input/validators"
	"ttt/cli/players"
	"ttt/core"
	"ttt/core/player"
)

type TTT struct{}

func (ttt TTT) CreateCliGame() Game {
	return NewCliGame(ttt.createGame(), ttt.createDisplay())
}

func (ttt TTT) createDisplay() DisplayWriter {
	return display.NewDisplayWriter(os.Stdout, new(display.Script))
}

func (ttt TTT) createGame() *core.Game {
	playerOne, playerTwo := ttt.createPlayers()
	board := core.NewBoard(9)
	game := core.NewGame(playerOne, playerTwo, board, new(core.Rules))
	return &game
}

func (ttt TTT) createPlayers() (core.Player, core.Player) {
	reader := input.NewReader(bufio.NewReader(os.Stdin))
	validator := new(validators.Move)
	playerOne := players.NewHuman("X", reader, validator)
	playerTwo := player.NewComputer("O", player.NewNegamax(new(core.Rules)))
	return playerOne, playerTwo
}
