package ttt

import (
	"bufio"
	"os"
)

type TTT struct {
}

func StartGame() {
	new(TTT).createCliGame().start()
}

func (ttt TTT) createCliGame() CliGame {
	return NewCliGame(ttt.createGame(), ttt.createDisplay())
}

func (ttt TTT) createDisplay() Display {
	return NewDisplay(os.Stdout, &Script{})
}

func (ttt TTT) createGame() Game {
	playerOne, playerTwo := ttt.createPlayers()
	board := NewBoard(9)
	return NewGame(playerOne, playerTwo, board, rules{})
}

func (ttt TTT) createPlayers() (Player, Player) {
	reader := CliInputReader{bufio.NewReader(os.Stdin)}
	validator := MoveValidator{}
	playerOne := HumanPlayer{"X", reader, validator}
	playerTwo := HumanPlayer{"O", reader, validator}
	return playerOne, playerTwo
}
