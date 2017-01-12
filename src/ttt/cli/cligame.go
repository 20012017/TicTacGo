package cli

import (
	"fmt"
	"ttt/core"
)

type CliGame struct {
	game    core.Game
	display DisplayWriter
}

func NewCliGame(game core.Game, display DisplayWriter) CliGame {
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
	cliGame.display.Goodbye()
}

func (cliGame CliGame) displayResult(isWon bool, winner string) {
	if isWon {
		cliGame.display.Win(winner)
	} else {
		cliGame.display.Draw()
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
		cliGame.display.Write(fmt.Sprintf("%s\n", err.Error()))
		cliGame.display.Prompt()
		move, err = cliGame.currentPlayer().Move(cliGame.board())
	}
	return move
}

func (cliGame CliGame) initializeTurn() {
	cliGame.showBoard()
	cliGame.display.Prompt()
}

func (cliGame CliGame) welcome() {
	cliGame.display.Clear()
	cliGame.display.Welcome()
}

func (cliGame CliGame) showBoard() {
	cliGame.display.Clear()
	cliGame.display.ShowBoard(cliGame.board())
}

func (cliGame CliGame) board() core.TTTBoard {
	return cliGame.game.Board()
}

func (cliGame CliGame) currentPlayer() core.Player {
	return cliGame.game.CurrentPlayer()
}
