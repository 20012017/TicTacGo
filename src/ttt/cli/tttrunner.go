package cli

import (
	"bufio"
	"os"
	"ttt/core"
)

type TTT struct{}

func (ttt TTT) CreateCliGame() CliGame {
	return NewCliGame(ttt.createGame(), ttt.createDisplay())
}

func (ttt TTT) createDisplay() Display {
	return NewDisplay(os.Stdout, &Script{})
}

func (ttt TTT) createGame() core.Game {
	playerOne, playerTwo := ttt.createPlayers()
	board := core.NewBoard(9)
	return core.NewGame(playerOne, playerTwo, board, new(core.Rules))
}

func (ttt TTT) createPlayers() (core.Player, core.Player) {
	reader := CliInputReader{bufio.NewReader(os.Stdin)}
	validator := MoveValidator{}
	playerOne := HumanPlayer{"X", reader, validator}
	playerTwo := HumanPlayer{"O", reader, validator}
	return playerOne, playerTwo
}