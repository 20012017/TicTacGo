package cli

import (
	"fmt"
	"ttt/core"
)

type CliGame struct {
	game    core.Game
	display Display
}

func NewCliGame(game core.Game, display Display) CliGame {
	return CliGame{game, display}
}

func (cliGame CliGame) Start() {
	cliGame.welcome()
	cliGame.play()
	cliGame.end()
}

func (cliGame CliGame) end() {
	cliGame.showBoard()
	cliGame.displayResult(cliGame.game.Result())
	cliGame.display.goodbye()
}

func (cliGame CliGame) displayResult(isWon bool, winner string) {
	if isWon {
		cliGame.display.win(winner)
	} else {
		cliGame.display.draw()
	}
}

func (cliGame CliGame) play() {
	for !cliGame.game.IsOver() {
		cliGame.initializeTurn()
		cliGame.game.Play(cliGame.getValidMove())
	}
}

func (cliGame CliGame) getValidMove() int {
	move, err := cliGame.currentPlayer().Move(cliGame.board())
	for err != nil {
		cliGame.display.write(fmt.Sprintf("%s\n", err.Error()))
		cliGame.display.prompt()
		move, err = cliGame.currentPlayer().Move(cliGame.board())
	}
	return move
}

func (cliGame CliGame) initializeTurn() {
	cliGame.showBoard()
	cliGame.display.prompt()
}

func (cliGame CliGame) welcome() {
	cliGame.display.clear()
	cliGame.display.welcome()
}

func (cliGame CliGame) showBoard() {
	cliGame.display.clear()
	cliGame.display.showBoard(cliGame.board())
}

func (cliGame CliGame) board() core.Board {
	return cliGame.game.Board()
}

func (cliGame CliGame) currentPlayer() core.Player {
	return cliGame.game.CurrentPlayer()
}
